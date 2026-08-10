package main  

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world of goroutines")
}

func main() {
	go hello()
	time.Sleep(2 * time.Second)
	fmt.Println("main function")
}

// In line no.14 of the program above, we have called the Sleep method of the time package which sleeps the go routine in which it is being executed. In this case the main goroutine is put to sleep for 1 second. Now the call to go hello() has enough time to execute before the main Goroutine terminates. This program first prints Hello world goroutine, waits for 1 second and then prints main function.

// This way of using sleep in the main Goroutine to wait for other Goroutines to finish their execution is a hack we are using to understand how Goroutines work. Channels can be used to block the main Goroutine until all other Goroutines finish their execution. We will discuss channels in the next tutorial.

