package eventstore

type MatchCreatedEventData struct {
	SeatsCount int `json:"seats_count"`
}

type PlayerJoinedEventData struct {
	SeatNumber int    `json:"seat_number"`
	PlayerID   string `json:"player_id"`
}
