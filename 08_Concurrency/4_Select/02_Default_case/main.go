package main  

// The default case in a select statement is executed when none of the other cases is ready. This is generally used to prevent the select statement from blocking.

import (
	"fmt"
	"time"

)

func process(ch chan string) {
	time.Sleep(10500 *time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 *time.Millisecond)
		select{
		case v := <- ch :
			fmt.Println("Received value:", v)
		default:
			fmt.Println("no value received")
		}
	}
}

// In the program above, the process function in line no. 8 sleeps for 10500 milliseconds (10.5 seconds) and then writes process successful to the ch channel. This function is called concurrently in line no. 15 of the program.

// After calling the process Goroutine concurrently, an infinite for loop is started in the main Goroutine. The infinite loop sleeps for 1000 milliseconds (1 second) during the start of each iteration and then performs a select operation. During the first 10500 milliseconds, the first case of the select statement namely case v := <-ch: will not be ready since the process Goroutine will write to the ch channel only after 10500 milliseconds. Hence the default case will be executed during this time and the program will print no value received 10 times.

// After 10.5 seconds, the process Goroutine writes process successful to ch in line no. 10. Now the first case of the select statement will be executed and the program will print received value: process successful and then it will terminate. This program will output,

// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// received value:  process successful