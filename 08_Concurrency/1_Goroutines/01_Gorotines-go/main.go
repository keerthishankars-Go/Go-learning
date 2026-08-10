package main 

// Prefix the function or method call with the keyword go and you will have a new Goroutine running concurrently.

import "fmt"

func hello() {
	fmt.Println("Hello world goroutine")
}

func main() {
	go hello()
	fmt.Println("main function")
}

//In line no. 12, go hello() starts a new Goroutine. Now the hello() function will run concurrently along with the main() function. The main function runs in its own Goroutine and it’s called the main Goroutine.

// Run this program and you will have a surprise!

// This program only outputs the text main function. What happened to the Goroutine we started? We need to understand the two main properties of goroutines to understand why this happens.

// When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.
// The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.

// After the call to go hello() in line no. 12, the control returned immediately to the next line of code without waiting for the hello goroutine to finish and printed main function. Then the main Goroutine terminated since there is no other code to execute and hence the hello Goroutine did not get a chance to run.