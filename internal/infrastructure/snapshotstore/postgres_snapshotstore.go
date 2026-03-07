package snapshotstore

import (
	"context"
	"errors"

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

func (s *PostgresSnapshotStore) Save(ctx context.Context, matchID domain.MatchID, snapshot *domain.Match) error {
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
    INSERT INTO snapshots (match_id, state, seats_count, current_turn)
    VALUES ($1, $2, $3, $4)
    `,
		matchID.String(), snapshot.State.String(), snapshot.SeatsCount, snapshot.CurrentTurn,
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

func (s *PostgresSnapshotStore) Load(ctx context.Context, matchID domain.MatchID) (*domain.Match, int, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback(ctx)
	rows, err := tx.Query(
		ctx,
		`
		SELECT id, state, seats_count, current_turn, version FROM snapshots WHERE match_id = $1 ORDER BY version DESC LIMIT 1
		`,
		matchID.String(),
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, 0, errors.New("No snapshot found")
	}
	var storedMatch StoredMatch
	err = rows.Scan(&storedMatch.MatchID, &storedMatch.State, &storedMatch.SeatsCount, &storedMatch.CurrentTurn, &storedMatch.Version)
	if err != nil {
		return nil, 0, err
	}
	match, version, err := storedMatch.ToDomainMatch()
	if err != nil {
		return nil, 0, err
	}
	return match, version, nil
}
