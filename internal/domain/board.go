package domain

import (
	"math/rand"

	"youpiteron.dev/wildlands-backend/internal/utils"
)

const BOARD_SIZE = 11

const MOUNTAIN_COUNT = 5
const MIN_DISTANCE_BETWEEN_MOUNTAINS = 3

const RUINS_COUNT = 5
const MIN_DISTANCE_BETWEEN_RUINS = 3

type Terrain int

const (
	TerrainEmpty Terrain = iota
	TerrainMountain
	TerrainMonsters
	TerrainForest
	TerrainWater
)

type Cell struct {
	Terrain  Terrain
	HasRuins bool
}

func EmptyCell() Cell {
	return Cell{
		Terrain:  TerrainEmpty,
		HasRuins: false,
	}
}

type Board struct {
	BoardNumber int
	PlayerID    PlayerID
	Cells       [][]Cell
	Score       int
}

func GenerateBoard(seed int64, boardNumber int, playerID PlayerID) *Board {
	boardSeed := DeriveSeedFromInt(seed, boardNumber)
	rng := rand.New(rand.NewSource(boardSeed))
	cells := make([][]Cell, BOARD_SIZE)
	for i := range cells {
		cells[i] = make([]Cell, BOARD_SIZE)
		for j := range cells[i] {
			cells[i][j] = EmptyCell()
		}
	}
	placeMountains(cells, rng)
	return &Board{
		BoardNumber: boardNumber,
		PlayerID:    playerID,
		Cells:       cells,
		Score:       0,
	}
}

func placeMountains(cells [][]Cell, rng *rand.Rand) {
	mountains := []utils.Point{}

	for len(mountains) < MOUNTAIN_COUNT {
		x := rng.Intn(BOARD_SIZE)
		y := rng.Intn(BOARD_SIZE)

		p := utils.NewPoint(x, y)

		if cells[y][x].Terrain != TerrainEmpty || cells[y][x].HasRuins {
			continue
		}

		valid := true

		for _, m := range mountains {
			if utils.ChebyshevDistance(p, m) < MIN_DISTANCE_BETWEEN_MOUNTAINS {
				valid = false
				break
			}
		}

		if !valid {
			continue
		}

		cells[y][x].Terrain = TerrainMountain
		mountains = append(mountains, p)
	}
}

func placeRuins(cells [][]Cell, rng *rand.Rand) {
	ruins := []utils.Point{}

	for len(ruins) < RUINS_COUNT {
		x := rng.Intn(BOARD_SIZE)
		y := rng.Intn(BOARD_SIZE)

		p := utils.NewPoint(x, y)

		if cells[y][x].Terrain == TerrainMountain || cells[y][x].HasRuins {
			continue
		}

		valid := true

		for _, r := range ruins {
			if utils.ChebyshevDistance(p, r) < MIN_DISTANCE_BETWEEN_RUINS {
				valid = false
				break
			}
		}

		if !valid {
			continue
		}

		cells[y][x].HasRuins = true
		ruins = append(ruins, p)
	}
}
