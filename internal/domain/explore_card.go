package domain

import "github.com/google/uuid"

type ExploreCardOption struct {
	Shape   *Shape
	Terrain Terrain
	Coins   int
}

func NewExploreCardOption(shape *Shape, terrain Terrain, coins int) ExploreCardOption {
	return ExploreCardOption{
		Shape:   shape,
		Terrain: terrain,
		Coins:   coins,
	}
}

type ExploreCardID uuid.UUID

type ExploreCard struct {
	ID      ExploreCardID
	Options []ExploreCardOption
	IsRuins bool
}

func NewExploreCard(id ExploreCardID, options []ExploreCardOption) *ExploreCard {
	return &ExploreCard{
		ID:      id,
		Options: options,
		IsRuins: false,
	}
}

func NewRuinsExploreCard(id ExploreCardID) *ExploreCard {
	return &ExploreCard{
		ID:      id,
		Options: []ExploreCardOption{},
		IsRuins: true,
	}
}
