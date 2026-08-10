package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

// The above program starts two Goroutines in line nos. 21 and 22. These two Goroutines now run concurrently. The numbers Goroutine sleeps initially for 250 milliseconds and then prints 1, then sleeps again and prints 2 and the same cycle happens till it prints 5. Similarly the alphabets Goroutine prints alphabets from a to e and has 400 milliseconds of sleep time. The main Goroutine initiates the numbers and alphabets Goroutines, sleeps for 3000 milliseconds and then terminates.

// This program outputs

// 1 a 2 3 b 4 c 5 d e main terminated