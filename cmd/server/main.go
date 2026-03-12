package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"youpiteron.dev/wildlands-backend/internal/domain"
)

func main() {
	seed, err := domain.GenerateSeed()
	if err != nil {
		log.Fatalf("Failed to generate seed: %v", err)
	}
	board, err := domain.GenerateNewBoard(seed, 0, domain.PlayerID(uuid.New()))
	if err != nil {
		log.Fatalf("Failed to generate board: %v", err)
		return
	}
	fmt.Println(board.String())
}
