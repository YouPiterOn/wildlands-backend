package domain

type EventType int

const (
	EventTypePlayerJoined EventType = iota
	EventTypeGameStarted
	EventTypeExploreCardOpened
	EventTypeShapePlaced
	EventTypeTurnAdvanced
	EventTypeSeasonEnded
	EventTypeScoringApplied
	EventTypeGameFinished
)

func (e EventType) String() string {
	return []string{
		"PLAYER_JOINED",
		"GAME_STARTED",
		"EXPLORE_CARD_OPENED",
		"SHAPE_PLACED",
		"TURN_ADVANCED",
		"SEASON_ENDED",
		"SCORING_APPLIED",
		"GAME_FINISHED",
	}[e]
}

type Event interface {
	EventType() EventType
	apply(match *Match, metadata *MatchMetadata) (*Match, error)
}

// ================================================================
// Player Joined
// ================================================================
type EventPlayerJoined struct {
	MatchID  MatchID
	PlayerID PlayerID
}

func (e EventPlayerJoined) EventType() EventType {
	return EventTypePlayerJoined
}

func (e EventPlayerJoined) apply(match *Match, metadata *MatchMetadata) (*Match, error) {
	board, err := GenerateNewBoard(metadata.BoardsSeed, len(match.Boards), e.PlayerID)
	if err != nil {
		return nil, err
	}
	match.Boards = append(match.Boards, *board)
	return match, nil
}

// ================================================================
// Game Started
// ================================================================
type EventGameStarted struct {
	MatchID MatchID
}

func (e EventGameStarted) EventType() EventType {
	return EventTypeGameStarted
}

func (e EventGameStarted) apply(match *Match, _ *MatchMetadata) (*Match, error) {
	match.State = MatchStateStarted
	return match, nil
}

// ================================================================
// Card Opened
// ================================================================
type EventExploreCardOpened struct {
	MatchID MatchID
	CardID  ExploreCardID
}

func (e EventExploreCardOpened) EventType() EventType {
	return EventTypeExploreCardOpened
}
