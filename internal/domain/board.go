package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"youpiteron.dev/wildlands-backend/internal/utils"
)

const BOARD_SIZE = 11

const MOUNTAIN_COUNT = 5
const MIN_DISTANCE_BETWEEN_MOUNTAINS = 3

const RUINS_COUNT = 5
const MIN_DISTANCE_BETWEEN_RUINS = 3

const MAX_TRY_COUNT = 1000

type Terrain int

const (
	TerrainEmpty Terrain = iota
	TerrainMountain
	TerrainMonster
	TerrainForest
	TerrainVillage
	TerrainFarm
	TerrainWater
)

func (t Terrain) Icon() string {
	return []string{
		"...",
		"Mnt",
		"Mns",
		"Fst",
		"Vlg",
		"Frm",
		"Wtr",
	}[t]
}

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

func (c *Cell) String() string {
	if c.HasRuins {
		return fmt.Sprintf("%sR", c.Terrain.Icon())
	}
	return fmt.Sprintf("%s ", c.Terrain.Icon())
}

type Placement struct {
	Shape     *Shape
	Rotations int
	Flipped   bool
	Terrain   Terrain
	Position  utils.Vector2
}

type Board struct {
	BoardNumber int      `json:"board_number"`
	PlayerID    PlayerID `json:"player_id"`
	Cells       [][]Cell `json:"cells"`
	Score       int      `json:"score"`
	Coins       int      `json:"coins"`
}

func GenerateNewBoard(seed int64, boardNumber int, playerID PlayerID) (*Board, error) {
	boardSeed := DeriveSeedFromInt(seed, boardNumber)
	rng := rand.New(rand.NewSource(boardSeed))
	cells := make([][]Cell, BOARD_SIZE)
	for i := range cells {
		cells[i] = make([]Cell, BOARD_SIZE)
		for j := range cells[i] {
			cells[i][j] = EmptyCell()
		}
	}
	err := placeMountains(cells, rng)
	if err != nil {
		return nil, err
	}
	err = placeRuins(cells, rng)
	if err != nil {
		return nil, err
	}
	return &Board{
		BoardNumber: boardNumber,
		PlayerID:    playerID,
		Cells:       cells,
		Score:       0,
		Coins:       0,
	}, nil
}

func RestoreBoard(boardNumber int, playerID PlayerID, cells [][]Cell, score int, coins int) *Board {
	return &Board{
		BoardNumber: boardNumber,
		PlayerID:    playerID,
		Cells:       cells,
		Score:       score,
		Coins:       coins,
	}
}

func EmptyBoard(playerID PlayerID) *Board {
	cells := make([][]Cell, BOARD_SIZE)
	for i := range cells {
		cells[i] = make([]Cell, BOARD_SIZE)
		for j := range cells[i] {
			cells[i][j] = EmptyCell()
		}
	}
	return &Board{
		BoardNumber: 0,
		PlayerID:    playerID,
		Cells:       cells,
		Score:       0,
		Coins:       0,
	}
}

func (b *Board) CanPlaceShape(placement Placement) bool {
	shape := placement.Shape.ToRotated(placement.Rotations)
	if placement.Flipped {
		shape = shape.ToFlipped()
	}

	p := placement.Position

	for i := range shape.Grid {
		for j := range shape.Grid[i] {
			if shape.Grid[i][j] {
				if b.Cells[p.Y+i][p.X+j].Terrain != TerrainEmpty {
					return false
				}
			}
		}
	}

	return true
}

func (b *Board) PlaceShape(placement Placement) bool {
	shape := placement.Shape.ToRotated(placement.Rotations)
	if placement.Flipped {
		shape = shape.ToFlipped()
	}

	p := placement.Position

	for i := range shape.Grid {
		for j := range shape.Grid[i] {
			if shape.Grid[i][j] {
				if b.Cells[p.Y+i][p.X+j].Terrain != TerrainEmpty {
					return false
				}
			}
		}
	}

	for i := range shape.Grid {
		for j := range shape.Grid[i] {
			if shape.Grid[i][j] {
				b.Cells[p.Y+i][p.X+j].Terrain = placement.Terrain
			}
		}
	}

	return true
}

func (b *Board) AddCoins(coins int) {
	b.Coins += coins
}

func (b *Board) ScoreSeason(scoringCards []ScoringCard) {
	for _, scoringCard := range scoringCards {
		b.Score += scoringCard.ScoreBoard(b)
	}
	b.Score += b.Coins
}

func (b *Board) String() string {
	builder := strings.Builder{}
	for _, row := range b.Cells {
		builder.WriteString("| ")
		for _, cell := range row {
			builder.WriteString(cell.String())
			builder.WriteString(" | ")
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func placeMountains(cells [][]Cell, rng *rand.Rand) error {
	mountains := []utils.Vector2{}
	tryCount := 0

	for len(mountains) < MOUNTAIN_COUNT && tryCount < MAX_TRY_COUNT {
		tryCount++
		x := rng.Intn(BOARD_SIZE)
		y := rng.Intn(BOARD_SIZE)

		if x == 0 || x == BOARD_SIZE-1 || y == 0 || y == BOARD_SIZE-1 {
			continue
		}

		p := utils.NewVector2(x, y)

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

	if tryCount >= MAX_TRY_COUNT {
		return errors.New("Failed to place mountains")
	}
	return nil
}

func placeRuins(cells [][]Cell, rng *rand.Rand) error {
	ruins := []utils.Vector2{}
	tryCount := 0

	for len(ruins) < RUINS_COUNT && tryCount < MAX_TRY_COUNT {
		tryCount++
		x := rng.Intn(BOARD_SIZE)
		y := rng.Intn(BOARD_SIZE)

		p := utils.NewVector2(x, y)

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

	if tryCount >= MAX_TRY_COUNT {
		return errors.New("Failed to place ruins")
	}
	return nil
}
