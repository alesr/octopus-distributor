package arithmetic

import (
	"errors"
	"fmt"
	"strconv"
)

var errZeroDivError = errors.New("cannot divide by zero")

func Exec(arithCh chan []string) {

	a := <-arithCh

	arith := parse(a)

	switch arith["task"] {
	case "add":
		add(arith)
	case "sub":
		subtract(arith)
	case "mult":
		multiply(arith)
	case "div":
		divide(arith)
	}

	fmt.Println(arith)
}

// Parse the request
func parse(a []string) map[string]string {

	arith := make(map[string]string)
	arith["id"] = a[3]
	arith["task"] = a[0]
	arith["a"] = a[1]
	arith["b"] = a[2]

	return arith
}

func add(arith map[string]string) {

	a, b, err := toInt(arith["a"], arith["b"])
	if err != nil {
		arith["result"] = err.Error()
		return
	}

	arith["result"] = strconv.Itoa(a + b)
}

func subtract(arith map[string]string) {

	a, b, err := toInt(arith["a"], arith["b"])
	if err != nil {
		arith["result"] = err.Error()
		return
	}
	arith["result"] = strconv.Itoa(a - b)
}

func multiply(arith map[string]string) {

	a, b, err := toInt(arith["a"], arith["b"])
	if err != nil {
		arith["result"] = err.Error()
		return
	}
	arith["result"] = strconv.Itoa(a * b)
}

func divide(arith map[string]string) {

	a, b, err := toInt(arith["a"], arith["b"])
	if err != nil {
		arith["result"] = err.Error()
		return
	}

	if b == 0 {
		arith["result"] = errZeroDivError.Error()
		return
	}

	// only the int part
	arith["result"] = strconv.Itoa(a / b)
}

func toInt(a, b string) (int, int, error) {

	intA, err := strconv.Atoi(a)
	if err != nil {
		return 0, 0, err
	}

	intB, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, err
	}

	return intA, intB, nil
}
