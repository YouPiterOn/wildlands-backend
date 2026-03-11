package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchRepository interface {
	Load(ctx context.Context, id domain.MatchID) (*domain.Match, error)
}
