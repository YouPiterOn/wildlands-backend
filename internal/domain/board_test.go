package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestBoard_GenerateNewBoard_IdenticalBoards(t *testing.T) {
	seed, err := GenerateSeed()
	if err != nil {
		t.Fatalf("Failed to generate seed: %v", err)
	}
	board, err := GenerateNewBoard(seed, 0, PlayerID(uuid.New()))
	if err != nil {
		t.Fatalf("Failed to generate board: %v", err)
	}
	board2, err := GenerateNewBoard(seed, 0, PlayerID(uuid.New()))
	if err != nil {
		t.Fatalf("Failed to generate board: %v", err)
	}
	if board.String() != board2.String() {
		t.Fatalf("Generated boards are not identical")
	}
}
