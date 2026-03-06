package application

import (
	"context"

	"youpiteron.dev/wildlands-backend/internal/api"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

type SnapshotCache struct {
	Match   *domain.Match
	Version int
	Stale   bool
}

type SnapshotService struct {
	snapshotStore api.SnapshotStore
	eventStore    api.EventStore
	snapshots     map[domain.MatchID]*SnapshotCache
}

func NewSnapshotService(snapshotStore api.SnapshotStore, eventStore api.EventStore) *SnapshotService {
	return &SnapshotService{snapshotStore: snapshotStore, eventStore: eventStore}
}

func (s *SnapshotService) GetSnapshot(ctx context.Context, matchID domain.MatchID) (*domain.Match, error) {
	snapshot, ok := s.snapshots[matchID]
	if !ok {
		s.snapshotStore.Load(ctx, matchID)
	}
	events, version, err := s.eventStore.LoadSince(ctx, matchID, snapshot.Version)
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		event.Apply(snapshot.Match)
	}
	snapshot.Version = version
	snapshot.Stale = false
	return snapshot.Match, nil
}
