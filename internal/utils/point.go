package utils

import "math"

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x int, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

func ChebyshevDistance(v1 Vector2, v2 Vector2) int {
	return int(math.Max(math.Abs(float64(v1.X-v2.X)), math.Abs(float64(v1.Y-v2.Y))))
}
