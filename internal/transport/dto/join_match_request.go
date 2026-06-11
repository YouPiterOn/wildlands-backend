package dto

type JoinMatchRequest struct {
	PlayerID string `json:"player_id"`
	MatchID  string `json:"match_id"`
}
