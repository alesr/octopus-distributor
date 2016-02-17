package main

import (
	"runtime"

	"github.com/alesr/octopus-distributor/subscriber"
)

func main() {

	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)

	subscriber.Run()
}
