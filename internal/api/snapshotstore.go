package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type SnapshotStore interface {
	Save(ctx context.Context, matchID domain.MatchID, version int, snapshot *domain.Match) error
	Load(ctx context.Context, matchID domain.MatchID) (*domain.Match, int, error)
}
