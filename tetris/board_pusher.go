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
