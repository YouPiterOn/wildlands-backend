package domain

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
	ID          string
	State       MatchState
	Seats       []Seat
	CurrentTurn int
}
