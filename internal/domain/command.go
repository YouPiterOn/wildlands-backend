package domain

import "errors"

type Command interface {
	GetMatchID() MatchID
	Handle(match *Match) ([]Event, error)
}

type CommandCreateMatch struct {
	MatchID MatchID
}

func NewCommandCreateMatch(matchID MatchID) Command {
	return CommandCreateMatch{
		MatchID: matchID,
	}
}

func (c CommandCreateMatch) GetMatchID() MatchID {
	return c.MatchID
}

func (c CommandCreateMatch) Handle(match *Match) ([]Event, error) {
	if match != nil {
		return nil, errors.New("Match already created")
	}
	events := []Event{
		EventMatchCreated{
			MatchID: c.MatchID,
		},
	}
	return events, nil
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
