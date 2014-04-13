package tetris

import (
	"encoding/json"
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
	cells [][]uint8
}

func (g *Grid) clearFilledLines() {
	fmt.Println("CHECKING LINES")
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
	fmt.Println("INSERT INTO ROW", row, "COL", col)
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

func (g *Grid) print(current *BoardTetrimino) {
	tcells := make([][]uint8, len(g.cells))

	for y := range g.cells {
		tcells[y] = make([]uint8, len(g.cells[y]))
		copy(tcells[y], g.cells[y])
	}

	for y := range current.Tetrimino.Shape {
		for x := range current.Tetrimino.Shape[y] {
			if y+current.row < len(g.cells) && current.Tetrimino.Shape[y][x] == 1 {
				tcells[y+current.row][x+current.col] = current.Tetrimino.Type
			}
		}
	}

	for row := range tcells {
		for col := range tcells[row] {
			if tcells[row][col] != 0 {
				fmt.Printf(" %c ", tcells[row][col])
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println("------------------------------------")
}

func newGrid(rows, cols int) *Grid {
	cells := make([][]uint8, rows)

	for i := range cells {
		cells[i] = make([]uint8, cols)
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
	game_over  bool
}

func (b *Board) AddTetrimino() bool {
	if b.next != nil {
		// Stash the current piece in the list of pieces
		fmt.Println("mewoz", b.next.row)

		b.tetriminos = append(b.tetriminos, b.current)

		b.grid.consumeTetrimino(b.current.row, b.current.col, b.current.Tetrimino)

		if !b.grid.tetriminoCausesCollision(b.next.row, b.next.col, b.next.Tetrimino) {

			fmt.Println("ENDING HERE ", b.next.row)
			return false
		}
		b.current = b.next

	} else {
		b.current = b.generateRandomTetrimino()

	}

	b.next = b.generateRandomTetrimino()

	return true

}

func (b *Board) Move(move_direction uint8) {
	b.moves <- move_direction
}

func (b *Board) move(move_direction uint8) {
	b.grid.clearFilledLines()

	switch move_direction {
	case DIRECTION_LEFT:
		if !b.grid.tetriminoCausesCollision(b.current.row, b.current.col-1, b.current.Tetrimino) {
			b.current.col -= 1
		}

	case DIRECTION_UP:
		transposed_tetrimino := b.current.Tetrimino.Rotate()
		if !b.grid.tetriminoCausesCollision(b.current.row, b.current.col+1, transposed_tetrimino) {
			b.current.Tetrimino = transposed_tetrimino
		}

		// Add counter clockwise rotation here

	case DIRECTION_RIGHT:
		if !b.grid.tetriminoCausesCollision(b.current.row, b.current.col+1, b.current.Tetrimino) {
			b.current.col += 1
		}

	case DIRECTION_DOWN:
		if !b.grid.tetriminoCausesCollision(b.current.row+1, b.current.col, b.current.Tetrimino) {
			b.current.row += 1
		} else {
			if !b.AddTetrimino() { //check is in AddTetrimino
				b.game_over = true
			}

		}
	}

	// b.grid.print(b.current)
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

	// ticker := time.NewTicker(1 * time.Second)

	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		if b.game_over {
			b.running = false
			break
		}

		go sendBoard(b)

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

func BoardToJson(b *Board) ([]byte, error) {
	// Converting to ints so the values don't condense to alphanumeric
	tcells := make([][]int, len(b.grid.cells))

	for y := range b.grid.cells {
		tcells[y] = make([]int, len(b.grid.cells[y]))
		for x := range b.grid.cells[y] {
			tcells[y][x] = int(b.grid.cells[y][x])
		}
	}

	// Track the falling shape
	for y := range b.current.Tetrimino.Shape {
		for x := range b.current.Tetrimino.Shape[y] {
			if y+b.current.row < len(b.grid.cells) && b.current.Tetrimino.Shape[y][x] == 1 {
				tcells[y+b.current.row][x+b.current.col] = int(b.current.Tetrimino.Type)
			}
		}
	}

	return json.Marshal(map[string]interface{}{
		"board": tcells,
	})
}
