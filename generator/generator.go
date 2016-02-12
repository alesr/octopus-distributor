package generator

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alesr/octopus-distributor/utilities"
)

// Messages randomly creates elements to be sent to our dear octopus,
// so he can distribute them to the right hands.
func Messages() {

	// the basic operations that the octopus are prepared to handle.
	taskList := []string{"arithmetic", "fibonacci", "reverse", "encode"}

	// index, err := utilities.Random(len(taskList))
	_, err := utilities.Random(len(taskList))

	if err != nil {
		// if you are here and you don't know why, checks if taskList is empty
		log.Fatal(err)
	}

	var msg []string
	// switch taskList[index] {
	switch "reverse" {
	case "arithmetic":
		msg, err = arithmetic()
		if err != nil {
			log.Fatal(err)
		}
	case "fibonacci":
		msg = fibonacci()
	case "reverse":
		msg, err = reverse()
		if err != nil {
			log.Fatal(err)
		}
	case "encode":
	}
	fmt.Println(msg)
}

// builds a fake arithmetic operation
func arithmetic() ([]string, error) {

	// let's take a random value for a and b
	maxValue := 100
	a, _ := utilities.Random(maxValue)
	b, _ := utilities.Random(maxValue)

	// now a random operation
	operationsList := []string{"add", "sub", "mult", "div"}
	index, err := utilities.Random(len(operationsList))
	if err != nil {
		return nil, err
	}
	operation := operationsList[index]
	return []string{operation, strconv.Itoa(a), strconv.Itoa(b)}, nil
}

// a fibonacci query
func fibonacci() []string {
	n, _ := utilities.Random(50)
	return []string{"fibonacci", strconv.Itoa(n)}
}

func reverse() ([]string, error) {

	// load file with sample text
	content, err := utilities.LoadFile("generator/data/text.in")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	index, err := utilities.Random(len(content))
	if err != nil {
		return nil, err
	}
	return []string{"reverse", content[index]}, nil
}
