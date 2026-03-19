package scoringcards

import (
	"testing"
)

func TestScoringCardDeepwood_ScoreBoard(t *testing.T) {
	card := NewScoringCardDeepwood()

	t.Run("valid cluster of 5", func(t *testing.T) {
		board := boardFromASCII([]string{"WWWWW"})
		got := card.ScoreBoard(board)
		if got != 6 {
			t.Errorf("ScoreBoard() = %d, want 5", got)
		}
	})

	t.Run("cluster touching village", func(t *testing.T) {
		board := boardFromASCII([]string{
			"WWWWW",
			"..V..",
		})
		got := card.ScoreBoard(board)
		if got != 0 {
			t.Errorf("ScoreBoard() = %d, want 0 (cluster touches village)", got)
		}
	})

	t.Run("cluster smaller than 5", func(t *testing.T) {
		board := boardFromASCII([]string{"WWWW"})
		got := card.ScoreBoard(board)
		if got != 0 {
			t.Errorf("ScoreBoard() = %d, want 0 (cluster size < 5)", got)
		}
	})

	t.Run("two valid clusters", func(t *testing.T) {
		board := boardFromASCII([]string{"WWWWW.WWWWW"})
		got := card.ScoreBoard(board)
		if got != 12 {
			t.Errorf("ScoreBoard() = %d, want 10", got)
		}
	})

	t.Run("diagonals do not connect", func(t *testing.T) {
		board := boardFromASCII([]string{
			"W.W",
			".W.",
			"W.W",
		})
		got := card.ScoreBoard(board)
		if got != 0 {
			t.Errorf("ScoreBoard() = %d, want 0 (diagonals must not connect)", got)
		}
	})
}
