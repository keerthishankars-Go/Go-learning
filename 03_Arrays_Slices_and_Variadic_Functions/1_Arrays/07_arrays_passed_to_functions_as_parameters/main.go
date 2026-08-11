package main  

// when arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func main () {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value..
	fmt.Println("after passing to function ", num)

}

// In the above program in line no. 16, the array num is actually passed by value to the function changeLocal and hence will not change because of the function call..