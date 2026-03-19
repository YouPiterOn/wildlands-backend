package scoringcards

import "youpiteron.dev/wildlands-backend/internal/domain"

func boardFromASCII(rows []string) *domain.Board {
	b := domain.EmptyBoard(domain.PlayerID{})
	for y, row := range rows {
		if y >= domain.BOARD_SIZE {
			break
		}
		for x, ch := range row {
			if x >= domain.BOARD_SIZE {
				break
			}
			switch ch {
			case 'W':
				b.Cells[y][x].Terrain = domain.TerrainForest
			case 'V':
				b.Cells[y][x].Terrain = domain.TerrainVillage
			case '.':
				b.Cells[y][x].Terrain = domain.TerrainEmpty
			}
		}
	}
	return b
}
