package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type EventStore interface {
	Append(ctx context.Context, matchID domain.MatchID, expectedVersion int, events ...domain.Event) error
	LoadAll(ctx context.Context, matchID domain.MatchID) ([]domain.Event, int, error)
	LoadSince(ctx context.Context, matchID domain.MatchID, version int) ([]domain.Event, int, error)
}
