package main 

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"

}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 *time.Second)
	select{
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

}

// In the program above, the server1 and server2 go routines are called in line no. 18 and 19 respectively. Then the main program sleeps for 1 second in line no. 20. When the control reaches the select statement in line no. 21, server1 would have written from server1 to the output1 channel and server2 would have written from server2 to the output2 channel and hence both the cases of the select statement are ready to be executed. If you run this program multiple times, the output will vary between from server1 or from server2 depending on which case is chosen at random.