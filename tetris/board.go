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
	X         int
	Y         int
	Tetrimino *Tetrimino
}

type Board struct {
	Rows    int
	Columns int

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
	fmt.Printf("MOVE DIRECTION: %c\n", move_direction)
	fmt.Printf("CURRENT: %s\n", b.current)
	fmt.Printf("TETRIMINOS: %s\n", b.tetriminos)
	switch move_direction {
	case DIRECTION_LEFT:
		if b.current.X > 0 {
			b.current.X -= 1
		}

	case DIRECTION_UP:
		// Add counter clockwise rotation here

	case DIRECTION_RIGHT:
		if b.current.X < b.Columns-1 {
			b.current.X += 1
		}

	case DIRECTION_DOWN:
		// Add better collision detection here
		if b.current.Y < b.Rows-1 {
			b.current.Y += 1
		} else {
			// We have met the bottom, make a new piece!
			b.AddTetrimino()
		}
	}
}

func (b *Board) generateRandomTetrimino() *BoardTetrimino {
	tetrimino := generateRandomTetrimino()

	return &BoardTetrimino{
		X:         b.Columns/2 - len(tetrimino.Shape)/2,
		Y:         0,
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

	ticker2 := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			b.move(DIRECTION_DOWN)

		case move_direction := <-b.moves:
			b.move(move_direction)

		case <-ticker2.C:
			err := sendBoard(b)
			if err != nil {
				fmt.Printf("Error : %s", err)
			}
		}

	}
}

func NewBoard(rows, columns int) *Board {
	board := &Board{
		Rows:       rows,
		Columns:    columns,
		tetriminos: make([]*BoardTetrimino, 0),
		moves:      make(chan uint8, 10),
	}

	board.AddTetrimino()

	return board
}
