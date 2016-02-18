package main

import (
	"flag"
	"runtime"

	"github.com/alesr/octopus-distributor/subscriber"
)

func main() {

	//
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Set flag
	flagPtr := flag.Int("request", 1000, "Type a number of requests to be simulated. Default: 1000")
	flag.Parse()

	subscriber.Run(*flagPtr)
}
