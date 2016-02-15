package subscriber

import (
	"strconv"
	"time"

	"github.com/alesr/octopus-distributor/arithmetic"
	"github.com/alesr/octopus-distributor/encode"
	"github.com/alesr/octopus-distributor/fibonacci"
	"github.com/alesr/octopus-distributor/publisher"
	"github.com/alesr/octopus-distributor/reverse"
)

var (
	arithCh = make(chan []string)
	fibCh   = make(chan []string)
	revCh   = make(chan []string)
	encCh   = make(chan []string)
)

// Receiver triggers the system to start receiving requests
func Receiver() {
	requestCh := make(chan []string)

	go func() {
		for {
			publisher.GetRequest(requestCh)
		}
	}()

	// i will be our ID
	for i := 1; i <= 1000; i++ {
		request := <-requestCh
		request = append(request, strconv.Itoa(i))
		distributor(request)
	}
	time.Sleep(time.Second * 1)
}

// Distribute requests to respective channels.
func distributor(request []string) {

	switch request[0] {

	case "add":
		go arithmetic.Exec(arithCh)
		arithCh <- request
	case "sub":
		go arithmetic.Exec(arithCh)
		arithCh <- request
	case "mult":
		go arithmetic.Exec(arithCh)
		arithCh <- request
	case "div":
		go arithmetic.Exec(arithCh)
		arithCh <- request
	case "fibonacci":
		go fibonacci.Exec(fibCh)
		fibCh <- request
	case "reverse":
		go reverse.Exec(revCh)
		revCh <- request
	case "encode":
		go encode.Exec(encCh)
		encCh <- request

	}
}
