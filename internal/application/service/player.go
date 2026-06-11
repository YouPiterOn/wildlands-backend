package service

import (
	"context"
	"log/slog"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type Player struct {
	playerStore api.PlayerStore
	logger      api.Logger
}

var _ api.PlayerService = (*Player)(nil)

func NewPlayer(playerStore api.PlayerStore, logger api.Logger) *Player {
	return &Player{playerStore: playerStore, logger: logger.With(slog.String("tag", "PlayerService"))}
}

func (p *Player) Create(ctx context.Context, name string) (*domain.Player, error) {
	player, err := p.playerStore.Create(ctx, name)
	if err != nil {
		p.logger.Error(utils.ErrPlayerCreate.Error(), slog.String("error", err.Error()))
		return nil, utils.ErrPlayerCreate
	}
	return player, nil
}

func (p *Player) GetByID(ctx context.Context, id domain.PlayerID) (*domain.Player, error) {
	player, err := p.playerStore.GetByID(ctx, id)
	if err != nil {
		p.logger.Error(utils.ErrPlayerGetByID.Error(), slog.String("error", err.Error()))
		return nil, utils.ErrPlayerGetByID
	}
	return player, nil
}
