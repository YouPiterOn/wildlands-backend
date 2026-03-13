package scoringcards

import (
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

type ScoringCardDeepwood struct{}

var _ domain.ScoringCard = (*ScoringCardDeepwood)(nil)

func NewScoringCardDeepwood() *ScoringCardDeepwood {
	return &ScoringCardDeepwood{}
}

func (c *ScoringCardDeepwood) ScoreBoard(board *domain.Board) int {
	const MIN_FOREST_SIZE = 5
	const SCORE_PER_DEEPWOOD = 6

	var visited [domain.BOARD_SIZE][domain.BOARD_SIZE]bool
	score := 0

	for y := range domain.BOARD_SIZE {
		for x := range domain.BOARD_SIZE {

			if visited[y][x] {
				continue
			}

			if board.Cells[y][x].Terrain != domain.TerrainForest {
				continue
			}

			size, valid := exploreTerrain(board, utils.NewVector2(x, y), domain.TerrainForest, &visited)

			if valid && size >= MIN_FOREST_SIZE {
				score += SCORE_PER_DEEPWOOD
			}
		}
	}

	return score
}
