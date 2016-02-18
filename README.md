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

```
git clone github.com/alesr/octopus-distributor
cd octopus-distributor/
go build
./octopus-distributor
```
To set the number of requests to simulate.
```
./octopus-distributor -request=10000
```
To write requests on debug.out under application root directory.
```
./octopus-distributor -debug=true
```


### How it works

In order to see our program working, we need some entity to simulate a high volume of tasks to be handled by the distributor.

The Publisher package assembles random tasks of four different types: arithmetic for basic mathematic operations, Fibonacci, reverse for mirrored text and encode for bcrypt hash values.

After building the tasks, it sends these messages through the request channel.

The Subscriber package at its side registers the publisher and starts receiving from the request channel. Basically, the Subscriber could register any Publisher interested in having tasks handled by the system by sending them trough the same request channel.

After receiving the request the Subscriber checks if the task can be handled. And if so, distributes it to be solved by one of the four external packages responsible for executing that particular task.

The four mentioned packages are the Arithmetic, Fibonacci, Reverse and Encode. Each one receives tasks through their own channel and giving back the resulting calculation through the result's channel.

Once the Subscriber receives requests from the results channel it sends the ordered results back to the Publisher. The Receiver function can then output the response putting an end to the cycle.

#### Some important considerations in this process to justify my implementation for the test.

Channels are typed conduits that behave like FIFO queues, hence, they preserve the order of the items that are sent to them and block from receiving new items while the other side is not ready to handle them. Also, it is important to note that each time the Subscriber receives new requests it launches a new goroutine to do the calculation. Therefore, the program solves the tasks concurrently and due to this reason there is no waiting any requests in line even though there is a line still.

Another thing is that the program is divided into three major components. The publisher, the subscriber and the other four packages (arith, fib, rev, enc) responsible to actually solve the tasks.

It was out of scope of implementation to create a real client sending tasks to be solved by the systems and as a result, the publisher package is advised to be seen only as an abstraction of that.

The subscriber package is a middleware entity between the publisher and the solver packages what justify the separation between those packages and the subscriber and the fact of the resulting calculation goes first to the subscriber and only then back to the publisher.
