package tetris

import (
	// "fmt"
	"fmt"
	"math/rand"
	"time"
)

const (
	ROTATION_0   = 0
	ROTATION_90  = 90
	ROTATION_180 = 180
	ROTATION_270 = 270

	TETRIMINO_I = 'I'
	TETRIMINO_J = 'J'
	TETRIMINO_L = 'L'
	TETRIMINO_O = 'O'
	TETRIMINO_S = 'S'
	TETRIMINO_T = 'T'
	TETRIMINO_Z = 'Z'
)

type Tetrimino struct {
	Type  uint8
	Shape [][]byte
}

func (t *Tetrimino) GetLeftmostColumn() int {
	min_col := len(t.Shape[0]) - 1

	for row := range t.Shape {
		for col := range t.Shape[row] {
			if t.Shape[row][col] == 1 && col < min_col {
				min_col = col
			}
		}
	}
	fmt.Println("%i----%i", t.Shape, min_col)

	return min_col
}

func (t *Tetrimino) GetRightmostCol() int {
	max_col := 0

	for row := range t.Shape {
		for col := range t.Shape[row] {
			if t.Shape[row][col] == 1 && col > max_col {
				max_col = col
			}
		}
	}

	return max_col
}

func (t *Tetrimino) GetLowestRow() int {
	max_row := 0

	for row := range t.Shape {
		for col := range t.Shape[row] {
			if t.Shape[row][col] == 1 && row > max_row {
				max_row = row
			}
		}
	}

	return max_row
}

func (t *Tetrimino) Rotate() *Tetrimino {
	size := int(len(t.Shape[0]))
	transposed := createNewShape(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			transposed[size-i-1][j] = t.Shape[j][i]
		}
	}

	for row := range t.Shape {
		for col := range t.Shape[row] {
			t.Shape[row][col] = transposed[row][col]
		}
	}

	return &Tetrimino{
		Type:  t.Type,
		Shape: transposed,
	}

}

func createNewShape(size int) [][]byte {
	shape := make([][]byte, size)
	for i := range shape {
		shape[i] = make([]byte, size)
	}
	return shape
}

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomTetrimino() *Tetrimino {
	// return &Tetriminos[0]
	return &Tetriminos[generator.Int()%len(Tetriminos)]
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
