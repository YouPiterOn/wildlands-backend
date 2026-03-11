package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type SnapshotStore interface {
	Save(ctx context.Context, snapshot *domain.Match) error
	Load(ctx context.Context, matchID domain.MatchID) (*domain.Match, error)
}
