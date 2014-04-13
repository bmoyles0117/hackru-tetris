package controllers

import (
	"fmt"
	"github.com/bmoyles0117/hackru-tetris/tetris"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var board *tetris.Board

func getBoard() *tetris.Board {
	if board == nil {
		board = tetris.NewBoard(22, 12)
		go board.Run()
	}

	return board
}

func (c App) Index() revel.Result {
	fmt.Println("BOARD: ", getBoard())

	return c.Render()
}

func (c App) Move() revel.Result {
	direction_string := c.Params.Get("Body")

	getBoard().Move(direction_string[0])

	// fmt.Println(c.Params.Get("test"))

	return c.RenderText("Moved!")
}
