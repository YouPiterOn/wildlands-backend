package domain

type ExploreCardShapeOption struct {
	Shape *Shape
	Coins int
}

func NewExploreCardShapeOption(shape *Shape, coins int) ExploreCardShapeOption {
	return ExploreCardShapeOption{
		Shape: shape,
		Coins: coins,
	}
}

type ExploreCardName string

const (
	ExploreCardNameForgottenForest ExploreCardName = "FORGOTTEN_FOREST"
	ExploreCardNameForestHouses    ExploreCardName = "FOREST_HOUSES"
)

type ExploreCard struct {
	Name           ExploreCardName
	Duration       int
	TerrainOptions []Terrain
	ShapeOptions   []ExploreCardShapeOption
}

func NewExploreCard(name ExploreCardName, duration int, terrainOptions []Terrain, shapeOptions []ExploreCardShapeOption) ExploreCard {
	return ExploreCard{
		Name:           name,
		Duration:       duration,
		TerrainOptions: terrainOptions,
		ShapeOptions:   shapeOptions,
	}
}

var ExploreCards = []ExploreCard{
	NewExploreCard(
		ExploreCardNameForgottenForest,
		1,
		[]Terrain{TerrainForest},
		[]ExploreCardShapeOption{
			NewExploreCardShapeOption(ShapeDiagonal2x2(), 1),
			NewExploreCardShapeOption(ShapeZigzag2x3(), 0),
		},
	),
	NewExploreCard(
		ExploreCardNameForestHouses,
		2,
		[]Terrain{TerrainForest, TerrainVillage},
		[]ExploreCardShapeOption{
			NewExploreCardShapeOption(ShapeZigzag4x2(), 0),
		},
	),
}
