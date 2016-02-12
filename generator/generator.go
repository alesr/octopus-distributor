package generator

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var (
	errRandomFuncError = errors.New("random: function expect an int arg greater than zero")
)

// Messages randomly creates elements to be sent to our dear octopus,
// so he can distribute them to the right hands.
func Messages() {

	// the basic operations that the octopus are prepared to handle.
	taskList := []string{"arithmetic", "fibonacci", "reverse", "encode"}

	index, err := random(len(taskList))

	if err != nil {
		// if you are here and you don't know why, checks if taskList is empty
		log.Fatal(err)
	}

	var msg []string
	switch taskList[index] {
	case "arithmetic":
		msg, err = arithmetic()
		if err != nil {
			log.Fatal(err)
		}
	case "fibonacci":
		msg = fibonacci()
	case "reverse":
	case "encode":
	}

	fmt.Println(msg)
}

// returns a pseudo random number between zero and max
func random(max int) (int, error) {
	// it means that the source slice is empty.
	if max < 0 {
		return 0, errRandomFuncError
	}

	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max), nil
}

// builds a fake arithmetic operation
func arithmetic() ([]string, error) {

	// let's take a random value for a and b
	maxValue := 100
	a, _ := random(maxValue)
	b, _ := random(maxValue)

	// now a random operation
	operationsList := []string{"add", "sub", "mult", "div"}
	index, err := random(len(operationsList))
	if err != nil {
		return nil, errRandomFuncError
	}
	operation := operationsList[index]
	return []string{operation, strconv.Itoa(a), strconv.Itoa(b)}, nil
}

// a fibonacci operation
func fibonacci() []string {
	n, _ := random(50)
	return []string{"fibonacci", strconv.Itoa(n)}
}
