package domain

import (
	"testing"

	"youpiteron.dev/wildlands-backend/internal/utils"
)

func TestBoard_PlaceShape_Success(t *testing.T) {
	board := EmptyBoard(PlayerID{})
	placement := Placement{
		Shape:     ShapeSquare2(),
		Rotations: 0,
		Flipped:   false,
		Terrain:   TerrainForest,
		Position:  utils.NewVector2(1, 1),
	}
	ok := board.PlaceShape(placement)
	if !ok {
		t.Fatal("PlaceShape expected true, got false")
	}
	// Shape 2x2 at (1,1) covers cells (1,1), (2,1), (1,2), (2,2)
	for _, pt := range []struct{ y, x int }{{1, 1}, {2, 1}, {1, 2}, {2, 2}} {
		if board.Cells[pt.y][pt.x].Terrain != TerrainForest {
			t.Errorf("cell (%d,%d) expected Forest, got %v", pt.y, pt.x, board.Cells[pt.y][pt.x].Terrain)
		}
	}
}

func TestBoard_PlaceShape_Failure(t *testing.T) {
	board := EmptyBoard(PlayerID{})
	board.Cells[2][2].Terrain = TerrainMountain
	placement := Placement{
		Shape:     ShapeSquare2(),
		Rotations: 0,
		Flipped:   false,
		Terrain:   TerrainForest,
		Position:  utils.NewVector2(1, 1), // 2x2 covers (1,1),(2,1),(1,2),(2,2) — (2,2) is mountain
	}
	ok := board.PlaceShape(placement)
	if ok {
		t.Fatal("PlaceShape expected false (overlap with mountain), got true")
	}
	if board.Cells[2][2].Terrain != TerrainMountain {
		t.Error("PlaceShape must not mutate board on failure: (2,2) should still be Mountain")
	}
	// Other cells in the shape footprint should still be empty
	if board.Cells[1][1].Terrain != TerrainEmpty {
		t.Error("PlaceShape must not mutate board on failure: (1,1) should still be Empty")
	}
}

func TestBoard_GenerateNewBoard_IdenticalBoards(t *testing.T) {
	seed, err := GenerateSeed()
	if err != nil {
		t.Fatalf("Failed to generate seed: %v", err)
	}
	board, err := GenerateNewBoard(seed, 0, PlayerID{})
	if err != nil {
		t.Fatalf("Failed to generate board: %v", err)
	}
	board2, err := GenerateNewBoard(seed, 0, PlayerID{})
	if err != nil {
		t.Fatalf("Failed to generate board: %v", err)
	}
	if board.String() != board2.String() {
		t.Fatalf("Generated boards are not identical")
	}
}
