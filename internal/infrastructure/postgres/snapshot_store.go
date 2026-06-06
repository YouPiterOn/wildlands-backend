package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type SnapshotStore struct {
	pool *pgxpool.Pool
}

var _ api.SnapshotStore = (*SnapshotStore)(nil)

func NewSnapshotStore(pool *pgxpool.Pool) *SnapshotStore {
	return &SnapshotStore{pool: pool}
}

func (s *SnapshotStore) Save(ctx context.Context, snapshot *domain.Match) error {
	storedMatch, err := ToSnapshotRow(snapshot)
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

func (s *SnapshotStore) Load(ctx context.Context, matchID domain.MatchID) (*domain.Match, error) {
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
	var snapshotRow SnapshotRow
	err = rows.Scan(&snapshotRow.MatchID, &snapshotRow.State, &snapshotRow.CurrentTurn, &snapshotRow.Version)
	if err != nil {
		return nil, err
	}
	match, err := snapshotRow.ToDomainMatch()
	if err != nil {
		return nil, err
	}
	return match, nil
}
