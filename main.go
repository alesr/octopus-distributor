package main

import (
	"fmt"

	"github.com/alesr/octopus-distributor/generator"
)

func main() {

	// creates a channel for requests
	// these requests going to be sent by a hyp
	requestCh := make(chan []string)
	go generator.Messenger(requestCh)

	for i := 0; i < 10; i++ {
		select {
		case msg := <-requestCh:
			fmt.Println(msg)
		}
	}

}
