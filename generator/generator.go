package generator

/*
To make things interesting we want a high volume of requests being sent
to your queue manager a.k.a Octopus Distributor.

For that, we need to simulate a client sending as many requests as possible.
Meet, Generator.

What the generator package does is to group functions that help us to build
four types of fake messages that going to work as requests to our queue.

The messages types are arithmetic, fibonacci, reverse and encode.

After build these requests we send them randomly over a channel our octopus so he can do his  distribution.

This way we should expect the channel receive messages like:

[fibo 5]
[reverse hello, world]
[add 2 2]
[div 4 2]
[encode foo bar]
[fib 10]
*/

import (
	"log"
	"strconv"

	"github.com/alesr/octopus-distributor/utilities"
)

// On text.in inside our data folder, we have some sample strings that we going  to use
// to form the text to be used by the reverse and encode function so build fake messages.
// Since we don't want to load that file everytime we build a message,
// let's do that just once on a init function and keep on a variable called content.
var content []string

func init() {

	// load file with sample text
	txt, err := utilities.LoadFile("generator/data/text.in")
	if err != nil {
		log.Fatal(err)
	}
	content = txt
}

// GetRequest call messenger() to delivery some message
// and send it over the request channel
func GetRequest(requestCh chan []string, doneCh chan bool) {

	// the default behaviour is to send messages to request	channel.
	// as soon we receive from done channel, we return the function
	// and the channel can be safely closed by the caller
	go func() {
		for {
			select {
			case <-doneCh:
				return
			default:
				msg, err := messenger()
				if err != nil {
					log.Fatal(err)
				}
				requestCh <- msg
			}
		}
	}()
}

// messenger randomly chooses a message type and call the proper function to build this message.
// After that, sends the message to our dear octopus.
func messenger() ([]string, error) {

	// the basic operations that the octopus are prepared to handle.
	taskList := []string{"arithmetic", "fibonacci", "reverse", "encode"}

	// choose a random index inside taskList
	index, err := utilities.Random(len(taskList))
	if err != nil {
		// if you are here and you don't know why, checks if taskList is empty
		log.Fatal(err)
	}

	// holds the message to be sent
	var msg []string

	// depending on which task has been chosen,
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

// builds basic arithmetic operations
func arithmetic() ([]string, error) {

	// let's take a random value for a and b
	maxValue := 100
	a, _ := utilities.Random(maxValue)
	b, _ := utilities.Random(maxValue)

	// now we select random operation
	operationsList := []string{"add", "sub", "mult", "div"}
	index, err := utilities.Random(len(operationsList))
	if err != nil {
		return nil, err
	}
	operation := operationsList[index]

	// message ready, expect something like [mult 3 2]
	return []string{operation, strconv.Itoa(a), strconv.Itoa(b)}, nil
}

// a fibonacci message
func fibonacci() []string {
	// fib 30 should be big enough for our case
	n, _ := utilities.Random(30)
	return []string{"fibonacci", strconv.Itoa(n)}
}

// a reverse query
func reverse() ([]string, error) {

	// get a rand number between zero and content length
	// which means a line at text.in
	index, err := utilities.Random(len(content))
	if err != nil {
		return nil, err
	}

	// eg.: [reverse Papa Americano]
	return []string{"reverse", content[index]}, nil
}

func encode() ([]string, error) {

	// same story once again
	index, err := utilities.Random(len(content))
	if err != nil {
		return nil, err
	}

	// eg.: [encode hello, world]
	return []string{"encode", content[index]}, nil
}
