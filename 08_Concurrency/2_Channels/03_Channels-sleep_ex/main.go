package main 

import ( 
	"fmt"
	"time"

)

func hello(done chan bool) {
	fmt.Println("hello goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello goroutine ia awake and going to write to done")
	done <- true 
}

func main() {

	done := make(chan bool)
	fmt.Println("main going to call hello goroutine")
	go hello(done)
	<- done
	fmt.Println("main received data")

}


// In the above program, we have introduced a sleep of 4 seconds to the hello function in line no. 11.

// This program will first print Main going to call hello go goroutine. Then the hello Goroutine will be started and it will print hello go routine is going to sleep. After this is printed, the hello Goroutine will sleep for 4 seconds and during this time main Goroutine will be blocked since it is waiting for data from the done channel in line no. 21 <-done. After 4 seconds hello go routine awake and going to write to done will be printed followed by Main received data.