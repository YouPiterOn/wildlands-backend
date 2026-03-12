package application

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchService struct {
	eventStore      api.EventStore
	matchRepository api.MatchRepository
}

func NewMatchService(eventStore api.EventStore, matchRepository api.MatchRepository) *MatchService {
	return &MatchService{eventStore: eventStore, matchRepository: matchRepository}
}

func (s *MatchService) HandleCommand(ctx context.Context, command domain.Command) ([]domain.Event, error) {
	match, err := s.matchRepository.Load(ctx, command.GetMatchID())
	if err != nil {
		return nil, err
	}
	events, err := command.Handle(match)
	if err != nil {
		return nil, err
	}
	err = s.eventStore.Append(ctx, match.ID, match.Version, events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
