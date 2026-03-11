package snapshotstore

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type PostgresSnapshotStore struct {
	pool *pgxpool.Pool
}

var _ api.SnapshotStore = (*PostgresSnapshotStore)(nil)

func NewPostgresSnapshotStore(pool *pgxpool.Pool) *PostgresSnapshotStore {
	return &PostgresSnapshotStore{pool: pool}
}

func (s *PostgresSnapshotStore) Save(ctx context.Context, snapshot *domain.Match) error {
	storedMatch, err := ToStoredMatch(snapshot)
	if err != nil {
		return err
	}
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(
		ctx,
		`
    INSERT INTO snapshots (match_id, state, current_turn, version)
    VALUES ($1, $2, $3, $4)
    `,
		storedMatch.MatchID, storedMatch.State, storedMatch.CurrentTurn, storedMatch.Version,
	)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresSnapshotStore) Load(ctx context.Context, matchID domain.MatchID) (*domain.Match, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	rows, err := tx.Query(
		ctx,
		`
		SELECT id, state, current_turn, version FROM snapshots WHERE match_id = $1 ORDER BY version DESC LIMIT 1
		`,
		matchID.String(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	var storedMatch StoredMatch
	err = rows.Scan(&storedMatch.MatchID, &storedMatch.State, &storedMatch.CurrentTurn, &storedMatch.Version)
	if err != nil {
		return nil, err
	}
	match, err := storedMatch.ToDomainMatch()
	if err != nil {
		return nil, err
	}
	return match, nil
}
