package controllers

import (
	"fmt"
	"github.com/bmoyles0117/hackru-tetris/app"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Move() revel.Result {
	direction_string := c.Params.Get("Body")

	received_number := c.Params.Get("From")
	recieved_data := received_number[2:5] + "*******" + received_number[9:12] + " sent : " + direction_string
	fmt.Println(recieved_data)

	app.Boards[c.Params.Get("To")].Move(direction_string[0])

	return c.RenderText("Moved!")
}

func (c App) Start() revel.Result {
	for i := range app.Boards {
		go app.Boards[i].Run()
	}

	return c.RenderText("OK")
}
