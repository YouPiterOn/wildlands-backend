package snapshotstore

import (
	"encoding/json"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type StoredMatch struct {
	MatchID     string
	State       string
	CurrentTurn int
	SeatsJson   []byte
	Version     int
}

func (s StoredMatch) ToDomainMatch() (*domain.Match, error) {
	matchID, err := domain.ParseMatchID(s.MatchID)
	if err != nil {
		return nil, err
	}
	state, err := domain.ParseMatchState(s.State)
	if err != nil {
		return nil, err
	}
	seats := []domain.Seat{}
	err = json.Unmarshal(s.SeatsJson, &seats)
	if err != nil {
		return nil, err
	}
	return &domain.Match{
		ID:          matchID,
		State:       state,
		Seats:       seats,
		CurrentTurn: s.CurrentTurn,
		Version:     s.Version,
	}, nil
}

func ToStoredMatch(match *domain.Match) (StoredMatch, error) {
	seatsJson, err := json.Marshal(match.Seats)
	if err != nil {
		return StoredMatch{}, err
	}
	return StoredMatch{
		MatchID:     match.ID.String(),
		State:       match.State.String(),
		CurrentTurn: match.CurrentTurn,
		SeatsJson:   seatsJson,
		Version:     match.Version,
	}, nil
}
