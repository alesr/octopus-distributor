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
	arithCh    = make(chan []string)
	fibCh      = make(chan []string)
	revCh      = make(chan []string)
	encCh      = make(chan []string)
	resultCh   = make(chan map[string]string)
	responseCh = make(chan map[string]string)
)

// Trigger the system to start receiving requests
func Run() {
	requestCh := make(chan []string)

	go func() {
		for {
			publisher.Sender(requestCh)
		}
	}()

	// i will be our ID
	for i := 1; i <= 10000; i++ {

		select {
		case request := <-requestCh:
			request = append(request, strconv.Itoa(i))
			distributor(request)
		case result := <-resultCh:

			go publisher.Receiver(responseCh)
			responseCh <- result
		}
	}

	time.Sleep(time.Second * 1)
}

// Distribute requests to respective channels.
// No waiting in line. Everybody gets its own goroutine!
func distributor(request []string) {

	switch request[0] {

	case "add":
		go arithmetic.Exec(arithCh, resultCh)
		arithCh <- request
	case "sub":
		go arithmetic.Exec(arithCh, resultCh)
		arithCh <- request
	case "mult":
		go arithmetic.Exec(arithCh, resultCh)
		arithCh <- request
	case "div":
		go arithmetic.Exec(arithCh, resultCh)
		arithCh <- request
	case "fibonacci":
		go fibonacci.Exec(fibCh, resultCh)
		fibCh <- request
	case "reverse":
		go reverse.Exec(revCh, resultCh)
		revCh <- request
	case "encode":
		go encode.Exec(encCh, resultCh)
		encCh <- request
	}

}
