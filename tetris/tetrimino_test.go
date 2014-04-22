package tetris

import (
	"testing"
)

func TestShapeBoundaries(t *testing.T) {
	for _, shape := range Tetriminos {
		if shape.GetLeftmostColumn() != 0 {
			t.Errorf("Unexpected leftmost column for %s... %d", shape.Type, shape.GetLeftmostColumn())
		}

		if shape.GetRightmostCol() != len(shape.Shape[0])-1 {
			t.Errorf("Unexpected rightmost column for %s... %d", shape.Type, shape.GetRightmostCol())
		}
	}

	rotated_l := Tetriminos[0].Rotate()
	if rotated_l.GetLeftmostColumn() != 1 {
		t.Errorf("Unexpected leftmost column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
	if rotated_l.GetLowestRow() != 3 {
		t.Errorf("Unexpected lowest column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
	if rotated_l.GetRightmostCol() != 1 {
		t.Errorf("Unexpected right column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}

	rotated_l = rotated_l.Rotate()
	if rotated_l.GetLeftmostColumn() != 0 {
		t.Errorf("Unexpected leftmost column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
	if rotated_l.GetLowestRow() != 2 {
		t.Errorf("Unexpected lowest column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
	if rotated_l.GetRightmostCol() != 3 {
		t.Errorf("Unexpected right column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}

	rotated_l = rotated_l.Rotate()
	if rotated_l.GetLeftmostColumn() != 2 {
		t.Errorf("Unexpected leftmost column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
	if rotated_l.GetLowestRow() != 3 {
		t.Errorf("Unexpected lowest column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
	if rotated_l.GetRightmostCol() != 2 {
		t.Errorf("Unexpected right column for rotated l: %d", rotated_l.GetLeftmostColumn())
	}
}
