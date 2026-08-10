package main

// To understand worker pools, we need to first know about WaitGroup as it will be used in the implementation of Worker pool.

// A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing. Let’s say we have 3 concurrently executing Goroutines spawned from the main Goroutine. The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished using WaitGroup.

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()

}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All oroutines finished executing..")
}

// WaitGroup is a struct type and we are creating a zero value variable of type WaitGroup in line no.18. The way WaitGroup works is by using a counter. When we call Add on the WaitGroup and pass it an int, the WaitGroup’s counter is incremented by the value passed to Add. The way to decrement the counter is by calling Done() method on the WaitGroup. The Wait() method blocks the Goroutine in which it’s called until the counter becomes zero.

// In the above program, we call wg.Add(1) in line no. 20 inside the for loop which iterates 3 times. So the counter now becomes 3. The for loop also spawns 3 process Goroutines and then wg.Wait() called in line no. 23 makes the main Goroutine to wait until the counter becomes zero. The counter is decremented by the call to wg.Done in the process Goroutine in line no. 13. Once all the 3 spawned Goroutines finish their execution, that is once wg.Done() has been called three times, the counter will become zero, and the main Goroutine will be unblocked.

// It is important to pass the pointer of wg in line no. 21. If the pointer is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.

// This program outputs.

// started Goroutine  2
// started Goroutine  0
// started Goroutine  1
// Goroutine 0 ended
// Goroutine 2 ended
// Goroutine 1 ended
// All go routines finished executing
