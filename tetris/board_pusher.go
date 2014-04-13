package tetris

import (
	"fmt"
	"github.com/timonv/pusher"
	"time"
)

func sendBoard(board *Board) error {

	data, err := BoardToJson(board)

	if err != nil {
		return err
	}

	client := pusher.NewClient("67131", "c2388f10a4afc865f3a5", "f4ab253ab07e424d993d")

	done := make(chan bool)

	go func() {

		err := client.Publish(string(data), "my_event", "my_channel")

		if err != nil {
			fmt.Printf("Error %s\n", err)
		} else {
			// fmt.Println("Pushed to pusher")
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
