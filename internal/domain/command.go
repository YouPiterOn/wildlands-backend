package domain

import "errors"

type Command interface {
	Handle(match *Match) ([]Event, error)
}

type CommandCreateMatch struct {
	MatchID MatchID
}

func (c CommandCreateMatch) Handle(match *Match) ([]Event, error) {
	if match != nil {
		return nil, errors.New("match already created")
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

func (c CommandJoinMatch) Handle(match *Match) ([]Event, error) {
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
