package main  

// The for range form of the for loop can be used to receive values from a channel until it is closed.

// Let’s rewrite the program above using a for range loop.

import "fmt"

func producer(chn1 chan int) {
	for i:=0; i < 10; i ++ {
		chn1 <- i
	}
	close(chn1)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

// The for range loop in line no. 19 receives data from the ch channel until it is closed. Once ch is closed, the loop automatically exits.