package application

import (
	"context"
	"log/slog"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type AggregateMatchRepository struct {
	snapshotStore api.SnapshotStore
	eventStore    api.EventStore
	logger        api.Logger
}

var _ api.MatchRepository = (*AggregateMatchRepository)(nil)

func NewAggregateMatchRepository(snapshotStore api.SnapshotStore, eventStore api.EventStore, logger api.Logger) *AggregateMatchRepository {
	return &AggregateMatchRepository{snapshotStore: snapshotStore, eventStore: eventStore, logger: logger}
}

func (r *AggregateMatchRepository) Load(ctx context.Context, id domain.MatchID) (*domain.Match, error) {
	match, err := r.snapshotStore.Load(ctx, id)
	if err != nil {
		match = domain.NewMatch(id)
	}

	events, _, err := r.eventStore.LoadSince(ctx, id, match.Version)
	if err != nil {
		r.logger.Error("error loading events",
			slog.String("match_id", id.String()),
			slog.Int("version", match.Version),
			slog.String("error", err.Error()))

		return nil, ErrEventsLoad
	}

	if len(events) == 0 {
		return match, nil
	}

	for _, e := range events {
		err := match.Apply(e)
		if err != nil {
			r.logger.Error("error applying event",
				slog.String("match_id", id.String()),
				slog.Int("version", match.Version),
				slog.String("event_type", e.EventType().String()),
				slog.String("error", err.Error()))

			return nil, ErrEventApply
		}
	}

	return match, nil
}
