package fibonacci

import "strconv"

func Exec(request []string, resultCh chan map[string]string) {

	fib := parse(request)

	n, err := strconv.Atoi(fib["n"])
	if err != nil {
		fib["result"] = err.Error()
		return
	}

	if n < 0 {
		fib["result"] = "n must be a positive number"
		return
	}

	fib["result"] = strconv.Itoa(nthFibonacci(n))
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
func nthFibonacci(n int) int {

	if n == 0 || n == 1 {
		return n
	}
	return nthFibonacci(n-1) + nthFibonacci(n-2)
}
