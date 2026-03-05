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
	MatchID    MatchID
	SeatsCount int
}

func (e EventMatchCreated) EventType() EventType {
	return EventTypeMatchCreated
}

func (e EventMatchCreated) Apply(match *Match) (*Match, error) {
	match.SeatsCount = e.SeatsCount
	match.State = MatchStateCreated
	return match, nil
}

type EventPlayerJoined struct {
	MatchID  MatchID
	PlayerID PlayerID
}

func (e EventPlayerJoined) EventType() EventType {
	return EventTypePlayerJoined
}

func (e EventPlayerJoined) Apply(match *Match) (*Match, error) {
	match.Seats = append(match.Seats, Seat{
		SeatNumber: len(match.Seats),
		PlayerID:   e.PlayerID,
		Score:      0,
	})
	return match, nil
}

type EventGameStarted struct {
	MatchID MatchID
}

func (e EventGameStarted) EventType() EventType {
	return EventTypeGameStarted
}

func (e EventGameStarted) Apply(match *Match) (*Match, error) {
	match.State = MatchStateStarted
	return match, nil
}
