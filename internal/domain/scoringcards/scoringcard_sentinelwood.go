package scoringcards

import "youpiteron.dev/wildlands-backend/internal/domain"

type ScoringCardSentinelwood struct{}

var _ domain.ScoringCard = (*ScoringCardSentinelwood)(nil)

func NewScoringCardSentinelwood() *ScoringCardSentinelwood {
	return &ScoringCardSentinelwood{}
}

func (c *ScoringCardSentinelwood) ScoreBoard(board *domain.Board) int {
	const SCORE_PER_SENTINELWOOD = 1

	score := 0

	for y := range domain.BOARD_SIZE {
		if y == 0 || y == domain.BOARD_SIZE-1 {
			for x := range domain.BOARD_SIZE {
				if board.Cells[y][x].Terrain == domain.TerrainForest {
					score += SCORE_PER_SENTINELWOOD
				}
			}
		} else {
			if board.Cells[y][0].Terrain == domain.TerrainForest {
				score += SCORE_PER_SENTINELWOOD
			}
			if board.Cells[y][domain.BOARD_SIZE-1].Terrain == domain.TerrainForest {
				score += SCORE_PER_SENTINELWOOD
			}
		}
	}

	return score
}
