package domain

type Seat struct {
	SeatNumber int      `json:"seat_number"`
	PlayerID   PlayerID `json:"player_id"`
	Score      int      `json:"score"`
}
