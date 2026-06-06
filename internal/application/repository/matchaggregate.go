package repository

import (
	"context"
	"log/slog"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type MatchAggregate struct {
	snapshotStore      api.SnapshotStore
	eventStore         api.EventStore
	matchMetadataStore api.MatchMetadataStore
	logger             api.Logger
}

var _ api.MatchAggregateRepository = (*MatchAggregate)(nil)

func NewMatchAggregate(
	snapshotStore api.SnapshotStore,
	eventStore api.EventStore,
	matchMetadataStore api.MatchMetadataStore,
	logger api.Logger,
) *MatchAggregate {
	return &MatchAggregate{
		snapshotStore:      snapshotStore,
		eventStore:         eventStore,
		matchMetadataStore: matchMetadataStore,
		logger:             logger,
	}
}

func (r *MatchAggregate) Load(ctx context.Context, id domain.MatchID) (*domain.Match, error) {
	match, err := r.snapshotStore.Load(ctx, id)
	if err != nil {
		match = domain.NewMatch(id)
	}

	metadata, err := r.matchMetadataStore.GetByMatchID(ctx, id)
	if err != nil {
		r.logger.Error("error loading match metadata",
			slog.String("match_id", id.String()),
			slog.String("error", err.Error()))

		return nil, utils.ErrMatchMetadataLoad
	}

	events, _, err := r.eventStore.LoadSince(ctx, id, match.Version)
	if err != nil {
		r.logger.Error("error loading events",
			slog.String("match_id", id.String()),
			slog.Int("version", match.Version),
			slog.String("error", err.Error()))

		return nil, utils.ErrEventLoad
	}

	if len(events) == 0 {
		return match, nil
	}

	for _, e := range events {
		err := match.Apply(e, metadata)
		if err != nil {
			r.logger.Error("error applying event",
				slog.String("match_id", id.String()),
				slog.Int("version", match.Version),
				slog.String("event_type", e.EventType().String()),
				slog.String("error", err.Error()))

			return nil, utils.ErrEventApply
		}
	}

	return match, nil
}
