package domain

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

type Match struct {
	ID          MatchID
	State       MatchState
	SeatsCount  int
	Seats       []Seat
	CurrentTurn int
}
