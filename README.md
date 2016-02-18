# octopus-distributor

### Project structure
```
.
├─Dockerfile
├─.travis.yml
├─README.md
├─main.go
├─publisher
| ├─publisher.go
| └─data
|   └─text.in
├─subscriber
| └─subscriber.go
├─arithmetic
| ├─arithmetic.go
| └─arithmetic_test.go
├─fibonacci
| ├─fibonacci.go
| └─fibonacci_test.go
├─reverse
| ├─reverse.go
| └─reverse_test.go
├─encode
| ├─encode.go
| └─encode_test.go
└─utilities
  ├─utilities.go
  └─utilities_test.go
```

### How to run

If you have Go installed:
```
git clone github.com/alesr/octopus-distributor
cd octopus-distributor/
go run main.go
```
if you have Docker:
```
docker run -it --rm --name octopus alesr/octopus-distributor
```

### How it works

In order to see our program working, we need some entity to simulate a high volume of tasks to be handled by the distributor.

The Publisher package assembles random tasks of four different types: arithmetic for basic mathematic operations, Fibonacci, reverse for mirrored text and encode for bcrypt hash values.

After building the tasks, it sends these messages through the request channel.

The Subscriber package at its side registers the publisher and starts receiving from the request channel. Basically, the Subscriber could register any Publisher interested in having tasks handled by the system by sending them trough the same request channel.

After receiving the request the Subscriber checks if the task can be handled. And if so, distribute it to be solved by one of the four external packages responsible for executing that particular task.

The four mentioned packages are the Arithmetic, Fibonacci, Reverse and Encode. Which one receiving tasks through their own channel and giving back the resulting calculation through the result's channel.

Once the Subscriber receives from the results channel. It sends the result back to the Publisher via the response channel so that the Receiver function can now get from the channel and output the response putting an end to the cycle.

#### Some important considerations in this process to justify my implementation for the test.

Channels are typed conduits that behave like FIFO queues, hence, they preserve the order of the items that are sent to them and block from receiving new items while the other side is not ready to handle them. Also, is important to not that each time the Subscriber receive new requests it launches a new goroutine to do the calculation. Meaning that the program solves the tasks concurrently and because of that, there is no waiting in line for requests, even if there is still a line.

Another thing is that the program is divided into three major components. The publisher, the subscriber and the other four packages (arith, fib, rev, enc) responsible to actually solve the tasks.

Wasn't in the scope of implementation to create a real client sending tasks to be solved by the systems. So, the publisher package must be seeing only as an abstraction of that.

The subscriber package is a middleware entity between the publisher and the solver packages what justify the separation between those packages and the subscriber and the fact of the resulting calculation goes first to the subscriber and only then back to the publisher.
