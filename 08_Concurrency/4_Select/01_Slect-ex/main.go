package main  

import (
	"fmt"
	"time"

)

func server1(ch chan string){
	time.Sleep(6 *time.Second)
	ch <- "from server1"
}

func server2(ch chan string){
	time.Sleep(3 *time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select{
	case s1 := <- output1:
		fmt.Println(s1)
	case s2 := <- output2:
		fmt.Println(s2)
	}

}

// In the program above, the server1 function in line no. 8 sleeps for 6 seconds then writes the text from server1 to the channel ch. The server2 function in line no. 12 sleeps for 3 seconds and then writes from server2 to the channel ch.

// The main function calls the go Goroutines server1 and server2 in line nos. 20 and 21 respectively.

// In line no. 22, the control reaches the select statement. The select statement blocks until one of its cases is ready. In our program above, the server1 Goroutine writes to the output1 channel after 6 seconds whereas the server2 writes to the output2 channel after 3 seconds. So the select statement will block for 3 seconds and will wait for server2 Goroutine to write to the output2 channel. After 3 seconds, the program prints,

// from server2