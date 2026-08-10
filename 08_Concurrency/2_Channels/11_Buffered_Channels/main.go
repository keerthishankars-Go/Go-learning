package main

// All the channels we discussed in the previous tutorial were basically unbuffered. As we discussed in the channels tutorial in detail, sends and receives to an unbuffered channel are blocking.

// It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

// Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

// ch := make(chan type, capacity)

// capacity in the above syntax should be greater than 0 for a channel to have a buffer. The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating channels in the previous tutorial.

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Keerthi"
	ch <- "shankar"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

// In the program above, in line no. 16 we create a buffered channel with a capacity of 2. Since the channel has a capacity of 2, it is possible to write 2 strings into the channel without being blocked. We write 2 strings to the channel in line no. 17 and 18 and the channel does not block. We read the 2 strings written in line nos. 20 and 21 respectively..
