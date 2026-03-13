package domain

import "youpiteron.dev/wildlands-backend/internal/utils"

var directions = []utils.Vector2{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

type ScoringCard interface {
	ScoreBoard(board *Board) int
}

type ScoringCardDeepwood struct{}

func (c ScoringCardDeepwood) ScoreBoard(board *Board) int {
	const MIN_FOREST_SIZE = 5
	const SCORE_PER_DEEPWOOD = 6

	var visited [11][11]bool
	score := 0

	for y := range BOARD_SIZE {
		for x := range BOARD_SIZE {

			if visited[y][x] {
				continue
			}

			if board.Cells[y][x].Terrain != TerrainForest {
				continue
			}

			size, valid := explore(board, utils.NewVector2(x, y), TerrainForest, &visited)

			if valid && size >= MIN_FOREST_SIZE {
				score += SCORE_PER_DEEPWOOD
			}
		}
	}

	return score
}

func explore(b *Board, position utils.Vector2, terrain Terrain, visited *[11][11]bool) (size int, valid bool) {
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

		if hasAdjacentTerrain(b, p, TerrainVillage) {
			valid = false
		}

		for _, d := range directions {
			nx := p.X + d.X
			ny := p.Y + d.Y

			if nx < 0 || ny < 0 || nx >= 11 || ny >= 11 {
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

func hasAdjacentTerrain(b *Board, position utils.Vector2, terrain Terrain) bool {
	for _, d := range directions {
		nx := position.X + d.X
		ny := position.Y + d.Y

		if nx < 0 || ny < 0 || nx >= 11 || ny >= 11 {
			continue
		}

		if b.Cells[ny][nx].Terrain == terrain {
			return true
		}
	}

	return false
}
