package domain

import (
	"errors"

	"github.com/google/uuid"
)

type MatchID uuid.UUID

func (id MatchID) String() string {
	return uuid.UUID(id).String()
}

func ParseMatchID(s string) (MatchID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return MatchID{}, err
	}
	return MatchID(id), nil
}

func NewMatchID() (MatchID, error) {
	matchId, err := uuid.NewRandom()
	if err != nil {
		return MatchID{}, err
	}
	return MatchID(matchId), nil
}

type MatchState int

const (
	MatchStateCreated MatchState = iota
	MatchStateStarted
	MatchStateFinished
)

func (s MatchState) String() string {
	return []string{
		"CREATED",
		"STARTED",
		"FINISHED",
	}[s]
}

func ParseMatchState(s string) (MatchState, error) {
	switch s {
	case "CREATED":
		return MatchStateCreated, nil
	case "STARTED":
		return MatchStateStarted, nil
	case "FINISHED":
		return MatchStateFinished, nil
	}
	return 0, errors.New("Invalid match state")
}

type Match struct {
	ID                 MatchID
	State              MatchState
	Boards             []Board
	CurrentExploreCard *ExploreDeckCard
	CurrentTurn        int
	Version            int
}

func NewMatch(id MatchID) *Match {
	return &Match{
		ID:                 id,
		State:              MatchStateCreated,
		Boards:             []Board{},
		CurrentExploreCard: nil,
		CurrentTurn:        0,
		Version:            0,
	}
}

func (m *Match) Clone() *Match {
	return &Match{
		ID:                 m.ID,
		State:              m.State,
		Boards:             m.Boards,
		CurrentExploreCard: m.CurrentExploreCard,
		CurrentTurn:        m.CurrentTurn,
		Version:            m.Version,
	}
}

func (m *Match) Apply(event Event, metadata *MatchMetadata) error {
	_, err := event.apply(m, metadata)
	if err != nil {
		return err
	}
	m.Version++
	return nil
}
