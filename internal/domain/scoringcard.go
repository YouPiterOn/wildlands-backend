package domain

type ScoringCard interface {
	ScoreBoard(board *Board) int
}
