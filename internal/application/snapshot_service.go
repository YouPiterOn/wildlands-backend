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
	snapshotCache := s.getOrCreateSnapshotCache(ctx, matchID)
	if !snapshotCache.Stale {
		return snapshotCache.Match, nil
	}
	events, version, err := s.eventStore.LoadSince(ctx, matchID, snapshotCache.Version)
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		event.Apply(snapshotCache.Match)
	}
	snapshotCache.Version = version
	snapshotCache.Stale = false
	s.snapshotStore.Save(ctx, matchID, snapshotCache.Match)
	return snapshotCache.Match, nil
}

func (s *SnapshotService) StaleSnapshotCache(ctx context.Context, matchID domain.MatchID) error {
	snapshot, ok := s.snapshots[matchID]
	if !ok {
		return nil
	}
	snapshot.Stale = true
	return nil
}

func (s *SnapshotService) getOrCreateSnapshotCache(ctx context.Context, matchID domain.MatchID) *SnapshotCache {
	snapshot, ok := s.snapshots[matchID]
	if !ok {
		match, version, err := s.snapshotStore.Load(ctx, matchID)
		if err != nil {
			match = nil
			version = 0
		}
		snapshot = &SnapshotCache{
			Match:   match,
			Version: version,
			Stale:   true,
		}
		s.snapshots[matchID] = snapshot
	}
	return snapshot
}
