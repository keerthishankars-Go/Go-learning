package main  

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {

	ch <- true 
	x = x + 1

	<- ch 

	wg.Done()

}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i ++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("Final output x ", x)
}

// In the program above, we have created a buffered channel of capacity 1 and this is passed to the increment Goroutine in line no. 18. This buffered channel is used to ensure that only one Goroutine access the critical section of code which increments x. This is done by passing true to the buffered channel in line no. 8 just before x is incremented. Since the buffered channel has a capacity of 1, all other Goroutines trying to write to this channel are blocked until the value is read from this channel after incrementing x in line no. 10. Effectively this allows only one Goroutine to access the critical section.

// This program also prints

// final value of x 1000


// In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine should access the critical section of code.

// In the case of the problem which we solved above, I would prefer to use mutex since this problem does not require any communication between the goroutines. Hence mutex would be a natural fit.