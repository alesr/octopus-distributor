package main

import (
	"runtime"

	"github.com/alesr/octopus-distributor/subscriber"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	subscriber.Run()
}
