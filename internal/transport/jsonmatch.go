package transport

import "youpiteron.dev/wildlands-backend/internal/domain"

type JsonMatch struct {
	MatchID     string      `json:"match_id"`
	Boards      []JsonBoard `json:"boards"`
	CurrentTurn int         `json:"current_turn"`
	Version     int         `json:"version"`
}

func ToJsonMatch(match *domain.Match) JsonMatch {
	boards := make([]JsonBoard, len(match.Boards))
	for i, board := range match.Boards {
		boards[i] = ToJsonBoard(&board)
	}
	return JsonMatch{
		MatchID:     match.ID.String(),
		Boards:      boards,
		CurrentTurn: match.CurrentTurn,
		Version:     match.Version,
	}
}
