package snapshotstore

import "youpiteron.dev/wildlands-backend/internal/domain"

type StoredMatch struct {
	MatchID     string
	State       string
	SeatsCount  int
	CurrentTurn int
	Version     int
}

func (s StoredMatch) ToDomainMatch() (*domain.Match, int, error) {
	matchID, err := domain.ParseMatchID(s.MatchID)
	if err != nil {
		return nil, 0, err
	}
	state, err := domain.ParseMatchState(s.State)
	if err != nil {
		return nil, 0, err
	}
	return &domain.Match{
		ID:          matchID,
		State:       state,
		SeatsCount:  s.SeatsCount,
		CurrentTurn: s.CurrentTurn,
	}, s.Version, nil
}

func ToStoredMatch(match *domain.Match, version int) (StoredMatch, error) {
	return StoredMatch{
		MatchID:     match.ID.String(),
		State:       match.State.String(),
		SeatsCount:  match.SeatsCount,
		CurrentTurn: match.CurrentTurn,
		Version:     version,
	}, nil
}
