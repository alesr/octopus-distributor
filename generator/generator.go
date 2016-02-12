package generator

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	errRandomFuncError = errors.New("random: function expect an int arg greater than zero")
)

// Messages randomly creates messages to be sent to our dear octopus so he can
// distribute them to the right hands.
func Messages() {

	// the basic operations that the octopus are prepared to handle.
	taskList := []string{"arithmetic", "fibonacci", "reverse", "encode"}

	index, err := random(len(taskList))

	if err != nil {
		// if you are here and you don't know why, checks if taskList is empty
		log.Fatal(err)
	}

	switch taskList[index] {
	case "arithmetic":

	case "fibonacci":
	case "reverse":
	case "encode":
	}

	fmt.Println()

}

// returns a pseudo random number between zero and max
func random(max int) (int, error) {

	// it means that the source slice is empty.
	if max < 0 {
		return 0, errRandomFuncError
	}

	rand.Seed(time.Now().Unix())
	return rand.Intn(max), nil
}
