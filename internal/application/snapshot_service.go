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
	snapshot, err := s.loadSnapshotCache(ctx, matchID)
	if err != nil {
		snapshot, err = s.createNewSnapshot(ctx, matchID)
		if err != nil {
			return nil, err
		}
	}
	if !snapshot.Stale {
		return snapshot.Match, nil
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

func (s *SnapshotService) StaleSnapshot(ctx context.Context, matchID domain.MatchID) error {
	snapshot, ok := s.snapshots[matchID]
	if !ok {
		return nil
	}
	snapshot.Stale = true
	return nil
}

func (s *SnapshotService) loadSnapshotCache(ctx context.Context, matchID domain.MatchID) (*SnapshotCache, error) {
	snapshot, ok := s.snapshots[matchID]
	if !ok {
		match, version, err := s.snapshotStore.Load(ctx, matchID)
		if err != nil {
			return nil, err
		}
		snapshot = &SnapshotCache{
			Match:   match,
			Version: version,
			Stale:   false,
		}
		s.snapshots[matchID] = snapshot
	}
	return snapshot, nil
}

func (s *SnapshotService) createNewSnapshot(ctx context.Context, matchID domain.MatchID) (*SnapshotCache, error) {
	events, version, err := s.eventStore.LoadAll(ctx, matchID)
	if err != nil {
		return nil, err
	}
	match := (*domain.Match)(nil)
	for _, event := range events {
		match, err = event.Apply(match)
		if err != nil {
			return nil, err
		}
	}
	return &SnapshotCache{
		Match:   match,
		Version: version,
		Stale:   false,
	}, nil
}
