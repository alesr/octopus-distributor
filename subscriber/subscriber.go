package subscriber

import (
	"strconv"

	"github.com/alesr/octopus-distributor/arithmetic"
	"github.com/alesr/octopus-distributor/encode"
	"github.com/alesr/octopus-distributor/fibonacci"
	"github.com/alesr/octopus-distributor/publisher"
	"github.com/alesr/octopus-distributor/reverse"
)

var (
	// Those channels are used to share data between the subscriber
	// and the process responsible to solve the tasks
	arithCh = make(chan []string)
	fibCh   = make(chan []string)
	revCh   = make(chan []string)
	encCh   = make(chan []string)

	// This channels communicate the results coming
	// from the task solvers back to the subscriber.
	resultCh   = make(chan map[string]string)
	responseCh = make(chan map[string]string)
)

// Run trigger the system to start receiving requests
func Run() {

	// Since the programs starts here, let's make a channel to receive requests
	requestCh := make(chan []string)

	// If you want to play with us you need to register your Sender here
	publisher.Sender(requestCh)

	for i := 1; i <= 100000; i++ {

		select {
		case request := <-requestCh:
			request = append(request, strconv.Itoa(i))
			distributor(request)
		case result := <-resultCh:
			publisher.Receiver(responseCh)
			responseCh <- result
		}
	}

	//time.Sleep(time.Second * 3)
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
