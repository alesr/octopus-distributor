package distributor

import (
	"fmt"
	"strconv"

	"github.com/alesr/octopus-distributor/assembler"
)

// Receiver triggers the system to start receive requests and steam the machine.
func Receiver() {
	requestCh := make(chan []string)
	go assembler.GetRequest(requestCh)

	for i := 0; i < 10; i++ {
		identifier(i, <-requestCh)
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
func identifier(id int, req []string) {

	request := make(map[string]string)

	request["id"] = strconv.Itoa(id)
	request["op"] = req[0]

	// no mater how many args, we handle it!
	for i, arg := range req[1:] {
		key := fmt.Sprintf("arg %s", strconv.Itoa(i+1))
		request[key] = arg
	}
	fmt.Println(request)
}
