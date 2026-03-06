package application

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchService struct {
	eventStore api.EventStore
}

func NewMatchService(eventStore api.EventStore) *MatchService {
	return &MatchService{eventStore: eventStore}
}

func (s *MatchService) CreateMatch(ctx context.Context, matchID domain.MatchID, seatsCount int) error {
	command := domain.CommandCreateMatch{
		MatchID:    matchID,
		SeatsCount: seatsCount,
	}
	events, err := command.Handle(nil)
	if err != nil {
		return err
	}
}
