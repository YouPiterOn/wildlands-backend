package domain

type EventType int

const (
	EventTypeMatchCreated EventType = iota
	EventTypePlayerJoined
	EventTypeGameStarted
	EventTypeCardOpened
	EventTypeShapePlaced
	EventTypeTurnAdvanced
	EventTypeSeasonEnded
	EventTypeScoringApplied
	EventTypeGameFinished
)

func (e EventType) String() string {
	return []string{
		"MATCH_CREATED",
		"PLAYER_JOINED",
		"GAME_STARTED",
		"CARD_OPENED",
		"SHAPE_PLACED",
		"TURN_ADVANCED",
		"SEASON_ENDED",
		"SCORING_APPLIED",
		"GAME_FINISHED",
	}[e]
}

type Event interface {
	EventType() EventType
	Apply(match *Match) (*Match, error)
}

type EventMatchCreated struct {
	MatchID    string
	SeatsCount int
}

func (e EventMatchCreated) EventType() EventType {
	return EventTypeMatchCreated
}

func (e EventMatchCreated) Apply(match *Match) (*Match, error) {
	match.Seats = make([]Seat, e.SeatsCount)
	match.State = MatchStateCreated
	return match, nil
}

type EventPlayerJoined struct {
	MatchID    string
	PlayerID   PlayerID
	SeatNumber int
}

func (e EventPlayerJoined) EventType() EventType {
	return EventTypePlayerJoined
}

func (e EventPlayerJoined) Apply(match *Match) (*Match, error) {
	match.Seats[e.SeatNumber].PlayerID = e.PlayerID
	return match, nil
}

type EventGameStarted struct {
	MatchID string
}

func (e EventGameStarted) EventType() EventType {
	return EventTypeGameStarted
}

func (e EventGameStarted) Apply(match *Match) (*Match, error) {
	match.State = MatchStateStarted
	return match, nil
}
