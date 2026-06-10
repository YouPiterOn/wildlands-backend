package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchMetadataStore struct {
	pool *pgxpool.Pool
}

var _ api.MatchMetadataStore = (*MatchMetadataStore)(nil)

func NewMatchMetadataStore(pool *pgxpool.Pool) *MatchMetadataStore {
	return &MatchMetadataStore{pool: pool}
}

func (s *MatchMetadataStore) Create(ctx context.Context, metadata *domain.MatchMetadata) error {
	storedMetadata, err := ToMatchMetadataRow(metadata)
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
		INSERT INTO match_metadata (match_id, boards_seed, explore_deck_seed, scoring_cards_seed)
		VALUES ($1, $2, $3, $4)
		`,
		storedMetadata.MatchID,
		storedMetadata.BoardsSeed,
		storedMetadata.ExploreDeckSeed,
		storedMetadata.ScoringCardsSeed,
	)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *MatchMetadataStore) Update(ctx context.Context, metadata *domain.MatchMetadata) error {
	storedMetadata, err := ToMatchMetadataRow(metadata)
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
		UPDATE match_metadata
		SET boards_seed = $2, explore_deck_seed = $3, scoring_cards_seed = $4
		WHERE match_id = $1
		`,
		storedMetadata.MatchID,
		storedMetadata.BoardsSeed,
		storedMetadata.ExploreDeckSeed,
		storedMetadata.ScoringCardsSeed,
	)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (s *MatchMetadataStore) GetByMatchID(ctx context.Context, matchID domain.MatchID) (*domain.MatchMetadata, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(
		ctx,
		`
		SELECT match_id, boards_seed, explore_deck_seed, scoring_cards_seed
		FROM match_metadata
		WHERE match_id = $1
		`,
		matchID.String(),
	)
	var storedMetadata MatchMetadataRow
	err = row.Scan(
		&storedMetadata.MatchID,
		&storedMetadata.BoardsSeed,
		&storedMetadata.ExploreDeckSeed,
		&storedMetadata.ScoringCardsSeed,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return storedMetadata.ToDomainMatchMetadata()
}
