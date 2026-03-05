package eventstore

type MatchCreatedEventData struct {
	SeatsCount int `json:"seats_count"`
}

type PlayerJoinedEventData struct {
	PlayerID string `json:"player_id"`
}
