package main

import (
	"fmt"

	"github.com/alesr/octopus-distributor/assembler"
)

func main() {

	requestCh := make(chan []string)
	doneCh := make(chan bool)
	go assembler.GetRequest(requestCh, doneCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-requestCh)
	}
	doneCh <- true
	close(doneCh)
}
