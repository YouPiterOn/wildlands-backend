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
		SELECT state, seats_count, current_turn FROM snapshots WHERE match_id = $1
		`,
		matchID.String(),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var state domain.MatchState
		var seatsCount int
		var currentTurn int
		err := rows.Scan(&state, &seatsCount, &currentTurn)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
