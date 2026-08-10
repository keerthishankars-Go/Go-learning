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
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return 
	}
	if r.width < 0 {
		fmt.Printf("The rect %v's width should be greater than zero\n", r)
		wg.Done()
		return 
	}
	area := r.length * r.width 
	fmt.Printf("rect %v's area %d\n", r, area )
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5 , -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All goroutine finished executing..")
}

// Defer is used in places where a function call should be executed irrespective of the code flow. Let’s understand this with the example of a program which makes use of WaitGroup. We will first write the program without using defer and then we will modify it to use defer and understand how useful defer is.

// In the program above, we have created a rect struct in line no. 8 and a method area on rect in line no. 13 which calculates the area of the rectangle. This method checks whether the length and width of the rectangle is less than zero. If it is so, it prints a corresponding error message else it prints the area of the rectangle.

// The main function creates 3 variables r1, r2 and r3 of type rect. They are then added to the rects slice in line no. 34. This slice is then iterated using a for range loop and the area method is called as a concurrent Goroutine in line no. 37. The WaitGroup wg is used to ensure that the main function is waiting until all Goroutines finish executing. This WaitGroup is passed to the area method as an argument and the area method calls wg.Done() in line nos. 16, 21 and 26 to notify the main function that the Goroutine is done with its job. If you notice closely, you can see that these calls happen just before the area method returns. wg.Done() should be called before the method returns irrespective of the path the code flow takes and hence these calls can be effectively replaced by a single defer call.