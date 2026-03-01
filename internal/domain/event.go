package domain

type EventType int

const (
	EventMatchCreated EventType = iota
	EventPlayerJoined
	EventGameStarted
	EventCardOpened
	EventShapePlaced
	EventTurnAdvanced
	EventSeasonEnded
	EventScoringApplied
	EventGameFinished
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
