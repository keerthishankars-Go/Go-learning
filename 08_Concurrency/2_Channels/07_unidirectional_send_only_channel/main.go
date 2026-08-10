package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	ch1 := make(chan int)
	go sendData(ch1)
	fmt.Println(<-ch1)
}

// In line no. 10 of the program above, a bidirectional channel chnl is created. It is passed as a parameter to the sendData Goroutine in line no. 11. The sendData function converts this channel to a send only channel in line no. 5 in the parameter sendch chan<- int. So now the channel is send only inside the sendData Goroutine but it’s bidirectional in the main Goroutine. This program will print 10 as the output.
