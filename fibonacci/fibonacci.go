package fibonacci

import "strconv"

func Exec(fibCh chan []string, resultCh chan map[string]string) {

	f := <-fibCh
	fib := parse(f)
	nthFibonacci(fib)
	resultCh <- fib
}

// Parse the request
func parse(f []string) map[string]string {

	fib := make(map[string]string)
	fib["id"] = f[2]
	fib["task"] = f[0]
	fib["n"] = f[1]

	return fib
}

// Return the nth Fibonacci value
func nthFibonacci(fib map[string]string) {

	c := fibonacciGen()

	n, err := strconv.Atoi(fib["n"])
	if err != nil {
		fib["result"] = err.Error()
		return
	}

	if n == 0 {
		fib["result"] = "0"
		return
	}

	fibList := make([]int, n)
	for i := 0; i <= n; i++ {
		fibList = append(fibList, <-c)
	}
	fib["result"] = strconv.Itoa(fibList[len(fibList)-1])
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
