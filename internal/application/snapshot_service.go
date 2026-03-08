package application

import (
	"context"
	"sync"

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

	mu        sync.RWMutex
	snapshots map[domain.MatchID]*SnapshotCache
}

func NewSnapshotService(snapshotStore api.SnapshotStore, eventStore api.EventStore) *SnapshotService {
	return &SnapshotService{
		snapshotStore: snapshotStore,
		eventStore:    eventStore,
		snapshots:     make(map[domain.MatchID]*SnapshotCache),
		mu:            sync.RWMutex{},
	}
}

func (s *SnapshotService) GetSnapshot(ctx context.Context, matchID domain.MatchID) (*domain.Match, error) {
	snapshotCache, err := s.getOrCreateSnapshotCache(ctx, matchID)
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if !snapshotCache.Stale {
		return snapshotCache.Match.Clone(), nil
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
	go s.snapshotStore.Save(ctx, matchID, version, snapshotCache.Match)
	return snapshotCache.Match.Clone(), nil
}

func (s *SnapshotService) MarkSnapshotStale(matchID domain.MatchID) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	snapshot, ok := s.snapshots[matchID]
	if !ok {
		return nil
	}
	snapshot.Stale = true
	return nil
}

func (s *SnapshotService) getOrCreateSnapshotCache(ctx context.Context, matchID domain.MatchID) (*SnapshotCache, error) {
	snapshot, ok := s.getSnapshotCache(matchID)
	if !ok {
		match, version, err := s.snapshotStore.Load(ctx, matchID)
		if err != nil {
			return nil, err
		}
		snapshot = &SnapshotCache{
			Match:   match,
			Version: version,
			Stale:   true,
		}
		s.setSnapshotCache(matchID, snapshot)
	}
	return snapshot, nil
}

func (s *SnapshotService) getSnapshotCache(matchID domain.MatchID) (*SnapshotCache, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	snapshot, ok := s.snapshots[matchID]
	return snapshot, ok
}

func (s *SnapshotService) setSnapshotCache(matchID domain.MatchID, snapshot *SnapshotCache) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.snapshots[matchID] = snapshot
}
