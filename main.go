package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	failures := make(chan error)

	go func() {
		time.Sleep(12 * time.Second)
		failures <- errors.New("error added to channel after 12 seconds")
	}()

	go func() {
		time.Sleep(8 * time.Second)
		failures <- errors.New("error added to channel after 8 seconds")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		failures <- nil
	}()

	for {
		select {
		case err := <-failures:
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("No error happened")
		}
	}

}
