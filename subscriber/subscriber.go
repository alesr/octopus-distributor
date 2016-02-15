package subscriber

import (
	"log"
	"reflect"
	"time"

	"github.com/alesr/octopus-distributor/publisher"
)

var (
	arithCh = make(chan arithmetic)
	fibCh   = make(chan fibonacci)
	revCh   = make(chan reverse)
	encCh   = make(chan encode)
)

// Receiver triggers the system to start receiving requests
func Receiver() {
	requestCh := make(chan []string)
	go publisher.GetRequest(requestCh)

	// i will be our ID
	for i := 1; i <= 100000; i++ {
		classifier(i, <-requestCh)
	}
	time.Sleep(time.Second * 3)
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
		task = &encode{request: request, id: i}
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

		go runArithmetic(arithCh)
		arithCh <- *arith

	case "*subscriber.fibonacci":
		fib := task.(*fibonacci)

		go runFibonacci(fibCh)
		fibCh <- *fib

	case "*subscriber.reverse":
		rev := task.(*reverse)

		go runReverse(revCh)
		revCh <- *rev

	case "*subscriber.encode":
		enc := task.(*encode)

		go runEncode(encCh)
		encCh <- *enc
	}
}
