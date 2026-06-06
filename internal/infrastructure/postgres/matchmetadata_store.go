package postgres

import (
	"context"

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
	return nil
}

func (s *MatchMetadataStore) Update(ctx context.Context, metadata *domain.MatchMetadata) error {
	return nil
}

func (s *MatchMetadataStore) GetByMatchID(ctx context.Context, matchID domain.MatchID) (*domain.MatchMetadata, error) {
	return nil, nil
}
