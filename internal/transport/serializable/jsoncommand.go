package serializable

import (
	"errors"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

type JsonCommand struct {
	Type     string `json:"type"`
	MatchID  string `json:"match_id,omitempty"`
	PlayerID string `json:"player_id,omitempty"`
}

func ToDomainCommand(command JsonCommand) (domain.Command, error) {
	switch domain.CommandType(command.Type) {

	case domain.CommandTypeJoinMatch:
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
