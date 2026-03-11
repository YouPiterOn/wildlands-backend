package transport

type JsonCommand struct {
	Type     string `json:"type"`
	MatchID  string `json:"match_id,omitempty"`
	PlayerID string `json:"player_id,omitempty"`
}
