package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()

	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)

	}
	w.Wait()
	fmt.Println("final value of x ", x)
}

// Mutex means:

// Mutual Exclusion..

// Meaning:

// Only one goroutine is allowed inside this section at a time.

// Mutex is a struct type and we create a zero valued variable m of type Mutex in line no. 15. In the above program we have changed the increment function so that the code which increments x x = x + 1 is between m.Lock() and m.Unlock(). Now this code is void of any race conditions since only one Goroutine is allowed to execute this piece of code at any point in time.

// Now if this program is run, it will output

// final value of x 1000

// It is important to pass the address of the mutex in line no. 18. If the mutex is passed by value instead of passing the address, each Goroutine will have its own copy of the mutex and the race condition will still occur.