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
	requestFlag := flag.Int("request", 1000, "Type a number of requests to be simulated. Default: 1000")
	debugFlag := flag.Bool("debug", false, "True or false to use the debug mode. debug.out containt the requests will be generated under application root.")
	flag.Parse()

	subscriber.Run(*requestFlag, *debugFlag)
}
