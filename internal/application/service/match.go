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
	playerStore              api.PlayerStore
	matchAggregateRepository api.MatchAggregateRepository
	logger                   api.Logger
}

var _ api.MatchService = (*Match)(nil)

func NewMatch(
	eventStore api.EventStore,
	matchMetadataStore api.MatchMetadataStore,
	playerStore api.PlayerStore,
	matchAggregateRepository api.MatchAggregateRepository,
	logger api.Logger,
) *Match {
	return &Match{
		eventStore:               eventStore,
		matchMetadataStore:       matchMetadataStore,
		playerStore:              playerStore,
		matchAggregateRepository: matchAggregateRepository,
		logger:                   logger.With(slog.String("tag", "MatchService")),
	}
}

func (m *Match) CreateMatch(ctx context.Context, playerID domain.PlayerID) (*domain.Match, error) {
	metadata, err := m.generateMatchMetadata()
	if err != nil {
		m.logger.Error(utils.ErrMatchMetadataGenerate.Error(),
			slog.String("error", err.Error()))

		return nil, utils.ErrMatchMetadataGenerate
	}
	err = m.matchMetadataStore.Create(ctx, metadata)
	if err != nil {
		m.logger.Error(utils.ErrMatchMetadataCreate.Error(),
			slog.String("match_id", metadata.MatchID.String()),
			slog.Int64("boards_seed", metadata.BoardsSeed),
			slog.Int64("explore_deck_seed", metadata.ExploreDeckSeed),
			slog.Int64("scoring_cards_seed", metadata.ScoringCardsSeed),
			slog.String("error", err.Error()))

		return nil, utils.ErrMatchMetadataCreate
	}

	m.logger.Info("match created",
		slog.String("match_id", metadata.MatchID.String()),
		slog.String("player_id", playerID.String()))

	return domain.NewMatch(metadata.MatchID), nil
}

func (m *Match) JoinMatch(ctx context.Context, matchID domain.MatchID, playerID domain.PlayerID) (*domain.Match, error) {
	player, err := m.playerStore.GetByID(ctx, playerID)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, utils.ErrPlayerNotFound
	}

	command := domain.CommandJoinMatch{
		MatchID:  matchID,
		PlayerID: playerID,
	}

	match, err := m.matchAggregateRepository.Load(ctx, command.GetMatchID())
	if err != nil {
		return nil, err
	}

	events, err := command.Handle(match)
	if err != nil {
		m.logger.Error(utils.ErrCommandHandle.Error(),
			slog.String("match_id", command.GetMatchID().String()),
			slog.Int("version", match.Version),
			slog.String("command_type", command.Type().String()),
			slog.String("error", err.Error()))

		return nil, utils.ErrCommandHandle
	}

	err = m.eventStore.Append(ctx, match.ID, match.Version, events)
	if err != nil {
		m.logger.Error(utils.ErrEventAppend.Error(),
			slog.String("match_id", match.ID.String()),
			slog.Int("version", match.Version),
			slog.String("error", err.Error()))

		return nil, utils.ErrEventAppend
	}

	match, err = m.matchAggregateRepository.Load(ctx, match.ID)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (m *Match) HandleCommand(ctx context.Context, command domain.Command) ([]domain.Event, error) {
	match, err := m.matchAggregateRepository.Load(ctx, command.GetMatchID())
	if err != nil {
		return nil, err
	}

	events, err := command.Handle(match)
	if err != nil {
		m.logger.Error(utils.ErrCommandHandle.Error(),
			slog.String("match_id", command.GetMatchID().String()),
			slog.Int("version", match.Version),
			slog.String("command_type", command.Type().String()),
			slog.String("error", err.Error()))

		return nil, utils.ErrCommandHandle
	}

	err = m.eventStore.Append(ctx, match.ID, match.Version, events)
	if err != nil {
		m.logger.Error(utils.ErrEventAppend.Error(),
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
