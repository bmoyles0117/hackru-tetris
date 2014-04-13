package tetris

import (
	"fmt"
)

type Grid struct {
	cells [][]uint8
}

func (g *Grid) clearFilledLines() {
	new_cells := make([][]uint8, len(g.cells))

	for row := 0; row < len(g.cells); row++ {
		new_cells[row] = make([]uint8, len(g.cells[row]))
	}

	new_row_index := len(new_cells) - 1

	for row := len(g.cells) - 1; row >= 0; row-- {
		is_complete := true

		for col := 0; col < len(g.cells[row]); col++ {
			if g.cells[row][col] == 0 {
				is_complete = false
				break
			}
		}

		if !is_complete {
			copy(new_cells[new_row_index], g.cells[row])

			new_row_index -= 1
		}
	}

	g.cells = new_cells
}

func (g *Grid) consumeTetrimino(row, col int, t *Tetrimino) {
	for shaperow := range t.Shape {
		for shapecol := range t.Shape[shaperow] {
			if 0 <= col+shapecol && col+shapecol < len(g.cells[0]) && row+shaperow < len(g.cells) {
				if t.Shape[shaperow][shapecol] == 1 {
					g.cells[row+shaperow][col+shapecol] = t.Type
				}
			}
		}
	}
}

func (g *Grid) tetriminoCausesCollision(row, col int, t *Tetrimino) bool {
	if col+t.GetRightmostCol() >= len(g.cells[0]) {
		fmt.Println("right barrier")
		return true
	}

	if col+t.GetLeftmostColumn() < 0 {
		fmt.Println("left barrier")

		return true
	}

	if row+t.GetLowestRow() >= len(g.cells) {
		fmt.Println("low barrier", len(g.cells), row, t.GetLowestRow())

		return true
	}

	for y := range t.Shape {
		for x := range t.Shape[y] {
			if t.Shape[y][x] == 0 {
				continue
			}

			if col+y >= len(g.cells) {
				continue
			}

			if row+x >= len(g.cells) {
				continue
			}

			if g.cells[row+y][col+x] != 0 {
				fmt.Println("piece conflict barrier")

				return true
			}
		}
	}

	return false
}

func newGrid(rows, cols int) *Grid {
	cells := make([][]uint8, rows)

	for i := range cells {
		cells[i] = make([]uint8, cols)
	}

	return &Grid{cells: cells}
}
