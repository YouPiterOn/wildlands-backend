package scoringcards

import (
	"strings"
	"testing"

	"youpiteron.dev/wildlands-backend/internal/domain"
)

func TestScoringCardSentinelwood_ScoreBoard(t *testing.T) {
	card := NewScoringCardSentinelwood()

	t.Run("empty board", func(t *testing.T) {
		board := boardFromASCII(nil)
		if got := card.ScoreBoard(board); got != 0 {
			t.Errorf("ScoreBoard() = %d, want 0", got)
		}
	})

	t.Run("forest only in interior does not score", func(t *testing.T) {
		rows := make([]string, domain.BOARD_SIZE)
		for i := range rows {
			rows[i] = strings.Repeat(".", domain.BOARD_SIZE)
		}
		r := []rune(rows[5])
		r[5] = 'W'
		rows[5] = string(r)
		board := boardFromASCII(rows)
		if got := card.ScoreBoard(board); got != 0 {
			t.Errorf("ScoreBoard() = %d, want 0 (interior forest)", got)
		}
	})

	t.Run("one forest on top edge", func(t *testing.T) {
		board := boardFromASCII([]string{"W" + strings.Repeat(".", domain.BOARD_SIZE-1)})
		if got := card.ScoreBoard(board); got != 1 {
			t.Errorf("ScoreBoard() = %d, want 1", got)
		}
	})

	t.Run("one forest on bottom edge", func(t *testing.T) {
		rows := make([]string, domain.BOARD_SIZE)
		for i := range rows {
			rows[i] = strings.Repeat(".", domain.BOARD_SIZE)
		}
		rows[domain.BOARD_SIZE-1] = strings.Repeat(".", domain.BOARD_SIZE-1) + "W"
		board := boardFromASCII(rows)
		if got := card.ScoreBoard(board); got != 1 {
			t.Errorf("ScoreBoard() = %d, want 1", got)
		}
	})

	t.Run("one forest on left edge middle row", func(t *testing.T) {
		rows := make([]string, domain.BOARD_SIZE)
		for i := range rows {
			rows[i] = strings.Repeat(".", domain.BOARD_SIZE)
		}
		rows[5] = "W" + strings.Repeat(".", domain.BOARD_SIZE-1)
		board := boardFromASCII(rows)
		if got := card.ScoreBoard(board); got != 1 {
			t.Errorf("ScoreBoard() = %d, want 1", got)
		}
	})

	t.Run("one forest on right edge middle row", func(t *testing.T) {
		rows := make([]string, domain.BOARD_SIZE)
		for i := range rows {
			rows[i] = strings.Repeat(".", domain.BOARD_SIZE)
		}
		rows[5] = strings.Repeat(".", domain.BOARD_SIZE-1) + "W"
		board := boardFromASCII(rows)
		if got := card.ScoreBoard(board); got != 1 {
			t.Errorf("ScoreBoard() = %d, want 1", got)
		}
	})

	t.Run("full top row of forest", func(t *testing.T) {
		board := boardFromASCII([]string{strings.Repeat("W", domain.BOARD_SIZE)})
		want := domain.BOARD_SIZE
		if got := card.ScoreBoard(board); got != want {
			t.Errorf("ScoreBoard() = %d, want %d", got, want)
		}
	})

	t.Run("four corners each count once", func(t *testing.T) {
		rows := make([]string, domain.BOARD_SIZE)
		for i := range rows {
			rows[i] = strings.Repeat(".", domain.BOARD_SIZE)
		}
		r0 := []rune(rows[0])
		r0[0], r0[domain.BOARD_SIZE-1] = 'W', 'W'
		rows[0] = string(r0)
		rLast := []rune(rows[domain.BOARD_SIZE-1])
		rLast[0], rLast[domain.BOARD_SIZE-1] = 'W', 'W'
		rows[domain.BOARD_SIZE-1] = string(rLast)
		board := boardFromASCII(rows)
		if got := card.ScoreBoard(board); got != 4 {
			t.Errorf("ScoreBoard() = %d, want 4", got)
		}
	})

	t.Run("every edge cell forest", func(t *testing.T) {
		board := boardWithAllEdgeForest()
		// Top + bottom rows: 2*BOARD_SIZE; middle rows: 2*(BOARD_SIZE-2) for left+right
		want := 2*domain.BOARD_SIZE + 2*(domain.BOARD_SIZE-2)
		if got := card.ScoreBoard(board); got != want {
			t.Errorf("ScoreBoard() = %d, want %d (full perimeter)", got, want)
		}
	})
}

// boardWithAllEdgeForest sets TerrainForest on every cell on the board perimeter.
func boardWithAllEdgeForest() *domain.Board {
	b := domain.EmptyBoard(domain.PlayerID{})
	n := domain.BOARD_SIZE
	for x := range n {
		b.Cells[0][x].Terrain = domain.TerrainForest
		b.Cells[n-1][x].Terrain = domain.TerrainForest
	}
	for y := 1; y < n-1; y++ {
		b.Cells[y][0].Terrain = domain.TerrainForest
		b.Cells[y][n-1].Terrain = domain.TerrainForest
	}
	return b
}
