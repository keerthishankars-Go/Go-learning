package main  

import "fmt"

func digits(number int, dchn1 chan int) {
	for number != 0 {
		digit := number % 10 
		dchn1 <- digit 
		number /= 10
	}
	close(dchn1)
}

func calcSquares(number int, squareOp chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit 

	}
	squareOp <- sum
}

func calcCubes(number int, cubeOp chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeOp <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <- sqrch, <- cubech
	fmt.Println("sum of squares and cubes = ", squares+cubes)
}

// Background: If you take a closer look at the program you can notice that the code which finds the individual digits of a number is repeated in both calcSquares function and calcCubes function. We will move that code to its own function and call it concurrently.

// The digits function in the program above now contains the logic for getting the individual digits from a number and it is called by both calcSquares and calcCubes functions concurrently. Once there are no more digits in the number, the channel is closed in line no. 13. The calcSquares and calcCubes Goroutines listen on their respective channels using a for range loop until it is closed. The rest of the program is the same. This program will also print

// Final output 1536