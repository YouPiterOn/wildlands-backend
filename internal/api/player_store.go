package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type PlayerStore interface {
	Create(ctx context.Context, name string) (*domain.Player, error)
	GetByID(ctx context.Context, id domain.PlayerID) (*domain.Player, error)
}
