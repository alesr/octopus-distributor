package subscriber

import (
	"log"
	"reflect"

	"github.com/alesr/octopus-distributor/publisher"
)

var (
	arithCh = make(chan arithmetic)
	fibCh   = make(chan fibonacci)
	revCh   = make(chan reverse)
	encCh   = make(chan Encode)
)

// Encode request data
type Encode struct {
	request      []string
	id           int
	text, result string
	err          error
}

// Parse the request as Encode struct
func (enc *Encode) parse() error {
	enc.text = enc.request[1]
	return nil
}

// Receiver triggers the system to start receiving requests
func Receiver() {
	requestCh := make(chan []string)
	go publisher.GetRequest(requestCh)

	for i := 0; i < 10; i++ {
		classifier(i, <-requestCh)
	}
}

// Add an ID to each request and initialize structs
func classifier(i int, request []string) {

	var task interface{}

	switch request[0] {
	case "add":
		task = &arithmetic{request: request, id: i}
	case "sub":
		task = &arithmetic{request: request, id: i}
	case "mult":
		task = &arithmetic{request: request, id: i}
	case "div":
		task = &arithmetic{request: request, id: i}
	case "fibonacci":
		task = &fibonacci{request: request, id: i}
	case "reverse":
		task = &reverse{request: request, id: i}
	case "encode":
		task = &Encode{request: request, id: i}
	default:
		log.Fatal("invalid request")
	}
	distributor(task)
}

// Organize requests into types and distribute them to respective channels.
func distributor(task interface{}) {

	// Checks the underlying type held by the empty interface
	// and assert the corresponding type to the task.
	// After that, call parse method to fill the struct fields
	// and send the problem to the right hands.
	switch reflect.TypeOf(task).String() {
	case "*subscriber.arithmetic":

		arith := task.(*arithmetic)

		if err := arith.parse(); err != nil {
			return
		}

		go calcArithmetic(arithCh)

		arithCh <- *arith

	case "*subscriber.fibonacci":

		fib := task.(*fibonacci)

		if err := fib.parse(); err != nil {
			return
		}

		go calcFibonacci(fibCh)
		fibCh <- *fib

	case "*subscriber.reverse":

		rev := task.(*reverse)

		if err := rev.parse(); err != nil {
			log.Fatal(err)
		}

		go reverser(revCh)
		revCh <- *rev

	case "*subscriber.Encode":
		enc := task.(*Encode)

		if err := enc.parse(); err != nil {
			log.Fatal(err)
		}
	}
}
