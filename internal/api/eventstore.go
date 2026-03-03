package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type EventStore interface {
	Append(ctx context.Context, event domain.Event) error
	Load(ctx context.Context, matchID string) ([]domain.Event, error)
}
