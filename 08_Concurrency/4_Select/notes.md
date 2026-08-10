The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random. The syntax is similar to switch except that each of the case statements will be a channel operation. Let’s dive right into some code for better understanding.

Practical use of select
The reason behind naming the functions in the above program as server1 and server2 is to illustrate the practical use of select.

Let’s assume we have a mission critical application and we need to return the output to the user as quickly as possible. The database for this application is replicated and stored in different servers across the world. Assume that the functions server1 and server2 are in fact communicating with 2 such servers. The response time of each server is dependant on the load on each and the network delay. We send the request to both the servers and then wait on the corresponding channels for the response using the select statement. The server which responds first is chosen by the select and the other response is ignored. This way we can send the same request to multiple servers and return the quickest response to the user :).

==============================================================================

Similarly, the default case will be executed even if the select has only nil channels.

package main

import "fmt"

func main() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")

	}
}

In the program above ch is nil and we are trying to read from ch in the select in line no. 8. If the default case was not present, the select would have blocked forever and caused a deadlock. Since we have a default case inside the select, it will be executed and the program will print,

default case executed

==============================================================================

**Gotcha - Empty select**

package main

func main() {
	select {}
}

What do you think will be the output of the program above?

We know that the select statement will block until one of its cases is executed. In this case, the select statement doesn’t have any cases and hence it will block forever resulting in a deadlock. This program will panic with the following output,

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
main.main()
	/tmp/sandbox246983342/prog.go:4 +0x25