package subscriber

import "fmt"

func fibonacci(fibCh chan Fibonacci) {
	for {
		select {
		case f := <-fibCh:
			go func(fibCh chan Fibonacci) {

				f.result = nthFibonacci(f.n)
				fmt.Println(f)
			}(fibCh)
		}
	}
}

// Return the nth Fibonacci value
func nthFibonacci(pos int) int {
	c := fibonacciGen()
	f := make([]int, pos)
	for i := 0; i <= pos; i++ {
		f = append(f, <-c)
	}
	return f[len(f)-1]
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
