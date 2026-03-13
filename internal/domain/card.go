package domain

import "github.com/google/uuid"

type CardOption struct {
	Shape   *Shape
	Terrain Terrain
	Coins   int
}

func NewCardOption(shape *Shape, terrain Terrain, coins int) CardOption {
	return CardOption{
		Shape:   shape,
		Terrain: terrain,
		Coins:   coins,
	}
}

type CardID uuid.UUID

type Card struct {
	ID      CardID
	Options []CardOption
}

func NewCard(id CardID, options []CardOption) *Card {
	return &Card{
		ID:      id,
		Options: options,
	}
}
