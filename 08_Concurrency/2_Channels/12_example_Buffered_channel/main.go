package main

//The values to the channel are written in a concurrent Goroutine and read from the main Goroutine..

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "into ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

// In the program above, a buffered channel ch of capacity 2 is created in line no. 19 of the main Goroutine and passed to the write Goroutine in line no. 20. Then the main Goroutine sleeps for 2 seconds. During this time, the write Goroutine is running concurrently. The write Goroutine has a for loop which writes numbers from 0 to 4 to the ch channel. The capacity of this buffered channel is 2 and hence the write Goroutine will be able to write values 0 and 1 to the ch channel immediately and then it blocks until at least one value is read from ch channel. So this program will print the following 2 lines immediately.

// successfully wrote 0 to ch
// successfully wrote 1 to ch

// After printing the above two lines, the writes to the ch channel in the write Goroutine are blocked until someone reads from the ch channel. Since the main Goroutine sleeps for 2 seconds before starting to read from the channel, the program will not print anything for the next 2 seconds. The main Goroutine wakes up after 2 seconds and starts reading from the ch channel using a for range loop in line no. 22, prints the read value and then sleeps for 2 seconds again and this cycle continues until the ch is closed. So the program will print the following lines after 2 seconds,

// read value 0 from ch
// successfully wrote 2 to ch

// This will continue until all values are written to the channel and it is closed in the write Goroutine. The final output would be,

// successfully wrote 0 to ch
// successfully wrote 1 to ch
// read value 0 from ch
// successfully wrote 2 to ch
// read value 1 from ch
// successfully wrote 3 to ch
// read value 2 from ch
// successfully wrote 4 to ch
// read value 3 from ch
// read value 4 from ch
