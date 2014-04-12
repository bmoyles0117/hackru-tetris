package tetris

import (
	"testing"
)

func TestGenerateTetrimino(t *testing.T) {
	var (
		to *Tetrimino
	)

	to = GenerateRandomTetrimino()

	if to.Rotation%90 != 0 {
		t.Errorf("Unexpected rotation: %d", to.Rotation)
	}
}
