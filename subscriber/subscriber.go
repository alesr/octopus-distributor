package subscriber

import (
	"strconv"

	"github.com/alesr/octopus-distributor/arithmetic"
	"github.com/alesr/octopus-distributor/encode"
	"github.com/alesr/octopus-distributor/fibonacci"
	"github.com/alesr/octopus-distributor/publisher"
	"github.com/alesr/octopus-distributor/reverse"
)

var resultCh = make(chan map[string]string)

// Run trigger the system to start receiving requests
func Run(n int) {

	// Since the program starts here, let's make a channel to receive requests.
	// The buffer size is completely arbitrary, just prevents the sender from blocking too soon.
	requestCh := make(chan []string, 100)
	idCh := make(chan string)

	// If you want to play with us you need to register your Sender here.
	go publisher.Sender(requestCh)
	go makeID(idCh)

	// Our request pool
	for i := 1; i <= n; i++ {

		// DEBUG
		//fmt.Println(runtime.NumGoroutine())

		// get request
		request := <-requestCh

		// add i as ID
		request = append(request, <-idCh)

		distributor(request)

		// Send the result back to the publisher
		publisher.Receiver(<-resultCh)
	}
}

func makeID(idCh chan string) {
	for i := 1; ; i++ {
		idCh <- strconv.Itoa(i)
	}
}

// Distribute requests to respective channels.
// No waiting in line. Everybody gets its own goroutine.
func distributor(request []string) {

	go func() {
		switch request[0] {
		case "sum":
			go arithmetic.Exec(request, resultCh)
		case "sub":
			go arithmetic.Exec(request, resultCh)
		case "mult":
			go arithmetic.Exec(request, resultCh)
		case "div":
			arithmetic.Exec(request, resultCh)
		case "fibonacci":
			go fibonacci.Exec(request, resultCh)
		case "reverse":
			go reverse.Exec(request, resultCh)
		case "encode":
			go encode.Exec(request, resultCh)
		}
	}()
}
