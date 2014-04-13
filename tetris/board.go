package tetris

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	DIRECTION_LEFT   = 'L'
	DIRECTION_ROTATE = 'Z'
	DIRECTION_RIGHT  = 'R'
	DIRECTION_DOWN   = 'D'
)

type BoardTetrimino struct {
	col       int
	row       int
	Tetrimino *Tetrimino
}

type Board struct {
	Key     string
	Rows    int
	Columns int

	grid       *Grid
	current    *BoardTetrimino
	next       *BoardTetrimino
	tetriminos []*BoardTetrimino
	moves      chan uint8
	running    bool
	game_over  bool
	moves_list []string
	callbacks  []func()
}

func (b *Board) AddTetrimino() bool {
	if b.next != nil {
		// Stash the current piece in the list of pieces
		b.tetriminos = append(b.tetriminos, b.current)

		b.grid.consumeTetrimino(b.current.row, b.current.col, b.current.Tetrimino)

		if b.grid.tetriminoCausesCollision(b.next.row, b.next.col, b.next.Tetrimino) {
			b.grid.consumeTetrimino(b.next.row, b.next.col, b.next.Tetrimino)

			return false
		}

		b.current = b.next

	} else {
		b.current = b.generateRandomTetrimino()

	}

	b.next = b.generateRandomTetrimino()

	return true
}

func (b *Board) OnGameover(callback func()) {
	b.callbacks = append(b.callbacks, callback)
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

	case DIRECTION_ROTATE:
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
			// Collision check is in AddTetrimino
			if !b.AddTetrimino() {
				for i := range b.callbacks {
					b.callbacks[i]()
				}

				b.game_over = true
			}

		}
	}
}

func (b *Board) generateRandomTetrimino() *BoardTetrimino {
	tetrimino := generateRandomTetrimino()

	return &BoardTetrimino{
		col:       b.Columns/2 - len(tetrimino.Shape)/2,
		row:       0,
		Tetrimino: tetrimino,
	}
}

func (b *Board) reset() {
	b.game_over = false
	b.running = false
	b.current = nil
	b.next = nil
	b.grid = newGrid(b.Rows, b.Columns)
	b.tetriminos = make([]*BoardTetrimino, 0)
	b.moves = make(chan uint8, 10)

	b.AddTetrimino()
}

func (b *Board) Run() {
	if b.running {
		return
	}

	b.reset()

	fmt.Println("STARTED BOARD")

	b.running = true

	// ticker := time.NewTicker(1 * time.Second)

	ticker := time.NewTicker(1000 * time.Millisecond)

	for {
		go sendBoard(b)

		if !b.running {
			return
		}

		if b.game_over {
			b.running = false
			return
		}

		select {
		case <-ticker.C:
			b.move(DIRECTION_DOWN)

		case move_direction := <-b.moves:
			b.move(move_direction)

		}

	}
}

func (b *Board) Stop() {
	b.running = false
}

func NewBoard(board_key string, rows, cols int) *Board {
	board := &Board{
		Key:        board_key,
		Rows:       rows,
		Columns:    cols,
		grid:       newGrid(rows, cols),
		tetriminos: make([]*BoardTetrimino, 0),
		moves:      make(chan uint8, 10),
		callbacks:  make([]func(), 0),
	}

	return board
}

func convertArrays(input [][]uint8) [][]int {
	converted := make([][]int, len(input))

	for row := range input {
		converted[row] = make([]int, len(input[row]))

		for col := range input[row] {
			converted[row][col] = int(input[row][col])
		}
	}

	return converted
}

func BoardToJson(b *Board) ([]byte, error) {
	tcells := convertArrays(b.grid.cells)

	// Track the falling shape
	for y := range b.current.Tetrimino.Shape {
		for x := range b.current.Tetrimino.Shape[y] {
			if y+b.current.row < len(b.grid.cells) && b.current.Tetrimino.Shape[y][x] == 1 {
				tcells[y+b.current.row][x+b.current.col] = int(b.current.Tetrimino.Type)
			}
		}
	}

	return json.Marshal(map[string]interface{}{
		"board_key": b.Key,
		"running":   b.running,
		"game_over": b.game_over,
		"cells":     tcells,
		"next_piece": map[string]interface{}{
			"color": b.next.Tetrimino.Type,
			"shape": convertArrays(b.next.Tetrimino.Shape),
		},
	})
}
