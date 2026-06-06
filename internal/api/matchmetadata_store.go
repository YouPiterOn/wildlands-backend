package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchMetadataStore interface {
	Create(ctx context.Context, metadata *domain.MatchMetadata) error
	Update(ctx context.Context, metadata *domain.MatchMetadata) error
	GetByMatchID(ctx context.Context, matchID domain.MatchID) (*domain.MatchMetadata, error)
}
