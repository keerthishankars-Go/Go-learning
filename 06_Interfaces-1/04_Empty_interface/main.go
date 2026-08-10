package main 

// An interface that has zero methods is called an empty interface. It is represented as interface{}. Since the empty interface has zero methods, all types implement the empty interface.

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello world"
	describe(s)

	i := 55
	describe(i)

	strt := struct{
		name string
	} {
		name: "Keerthi",
	}
	describe(strt)
}

//In the program above, in line no.7, the describe(i interface{}) function takes an empty interface as an argument and hence any type can be passed.

// We pass string, int and struct to the describe function in line nos. 13, 15 and 21 respectively. This program prints,

// Type = string, value = Hello World
// Type = int, value = 55
// Type = struct { name string }, value = {Naveen R}