package domain

import "errors"

type MatchID string

func (id MatchID) String() string {
	return string(id)
}

func ParseMatchID(s string) (MatchID, error) {
	return MatchID(s), nil
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
	ID          MatchID
	State       MatchState
	SeatsCount  int
	Seats       []Seat
	CurrentTurn int
	Version     int
}

func NewMatch(id MatchID) *Match {
	return &Match{
		ID:          id,
		State:       MatchStateCreated,
		SeatsCount:  0,
		Seats:       []Seat{},
		CurrentTurn: 0,
		Version:     0,
	}
}

func (m *Match) Clone() *Match {
	return &Match{
		ID:          m.ID,
		State:       m.State,
		SeatsCount:  m.SeatsCount,
		Seats:       m.Seats,
		CurrentTurn: m.CurrentTurn,
		Version:     m.Version,
	}
}

func (m *Match) Apply(event Event) error {
	_, err := event.apply(m)
	if err != nil {
		return err
	}
	m.Version++
	return nil
}
