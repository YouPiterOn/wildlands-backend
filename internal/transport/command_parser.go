package transport

import (
	"errors"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

func ParseCommand(command JsonCommand) (domain.Command, error) {
	switch command.Type {

	case "create_match":
		id, err := domain.ParseMatchID(command.MatchID)
		if err != nil {
			return nil, err
		}
		return domain.CommandCreateMatch{
			MatchID: id,
		}, nil

	case "join_match":
		matchID, err := domain.ParseMatchID(command.MatchID)
		if err != nil {
			return nil, err
		}

		playerID, err := domain.ParsePlayerID(command.PlayerID)
		if err != nil {
			return nil, err
		}

		return domain.CommandJoinMatch{
			MatchID:  matchID,
			PlayerID: playerID,
		}, nil

	default:
		return nil, errors.New("unknown command")
	}
}
