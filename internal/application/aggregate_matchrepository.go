package application

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type AggregateMatchRepository struct {
	snapshotStore api.SnapshotStore
	eventStore    api.EventStore
}

var _ api.MatchRepository = (*AggregateMatchRepository)(nil)

func NewAggregateMatchRepository(snapshotStore api.SnapshotStore, eventStore api.EventStore) *AggregateMatchRepository {
	return &AggregateMatchRepository{snapshotStore: snapshotStore, eventStore: eventStore}
}

func (r *AggregateMatchRepository) Load(ctx context.Context, id domain.MatchID) (*domain.Match, error) {
	match, err := r.snapshotStore.Load(ctx, id)
	if err != nil {
		match = domain.NewMatch(id)
	}

	events, _, err := r.eventStore.LoadSince(ctx, id, match.Version)
	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return match, nil
	}

	for _, e := range events {
		err := match.Apply(e)
		if err != nil {
			return nil, err
		}
	}

	return match, nil
}
