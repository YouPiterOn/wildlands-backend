package domain

import (
	"testing"
)

func shapeEqual(a, b *Shape) bool {
	if len(a.Grid) != len(b.Grid) {
		return false
	}
	for i := range a.Grid {
		if len(a.Grid[i]) != len(b.Grid[i]) {
			return false
		}
		for j := range a.Grid[i] {
			if a.Grid[i][j] != b.Grid[i][j] {
				return false
			}
		}
	}
	return true
}

func TestShape_ToRotated_ZeroRotations(t *testing.T) {
	s := ShapeDiagonal3()
	got := s.ToRotated(0)
	if !shapeEqual(s, got) {
		t.Errorf("ToRotated(0) should return identical shape\ngot:\n%v\nwant:\n%v", got.String(), s.String())
	}
}

func TestShape_ToRotated_OneRotation(t *testing.T) {
	// Diagonal \ becomes diagonal / (90° clockwise)
	// 1 0 0    0 0 1
	// 0 1 0 -> 0 1 0
	// 0 0 1    1 0 0
	orig := ShapeDiagonal3()
	got := orig.ToRotated(1)
	want := NewShape([][]bool{
		{false, false, true},
		{false, true, false},
		{true, false, false},
	})
	if !shapeEqual(got, want) {
		t.Errorf("ToRotated(1) mismatch\ngot:\n%v\nwant:\n%v", got.String(), want.String())
	}
}

func TestShape_ToRotated_TwoRotations(t *testing.T) {
	// Two 90° rotations = 180°; diagonal is symmetric so back to original
	orig := ShapeDiagonal3()
	got := orig.ToRotated(2)
	if !shapeEqual(got, orig) {
		t.Errorf("ToRotated(2) should equal original for diagonal\ngot:\n%v\nwant:\n%v", got.String(), orig.String())
	}
}

func TestShape_ToRotated_FourRotations(t *testing.T) {
	orig := ShapeDiagonal3()
	got := orig.ToRotated(4)
	if !shapeEqual(got, orig) {
		t.Errorf("ToRotated(4) should equal original\ngot:\n%v\nwant:\n%v", got.String(), orig.String())
	}
}

func TestShape_ToRotated_Modulo(t *testing.T) {
	orig := ShapeDiagonal3()
	// 5 % 4 == 1, so same as one rotation
	got := orig.ToRotated(5)
	want := orig.ToRotated(1)
	if !shapeEqual(got, want) {
		t.Errorf("ToRotated(5) should equal ToRotated(1)\ngot:\n%v\nwant:\n%v", got.String(), want.String())
	}
}

func TestShape_ToRotated_EmptyGrid(t *testing.T) {
	s := NewShape([][]bool{})
	got := s.ToRotated(1)
	if len(got.Grid) != 0 {
		t.Errorf("ToRotated on empty grid should return empty grid, got %d rows", len(got.Grid))
	}
}

func TestShape_ToFlipped(t *testing.T) {
	// 1 0 0  ->  0 0 1  (each row reversed)
	// 0 1 0  ->  0 1 0
	// 0 0 1  ->  1 0 0
	orig := ShapeDiagonal3()
	got := orig.ToFlipped()
	want := NewShape([][]bool{
		{false, false, true},
		{false, true, false},
		{true, false, false},
	})
	if !shapeEqual(got, want) {
		t.Errorf("ToFlipped() mismatch\ngot:\n%v\nwant:\n%v", got.String(), want.String())
	}
}

func TestShape_ToFlipped_TwiceRestoresOriginal(t *testing.T) {
	orig := ShapeDiagonal3()
	flipped := orig.ToFlipped()
	got := flipped.ToFlipped()
	if !shapeEqual(got, orig) {
		t.Errorf("flip twice should restore original\ngot:\n%v\nwant:\n%v", got.String(), orig.String())
	}
}

func TestShape_ToFlipped_DoesNotMutateOriginal(t *testing.T) {
	orig := ShapeDiagonal3()
	origStr := orig.String()
	_ = orig.ToFlipped()
	if orig.String() != origStr {
		t.Error("ToFlipped() must not mutate the original shape")
	}
}

func TestShape_ToRotated_DoesNotMutateOriginal(t *testing.T) {
	orig := ShapeDiagonal3()
	origStr := orig.String()
	_ = orig.ToRotated(1)
	if orig.String() != origStr {
		t.Error("ToRotated() must not mutate the original shape")
	}
}

// ShapeSquare2 is symmetric under rotation and flip; we still assert no panic and clone behavior
func TestShape_ToRotated_Square2(t *testing.T) {
	sq := ShapeSquare2()
	for _, r := range []int{0, 1, 2, 3} {
		got := sq.ToRotated(r)
		if len(got.Grid) != 2 || len(got.Grid[0]) != 2 {
			t.Errorf("ShapeSquare2 ToRotated(%d): expected 2x2, got %dx%d", r, len(got.Grid), len(got.Grid[0]))
		}
		if !shapeEqual(got, sq) {
			t.Errorf("ShapeSquare2 ToRotated(%d) should equal original (square is symmetric)", r)
		}
	}
}

func TestShape_ToFlipped_Square2(t *testing.T) {
	sq := ShapeSquare2()
	got := sq.ToFlipped()
	if !shapeEqual(got, sq) {
		t.Errorf("ShapeSquare2 ToFlipped() should equal original (square is symmetric)\ngot:\n%v", got.String())
	}
}
