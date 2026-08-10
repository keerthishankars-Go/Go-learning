package main   

// same example can be written with range loop..

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)

	for n := range ch {
		fmt.Println("Received", n)
	}
}

// The for range loop in line no. 12 of the program above will read all the values written to the channel and will quit once there are no more values to read since the channel is already closed.

// This program will print,

// Received: 5
// Received: 6