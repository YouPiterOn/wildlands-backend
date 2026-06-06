package service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type Match struct {
	eventStore               api.EventStore
	matchMetadataStore       api.MatchMetadataStore
	matchAggregateRepository api.MatchAggregateRepository
	logger                   api.Logger
}

func NewMatch(
	eventStore api.EventStore,
	matchMetadataStore api.MatchMetadataStore,
	matchAggregateRepository api.MatchAggregateRepository,
	logger api.Logger,
) *Match {
	return &Match{
		eventStore:               eventStore,
		matchMetadataStore:       matchMetadataStore,
		matchAggregateRepository: matchAggregateRepository,
		logger:                   logger,
	}
}

func (m *Match) CreateMatch(ctx context.Context) (domain.MatchID, error) {
	metadata, err := m.generateMatchMetadata()
	if err != nil {
		m.logger.Error("error generating match metadata",
			slog.String("error", err.Error()))

		return domain.MatchID{}, utils.ErrMatchMetadataGenerate
	}
	err = m.matchMetadataStore.Create(ctx, metadata)
	if err != nil {
		m.logger.Error("error creating match metadata",
			slog.String("match_id", metadata.MatchID.String()),
			slog.Int64("boards_seed", metadata.BoardsSeed),
			slog.Int64("explore_deck_seed", metadata.ExploreDeckSeed),
			slog.Int64("scoring_cards_seed", metadata.ScoringCardsSeed),
			slog.String("error", err.Error()))

		return domain.MatchID{}, utils.ErrMatchMetadataCreate
	}

	return metadata.MatchID, nil
}

func (m *Match) HandleCommand(ctx context.Context, command domain.Command) ([]domain.Event, error) {
	match, err := m.matchAggregateRepository.Load(ctx, command.GetMatchID())
	if err != nil {
		return nil, err
	}

	events, err := command.Handle(match)
	if err != nil {
		m.logger.Error("error handling command",
			slog.String("match_id", command.GetMatchID().String()),
			slog.Int("version", match.Version),
			slog.String("command_type", command.Type().String()),
			slog.String("error", err.Error()))

		return nil, utils.ErrCommandHandle
	}

	err = m.eventStore.Append(ctx, match.ID, match.Version, events)
	if err != nil {
		m.logger.Error("error appending events",
			slog.String("match_id", match.ID.String()),
			slog.Int("version", match.Version),
			slog.String("error", err.Error()))

		return nil, utils.ErrEventAppend
	}
	return events, nil
}

func (m *Match) generateMatchMetadata() (*domain.MatchMetadata, error) {
	matchId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	boardsSeed, err := domain.GenerateSeed()
	if err != nil {
		return nil, err
	}
	exploreDeckSeed, err := domain.GenerateSeed()
	if err != nil {
		return nil, err
	}
	scoringCardsSeed, err := domain.GenerateSeed()
	if err != nil {
		return nil, err
	}
	return domain.NewMatchMetadata(domain.MatchID(matchId), boardsSeed, exploreDeckSeed, scoringCardsSeed), nil
}
