package scoringcards

import (
	"youpiteron.dev/wildlands-backend/internal/domain"
	"youpiteron.dev/wildlands-backend/internal/utils"
)

var directions = []utils.Vector2{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

func exploreTerrain(b *domain.Board, position utils.Vector2, terrain domain.Terrain, visited *[domain.BOARD_SIZE][domain.BOARD_SIZE]bool) (size int, valid bool) {
	stack := []utils.Vector2{position}
	valid = true

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[p.Y][p.X] {
			continue
		}

		visited[p.Y][p.X] = true
		size++

		if hasAdjacentTerrain(b, p, domain.TerrainVillage) {
			valid = false
		}

		for _, d := range directions {
			nx := p.X + d.X
			ny := p.Y + d.Y

			if nx < 0 || ny < 0 || nx >= domain.BOARD_SIZE || ny >= domain.BOARD_SIZE {
				continue
			}

			if visited[ny][nx] {
				continue
			}

			if b.Cells[ny][nx].Terrain == terrain {
				stack = append(stack, utils.NewVector2(nx, ny))
			}
		}
	}

	return
}

func hasAdjacentTerrain(b *domain.Board, position utils.Vector2, terrain domain.Terrain) bool {
	for _, d := range directions {
		nx := position.X + d.X
		ny := position.Y + d.Y

		if nx < 0 || ny < 0 || nx >= domain.BOARD_SIZE || ny >= domain.BOARD_SIZE {
			continue
		}

		if b.Cells[ny][nx].Terrain == terrain {
			return true
		}
	}

	return false
}
