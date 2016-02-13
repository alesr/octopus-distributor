package main

import (
	"fmt"

	"github.com/alesr/octopus-distributor/generator"
)

func main() {

	requestCh := make(chan []string)
	doneCh := make(chan bool)
	go generator.GetRequest(requestCh, doneCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-requestCh)
	}
	doneCh <- true
	close(doneCh)
}
