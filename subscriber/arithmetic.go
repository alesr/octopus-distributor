package subscriber

import (
	"errors"
	"fmt"
	"strconv"
)

var errZeroDivError = errors.New("cannot divide by zero")

// arithmetic request data
type arithmetic struct {
	request          []string
	operation        string
	id, a, b, result int
	err              error
}

func runArithmetic(arithCh chan arithmetic) {
	for {
		select {
		case a := <-arithCh:

			if err := a.parser(); err != nil {
				return
			}

			switch a.operation {
			case "add":
				a.addition()
			case "sub":
				a.subtraction()
			case "mult":
				a.multiplication()
			case "div":
				a.division()
			}
			fmt.Println(a)
		}
	}
}

// Parse the request as Arithmetic struct
func (arith *arithmetic) parser() error {

	arith.operation = arith.request[0]

	aValue, err := strconv.Atoi(arith.request[1])
	if err != nil {
		arith.err = err
		return err
	}

	arith.a = aValue

	bValue, err := strconv.Atoi(arith.request[2])
	if err != nil {
		arith.err = err
		return err
	}
	arith.b = bValue
	arith.err = err
	return err
}

func (arith *arithmetic) addition() {
	arith.result = arith.a + arith.b
}

func (arith *arithmetic) subtraction() {
	arith.result = arith.a - arith.b
}

func (arith *arithmetic) multiplication() {
	arith.result = arith.a * arith.b
}

func (arith *arithmetic) division() {

	if arith.b == 0 {
		arith.err = errZeroDivError
		return
	}

	// only the int part
	arith.result = arith.a / arith.b
}

func (arith arithmetic) String() string {

	var symbol string

	switch arith.operation {
	case "add":
		symbol = "+"
	case "sub":
		symbol = "-"
	case "mult":
		symbol = "*"
	case "div":
		symbol = "/"
	}
	return fmt.Sprintf("ID: %d   %d %s %d = %d", arith.id, arith.a, symbol, arith.b, arith.result)
}
