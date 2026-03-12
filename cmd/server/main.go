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
	board := domain.GenerateBoard(seed, 0, domain.PlayerID(uuid.New()))
	fmt.Println(board.String())
}
