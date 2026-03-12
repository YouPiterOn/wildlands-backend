package domain

import "strings"

type Shape struct {
	Grid [][]bool
}

func NewShape(grid [][]bool) *Shape {
	return &Shape{Grid: grid}
}

func (s *Shape) ToRotated(rotations int) *Shape {
	rotations = rotations % 4
	if len(s.Grid) == 0 {
		return NewShape(make([][]bool, 0))
	}
	shape := s.Clone()
	for i := 0; i < rotations; i++ {
		shape = shape.toRotated90()
	}
	return shape
}

func (s *Shape) ToFlipped() *Shape {
	newGrid := make([][]bool, len(s.Grid))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(s.Grid[i]))
		for j := range s.Grid[i] {
			newGrid[i][len(s.Grid[i])-1-j] = s.Grid[i][j]
		}
	}
	return NewShape(newGrid)
}

func (s *Shape) Clone() *Shape {
	newGrid := make([][]bool, len(s.Grid))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(s.Grid[i]))
		copy(newGrid[i], s.Grid[i])
	}
	return NewShape(newGrid)
}

func (s *Shape) toRotated90() *Shape {
	newGrid := make([][]bool, len(s.Grid[0]))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(s.Grid))
	}
	for i := range s.Grid {
		for j := range s.Grid[i] {
			newGrid[j][len(s.Grid)-i-1] = s.Grid[i][j]
		}
	}
	return NewShape(newGrid)
}

func (s *Shape) String() string {
	builder := strings.Builder{}
	for _, row := range s.Grid {
		for _, cell := range row {
			if cell {
				builder.WriteString("1")
			} else {
				builder.WriteString("0")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func ShapeSquare2() *Shape {
	return NewShape([][]bool{
		{true, true},
		{true, true},
	})
}

func ShapeDiagonal3() *Shape {
	return NewShape([][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, true},
	})
}
