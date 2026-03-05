package domain

import "errors"

type Command interface {
	Handle(match *Match) ([]Event, error)
}

type CommandCreateMatch struct {
	MatchID    MatchID
	SeatsCount int
}

func (c CommandCreateMatch) Handle(match *Match) ([]Event, error) {
	if match != nil {
		return nil, errors.New("Match already created")
	}
	events := []Event{
		EventMatchCreated{
			MatchID:    c.MatchID,
			SeatsCount: c.SeatsCount,
		},
	}
	return events, nil
}

type CommandJoinMatch struct {
	MatchID  MatchID
	PlayerID PlayerID
}

func (c CommandJoinMatch) Handle(match *Match) ([]Event, error) {
	if match == nil {
		return nil, errors.New("Match does not exist")
	}
	if match.State != MatchStateCreated {
		return nil, errors.New("Match already started of ended")
	}
	if len(match.Seats) >= match.SeatsCount {
		return nil, errors.New("Match is already full")
	}
	events := []Event{
		EventPlayerJoined{
			PlayerID: c.PlayerID,
		},
	}
	return events, nil
}
