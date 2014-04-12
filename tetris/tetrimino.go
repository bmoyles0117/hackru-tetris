package tetris

import (
	"math/rand"
	"time"
)

const (
	ROTATION_0   = 0
	ROTATION_90  = 90
	ROTATION_180 = 180
	ROTATION_270 = 270

	TETRIMINO_I = 0
	TETRIMINO_J = 1
	TETRIMINO_L = 2
	TETRIMINO_O = 3
	TETRIMINO_S = 4
	TETRIMINO_T = 5
	TETRIMINO_Z = 6
)

type Tetrimino struct {
	Rotation int
	Type     int
	Shape    [][]byte
}

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomTetrimino() *Tetrimino {
	tetrimino := &Tetriminos[generator.Int()%len(Tetriminos)]
	tetrimino.Rotation = generator.Int() % 4 * 90

	return tetrimino
}

var Tetriminos = []Tetrimino{
	Tetrimino{
		Type: TETRIMINO_I,
		Shape: [][]byte{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	},
	Tetrimino{
		Type: TETRIMINO_J,
		Shape: [][]byte{
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		},
	},
	Tetrimino{
		Type: TETRIMINO_L,
		Shape: [][]byte{
			{1, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
	},
	Tetrimino{
		Type: TETRIMINO_O,
		Shape: [][]byte{
			{1, 1},
			{1, 1},
		},
	},
	Tetrimino{
		Type: TETRIMINO_S,
		Shape: [][]byte{
			{0, 1, 1},
			{1, 1, 0},
			{0, 0, 0},
		},
	},
	Tetrimino{
		Type: TETRIMINO_T,
		Shape: [][]byte{
			{0, 1, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
	},
	Tetrimino{
		Type: TETRIMINO_Z,
		Shape: [][]byte{
			{1, 1, 0},
			{0, 1, 1},
			{0, 0, 0},
		},
	},
}
