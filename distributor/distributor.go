package distributor

import (
	"log"

	"github.com/alesr/octopus-distributor/assembler"
)

// Arithmetic ...
type Arithmetic struct {
	request          []string
	operation        string
	id, a, b, result int
}

// Fibonacci ...
type Fibonacci struct {
	request       []string
	id, n, result int
}

// Reverse ...
type Reverse struct {
	request      []string
	id           int
	text, result string
}

// Encode ..
type Encode struct {
	request      []string
	id           int
	text, result string
}

// Parser ..
type Parser interface {
	parse() error
}

func (a *Arithmetic) parse(request map[string]string) error {
	return nil
}

func (f *Fibonacci) parse(request map[string]string) error {
	return nil
}

func (r *Reverse) parse(request map[string]string) error {
	return nil
}

func (e *Encode) parse(request map[string]string) error {
	return nil
}

// Receive triggers the system to start receive requests and steam the machine.
func Receive() {
	requestCh := make(chan []string)
	go assembler.GetRequest(requestCh)

	for i := 0; i < 10; i++ {
		register(i, <-requestCh)
	}
}

// to make things more interesting let's add an ID to each request.
// doesn't make sense add the ID in the assembler side since in a real situation
// we cannot expect that every "client" can control the ID attribution.
// even worse if we have many clients sending requests to the request channel.
// also, so far a request is formed by an operator which can be "arithmetic",
// "fibonacci", "reverse" or "encode" and arguments which are the N following
// values after the operator, like [add 2 3] where 2 and 3 are the arguments
// for the addition in this case. said that, would be good to organize the request
// in a map to make it easy to handle later.
// we'll organize the received request under the form:
// [id: XX, op: encode, argX: some string]
func register(i int, request []string) {

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
	case "enconde":
		task = &Encode{request: request, id: i}
	default:

		// REVIEW: to work on data later
		log.Fatal("invalid request")
	}

}
