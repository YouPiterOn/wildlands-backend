package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchAggregateRepository interface {
	Load(ctx context.Context, id domain.MatchID) (*domain.Match, error)
}
