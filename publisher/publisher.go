package publisher

/*
To make things interesting we want a high volume of requests being sent
to your queue manager a.k.a Octopus Distributor.

For that, we need to simulate a client sending as many requests as possible.
Meet, the Publisher package.

What the publisher package does is to group functions that help us to build
four types of fake messages that going to work as requests to our queue.

The messages types are arithmetic, fibonacci, reverse and encode.

After build these messages we send them randomly over the request channel.

This way, we should expect the channel receive messages like:

[fibo 5]
[reverse hello, world]
[add 2 2]
[div 4 2]
[encode foo bar]
[fib 10]
*/

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alesr/octopus-distributor/utilities"
)

// On the text.in file inside our data folder, we have some sample strings that
// we're going to use to form the text to be used by the reverse and encode
// functions to build fake messages.
// Since we don't want to load that file everytime we assemble a message, let's
// do that on a init function and keep on a variable called content.
var content []string

func init() {

	// Load file with sample text.
	txt, err := utilities.LoadFile("publisher/data/text.in")
	if err != nil {
		log.Fatal(err)
	}
	content = txt
}

// Sender call messenger() to build a message and delivery it as a request
// by sending it over the request channel.
func Sender(requestCh chan []string) {

	msg, err := messenger()
	if err != nil {
		log.Fatal(err)
	}
	requestCh <- msg
}

// Receiver output responses from response channel
func Receiver(responseCh chan map[string]string) {

	response := <-responseCh

	id := fmt.Sprintf("id: %s ", response["id"])

	switch response["task"] {
	case "add":
		fmt.Printf(id+"%s + %s = %s\n",
			response["a"], response["b"], response["result"])
	case "sub":
		fmt.Printf(id+"%s %s - %s = %s\n",
			response["a"], response["b"], response["result"])
	case "mult":
		fmt.Printf(id+"%s %s * %s = %s\n",
			response["a"], response["b"], response["result"])
	case "div":
		fmt.Printf(id+"%s %s / %s = %s\n",
			response["a"], response["b"], response["result"])
	case "fibonacci":
		fmt.Printf(id+"%s Fibonacci(%s) = %s\n",
			response["n"], response["result"])
	case "reverse":
		fmt.Printf(id+"%s Reverse: %s = %s\n",
			response["text"], response["result"])
	case "encode":
		fmt.Printf(id+"%s Encode: %s = %s\n",
			response["text"], response["result"])
	}
}

// Randomly choose a message type and call the proper function to build
// that message. Then return the message back to Sender function.
func messenger() ([]string, error) {

	// The basic tasks that the octopus are prepared to handle.
	taskList := []string{"arithmetic", "fibonacci", "reverse", "encode"}

	// Choose a random index from taskList
	index, err := utilities.Random(len(taskList))
	if err != nil {
		// If you are here and you don't know why, check if taskList is empty.
		return nil, err
	}

	// Hold the message to be sent.
	var msg []string

	// Depending on which task has been chosen,
	// we call the corresponding constructor function.
	switch taskList[index] {

	case "arithmetic":
		msg, err = arithmetic()
		if err != nil {
			return nil, err
		}

	case "fibonacci":
		msg = fibonacci()

	case "reverse":
		msg, err = reverse()
		if err != nil {
			return nil, err
		}

	case "encode":
		msg, err = encode()
		if err != nil {
			return nil, err
		}
	}
	return msg, nil
}

/// THE ASSEMBLERS

// Assemble basic arithmetic operations requests.
func arithmetic() ([]string, error) {

	// Take a random value for A and B up to 1000.
	maxValue := 1000
	a, _ := utilities.Random(maxValue)
	b, _ := utilities.Random(maxValue)

	// Now we select a random operation.
	operationsList := []string{"add", "sub", "mult", "div"}
	index, err := utilities.Random(len(operationsList))
	if err != nil {
		return nil, err
	}
	operation := operationsList[index]

	// Request ready, expect something like [mult 3 2]
	return []string{operation, strconv.Itoa(a), strconv.Itoa(b)}, nil
}

// A Fibonacci request.
func fibonacci() []string {

	// Fib 30 should be big enough for our case.
	n, _ := utilities.Random(30)
	return []string{"fibonacci", strconv.Itoa(n)}
}

// A reverse text request.
func reverse() ([]string, error) {

	// Get a random number between zero and the content length,
	// which is equivalent to a line at text.in.
	index, err := utilities.Random(len(content))
	if err != nil {
		return nil, err
	}

	// eg.: [reverse Papa Americano]
	return []string{"reverse", content[index]}, nil
}

// A encode text request.
func encode() ([]string, error) {

	// Same story once again.
	index, err := utilities.Random(len(content))
	if err != nil {
		return nil, err
	}

	// eg.: [encode hello, world]
	return []string{"encode", content[index]}, nil
}
