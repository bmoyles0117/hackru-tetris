package tetris

import (
	// "bytes"
	// "encoding/json"
	// "errors"
	// "net/http"
	"fmt"
	"github.com/timonv/pusher"
	"time"
)

func sendBoard(board *Board) error {

	// data, err := json.Marshal(map[string]interface{}{
	// 	"data": map[string]interface{}{
	// 		"message": "hello",
	// 	},
	// })

	data := `{"board":[
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,1,0,1,0,0,1,0,1,0,0,0],
	[0,1,0,0,0,0,0,0,0,1,1,0],
	[0,1,0,0,1,0,0,0,0,1,0,0],
	[0,0,0,0,0,0,1,0,0,0,0,0],
	[0,0,1,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,1,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,1,0,0,0],
	[0,0,0,0,1,0,1,0,0,0,0,0],
	[0,0,0,0,1,0,0,0,1,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,1,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,1,0,0,0,0],
	[0,0,0,0,0,0,1,0,0,0,0,0],
	[0,0,0,0,1,0,0,0,0,0,0,0],
	[0,0,0,0,1,0,1,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0],
	[0,0,0,0,0,0,0,0,0,0,0,0]

	]}`

	client := pusher.NewClient("71664", "80f71c71ecfd0ce866eb", "94130ff49932d65f5515")

	done := make(chan bool)

	go func() {

		err := client.Publish(string(data), "my_event", "my_channel")

		if err != nil {
			fmt.Printf("Error %s\n", err)
		} else {
			fmt.Println("Pushed to pusher")
		}
		done <- true
	}()

	select {

	case <-done:
		fmt.Println("Done :-)")

	case <-time.After(1 * time.Minute):
		fmt.Println("Timeout :-(")

	}
	return nil

}
