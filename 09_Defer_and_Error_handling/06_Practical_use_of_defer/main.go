package main 

import (

	"fmt"
	"sync"

)

type rect struct {
	length int 
	width int 
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return 
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return 
	}
	area := r.length * r.width 
	fmt.Printf("rect %v's area %d\n", r , area)

}

func main() {
	var wg sync.WaitGroup

	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v:= range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All goroutine finished executing")
}



// In the program above, we have removed the 3 wg.Done() calls in the above program and replaced it with a single defer wg.Done() call in line no. 14. This makes the code more simple and readable.

// There is one more advantage of using defer in the above program. Let’s say we add another return path to the area method using a new if condition. If the call to wg.Done() was not deferred, we have to be careful and ensure that we call wg.Done() in this new return path. But since the call to wg.Done() is defered, we need not worry about adding new return paths to this method.

