package main  

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()

}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i ++ {
		w.Add(1)
		go increment(&w)

	}
	w.Wait()
	fmt.Println("Final value of x ", x)

}



// In the program above, the increment function in line no. 7 increments the value of x by 1 and then calls Done() on the WaitGroup to notify its completion.

// We spawn 1000 increment Goroutines from line no. 15 of the program above. Each of these Goroutines run concurrently and the race condition occurs when trying to increment x is line no. 8 as multiple Goroutines try to access the value of x concurrently.

// Please run this program in your local as the playground is deterministic and the race condition will not occur in the playground. Run this program multiple times in your local machine and you can see that the output will be different for each time because of race condition. Some of the outputs which I encountered are final value of x 941, final value of x 928, final value of x 922 and so on.