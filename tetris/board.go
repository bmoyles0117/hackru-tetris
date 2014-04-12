package tetris

import (
	"fmt"
	"time"
)

const (
	DIRECTION_LEFT  = 'L'
	DIRECTION_UP    = 'U'
	DIRECTION_RIGHT = 'R'
	DIRECTION_DOWN  = 'D'
)

type BoardTetrimino struct {
	col       int
	row       int
	Tetrimino *Tetrimino
}

type Grid struct {
	cells [][]bool
}

func (g *Grid) consumeTetrimino(row, col int, t *Tetrimino) {
	fmt.Println("INSERT INTO ROW", row, "COL", col)
	for shaperow := range t.Shape {
		for shapecol := range t.Shape[shaperow] {
			if 0 <= col+shapecol && col+shapecol < len(g.cells[0]) && row+shaperow < len(g.cells) {
				if t.Shape[shaperow][shapecol] == 1 {
					g.cells[row+shaperow][col+shapecol] = true
				}
			}
		}
	}
}

func (g *Grid) tetriminoCausesCollision(row, col int, t *Tetrimino) bool {
	if col+t.GetRightmostCol() >= len(g.cells[0]) {
		return true
	}

	if col+t.GetLeftmostColumn() < 0 {
		return true
	}

	if row+t.GetLowestRow() >= len(g.cells) {
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

			fmt.Println("ROW", row, "X", x, "+", row+x)

			if row+x >= len(g.cells) {
				fmt.Println("CONT")
				continue
			}

			if g.cells[row+y][col+x] == true {
				return true
			}
		}
	}

	return false
}

func (g *Grid) print(current *BoardTetrimino) {
	tcells := make([][]bool, len(g.cells))

	for y := range g.cells {
		tcells[y] = make([]bool, len(g.cells[y]))
		copy(tcells[y], g.cells[y])
	}

	for y := range current.Tetrimino.Shape {
		for x := range current.Tetrimino.Shape[y] {
			if y+current.row < len(g.cells) && current.Tetrimino.Shape[y][x] == 1 {
				tcells[y+current.row][x+current.col] = current.Tetrimino.Shape[y][x] == 1
			}
		}
	}

	for row := range tcells {
		for col := range tcells[row] {
			if tcells[row][col] {
				fmt.Print(" X ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("------------------------------------")
}

func newGrid(rows, cols int) *Grid {
	cells := make([][]bool, rows)

	for i := range cells {
		cells[i] = make([]bool, cols)
	}

	return &Grid{cells: cells}
}

type Board struct {
	Rows    int
	Columns int

	grid       *Grid
	current    *BoardTetrimino
	next       *BoardTetrimino
	tetriminos []*BoardTetrimino
	moves      chan uint8
	running    bool
}

func (b *Board) AddTetrimino() {
	if b.next != nil {
		// Stash the current piece in the list of pieces
		b.tetriminos = append(b.tetriminos, b.current)

		b.grid.consumeTetrimino(b.current.row, b.current.col, b.current.Tetrimino)

		b.current = b.next
	} else {
		b.current = b.generateRandomTetrimino()
	}

	b.next = b.generateRandomTetrimino()
}

func (b *Board) Move(move_direction uint8) {
	b.moves <- move_direction
}

func (b *Board) move(move_direction uint8) {
	switch move_direction {
	case DIRECTION_LEFT:
		if !b.grid.tetriminoCausesCollision(b.current.row, b.current.col-1, b.current.Tetrimino) {
			b.current.col -= 1
		}

	case DIRECTION_UP:
		// Add counter clockwise rotation here

	case DIRECTION_RIGHT:
		if !b.grid.tetriminoCausesCollision(b.current.row, b.current.col+1, b.current.Tetrimino) {
			b.current.col += 1
		}

	case DIRECTION_DOWN:
		if !b.grid.tetriminoCausesCollision(b.current.row+1, b.current.col, b.current.Tetrimino) {
			b.current.row += 1
		} else {
			b.AddTetrimino()
		}
	}

	b.grid.print(b.current)
}

func (b *Board) generateRandomTetrimino() *BoardTetrimino {
	tetrimino := generateRandomTetrimino()

	return &BoardTetrimino{
		col:       b.Columns/2 - len(tetrimino.Shape)/2,
		row:       0,
		Tetrimino: tetrimino,
	}
}

func (b *Board) Run() {
	if b.running {
		return
	}

	fmt.Println("STARTED BOARD")

	b.running = true

	ticker := time.NewTicker(250 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			b.move(DIRECTION_DOWN)
		case move_direction := <-b.moves:
			b.move(move_direction)
		}
	}
}

func NewBoard(rows, cols int) *Board {
	board := &Board{
		Rows:       rows,
		Columns:    cols,
		grid:       newGrid(rows, cols),
		tetriminos: make([]*BoardTetrimino, 0),
		moves:      make(chan uint8, 10),
	}

	board.AddTetrimino()

	return board
}
