package utils

import "math"

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func ChebyshevDistance(p1 Point, p2 Point) int {
	return int(math.Max(math.Abs(float64(p1.X-p2.X)), math.Abs(float64(p1.Y-p2.Y))))
}
