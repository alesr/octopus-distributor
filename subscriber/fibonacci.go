package subscriber

import (
	"fmt"
	"strconv"
)

type fibonacci struct {
	request       []string
	id, n, result int
	err           error
}

func runFibonacci(fibCh chan fibonacci) {
	for {
		select {
		case f := <-fibCh:

			if err := f.parser(); err != nil {
				return
			}
			f.nthFibonacci()
			fmt.Println(f)
		}
	}
}

// Parse the request as Fibonacci struct
func (fib *fibonacci) parser() error {

	nValue, err := strconv.Atoi(fib.request[1])
	if err != nil {
		fib.err = err
		return err
	}

	fib.n = nValue
	return nil
}

// Return the nth Fibonacci value
func (fib *fibonacci) nthFibonacci() {
	c := fibonacciGen()
	fibList := make([]int, fib.n)
	for i := 0; i <= fib.n; i++ {
		fibList = append(fibList, <-c)
	}
	fib.result = fibList[len(fibList)-1]
}

// Emit a infinite stream of Fibonacci values.
func fibonacciGen() chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()
	return c
}
