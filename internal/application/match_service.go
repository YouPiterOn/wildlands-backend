package application

import (
	"context"
	"log/slog"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type MatchService struct {
	eventStore      api.EventStore
	matchRepository api.MatchRepository
	logger          api.Logger
}

func NewMatchService(eventStore api.EventStore, matchRepository api.MatchRepository, logger api.Logger) *MatchService {
	return &MatchService{eventStore: eventStore, matchRepository: matchRepository, logger: logger}
}

func (s *MatchService) HandleCommand(ctx context.Context, command domain.Command) ([]domain.Event, error) {
	match, err := s.matchRepository.Load(ctx, command.GetMatchID())
	if err != nil {
		return nil, err
	}

	events, err := command.Handle(match)
	if err != nil {
		s.logger.Error("error handling command",
			slog.String("match_id", command.GetMatchID().String()),
			slog.Int("version", match.Version),
			slog.String("command_type", command.Type().String()),
			slog.String("error", err.Error()))

		return nil, ErrCommandHandle
	}

	err = s.eventStore.Append(ctx, match.ID, match.Version, events)
	if err != nil {
		s.logger.Error("error appending events",
			slog.String("match_id", match.ID.String()),
			slog.Int("version", match.Version),
			slog.String("error", err.Error()))

		return nil, ErrEventsAppend
	}
	return events, nil
}
