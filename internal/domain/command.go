package domain

import "errors"

type CommandType string

const (
	CommandTypeJoinMatch CommandType = "JOIN_MATCH"
)

func (c CommandType) String() string {
	return string(c)
}

type Command interface {
	Type() CommandType
	GetMatchID() MatchID
	Handle(match *Match) ([]Event, error)
}

type CommandCreateMatch struct {
	MatchID MatchID
}

type CommandJoinMatch struct {
	MatchID  MatchID
	PlayerID PlayerID
}

func NewCommandJoinMatch(matchID MatchID, playerID PlayerID) Command {
	return CommandJoinMatch{
		MatchID:  matchID,
		PlayerID: playerID,
	}
}

func (c CommandJoinMatch) Type() CommandType {
	return CommandTypeJoinMatch
}

func (c CommandJoinMatch) GetMatchID() MatchID {
	return c.MatchID
}

func (c CommandJoinMatch) Handle(match *Match) ([]Event, error) {
	if match == nil {
		return nil, errors.New("Match does not exist")
	}
	if match.State != MatchStateCreated {
		return nil, errors.New("Match already started of ended")
	}
	events := []Event{
		EventPlayerJoined{
			PlayerID: c.PlayerID,
		},
	}
	return events, nil
}
