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
}

type EventMatchCreated struct {
	MatchID string
	Seats   []Seat
}

func (e EventMatchCreated) EventType() EventType {
	return EventTypeMatchCreated
}

type EventPlayerJoined struct {
	MatchID    string
	PlayerID   PlayerID
	SeatNumber int
}

func (e EventPlayerJoined) EventType() EventType {
	return EventTypePlayerJoined
}

type EventGameStarted struct {
	MatchID string
}

func (e EventGameStarted) EventType() EventType {
	return EventTypeGameStarted
}
