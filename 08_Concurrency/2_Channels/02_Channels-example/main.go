package main 

import "fmt"

func hello(done chan bool) {

	fmt.Println("Hello world goroutine")

	done <- true 

}

func main() {

	done := make(chan bool)
	go hello(done)

	<-done
	fmt.Println("main goroutine")
}

//In the above program, we create a done bool channel in line no. 15 and pass it as a parameter to the hello Goroutine. In line no. 18 we are receiving data from the done channel. This line of code is blocking which means that until some Goroutine writes data to the done channel, the control will not move to the next line of code. Hence this eliminates the need for the time.Sleep which was present in the original program to prevent the main Goroutine from exiting.

// The line of code <-done receives data from the done channel but does not use or store that data in any variable. This is perfectly legal.

// Now we have our main Goroutine blocked waiting for data on the done channel. The hello Goroutine receives this channel as a parameter, prints Hello world goroutine and then writes to the done channel. When this write is complete, the main Goroutine receives the data from the done channel, it is unblocked and then the text main function is printed.

// This program outputs

// Hello world goroutine
// main function

