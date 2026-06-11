package serializable

// =======Match=======

type CreateMatchRequest struct {
	PlayerID string `json:"player_id"`
}

type CreateMatchResponse struct {
	Match JsonMatch `json:"match"`
}

type JoinMatchRequest struct {
	PlayerID string `json:"player_id"`
	MatchID  string `json:"match_id"`
}

type JoinMatchResponse struct {
	Match JsonMatch `json:"match"`
}

// =======Player=======

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

type CreatePlayerResponse struct {
	Name     string `json:"name"`
	PlayerID string `json:"player_id"`
}
