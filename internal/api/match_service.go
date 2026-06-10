package api

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchService interface {
	CreateMatch(ctx context.Context, playerID domain.PlayerID) (*domain.Match, error)
	JoinMatch(ctx context.Context, matchID domain.MatchID, playerID domain.PlayerID) (*domain.Match, error)
	HandleCommand(ctx context.Context, command domain.Command) ([]domain.Event, error)
}
