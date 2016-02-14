package subscriber

import (
	"log"
	"reflect"
	"strconv"

	"github.com/alesr/octopus-distributor/publisher"
)

var (
	arithCh = make(chan []int)
	fibCh   = make(chan Fibonacci)
	revCh   = make(chan Reverse)
	encCh   = make(chan string)
)

// Arithmetic request data
type Arithmetic struct {
	request          []string
	operation        string
	id, a, b, result int
}

// Fibonacci request data
type Fibonacci struct {
	request       []string
	id, n, result int
}

// Reverse request data
type Reverse struct {
	request      []string
	id           int
	text, result string
}

// Encode request data
type Encode struct {
	request      []string
	id           int
	text, result string
}

// Parse the request as Arithmetic struct
func (arith *Arithmetic) parse() error {

	arith.operation = arith.request[0]

	aValue, err := strconv.Atoi(arith.request[1])
	if err != nil {
		return err
	}

	arith.a = aValue

	bValue, err := strconv.Atoi(arith.request[1])
	if err != nil {
		return err
	}
	arith.b = bValue
	return nil
}

// Parse the request as Fibonacci struct
func (fib *Fibonacci) parse() error {

	nValue, err := strconv.Atoi(fib.request[1])
	if err != nil {
		return err
	}

	fib.n = nValue
	return nil
}

// Parse the request as Reverse struct
func (rev *Reverse) parse() error {
	rev.text = rev.request[1]
	return nil
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

	for i := 0; i < 500; i++ {
		classifier(i, <-requestCh)
	}
}

// Add an ID to each request and initialize structs
func classifier(i int, request []string) {

	var task interface{}

	switch request[0] {
	case "add":
		task = &Arithmetic{request: request, id: i}
	case "sub":
		task = &Arithmetic{request: request, id: i}
	case "mult":
		task = &Arithmetic{request: request, id: i}
	case "div":
		task = &Arithmetic{request: request, id: i}
	case "fibonacci":
		task = &Fibonacci{request: request, id: i}
	case "reverse":
		task = &Reverse{request: request, id: i}
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
	case "*subscriber.Arithmetic":

		arith := task.(*Arithmetic)

		if err := arith.parse(); err != nil {
			log.Fatal(err)
		}

	case "*subscriber.Fibonacci":

		fib := task.(*Fibonacci)

		if err := fib.parse(); err != nil {
			log.Fatal(err)
		}

		go fibonacci(fibCh)
		fibCh <- *fib

	case "*subscriber.Reverse":

		rev := task.(*Reverse)

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
