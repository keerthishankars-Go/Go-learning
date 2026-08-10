package main

// The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the make function.

// The length of the buffered channel is the number of elements currently queued in it.

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "Keerthi"
	ch <- "shankar"
	
	fmt.Println("Capacity", cap(ch))
	fmt.Println("Length", len(ch))
	fmt.Println("Read from channel", <- ch)
	fmt.Println("new length is", len(ch))
}

// In the program above, the channel is created with a capacity of 3, that is, it can hold 3 strings. We then write 2 strings to the channel in line nos. 9 and 10 respectively. Now the channel has 2 strings queued in it and hence its length is 2. In line no. 13, we read a string from the channel. Now the channel has only one string queued in it and hence its length becomes 1. This program will print,

// capacity is 3
// length is 2
// read value naveen
// new length is 1