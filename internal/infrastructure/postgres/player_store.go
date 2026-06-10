package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type PlayerStore struct {
	pool *pgxpool.Pool
}

var _ api.PlayerStore = (*PlayerStore)(nil)

func NewPlayerStore(pool *pgxpool.Pool) *PlayerStore {
	return &PlayerStore{pool: pool}
}

func (s *PlayerStore) Create(ctx context.Context, name string) (*domain.Player, error) {
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
	id, err := domain.NewPlayerID()
	if err != nil {
		return nil, err
	}
	storedPlayer, err := ToPlayerRow(&domain.Player{ID: id, Name: name})
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(
		ctx,
		`
		INSERT INTO players (id, name)
		VALUES ($1, $2)
		`,
		storedPlayer.ID, storedPlayer.Name,
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return storedPlayer.ToDomainPlayer()
}

func (s *PlayerStore) GetByID(ctx context.Context, id domain.PlayerID) (*domain.Player, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(
		ctx,
		`
		SELECT id, name FROM players WHERE id = $1
		`,
		id.String(),
	)
	var storedPlayer PlayerRow
	err = row.Scan(&storedPlayer.ID, &storedPlayer.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return storedPlayer.ToDomainPlayer()
}
