package main  

// It is also possible to append one slice to another using the ... operator.

import (
	"fmt"
)

func main () {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"mangoes", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food", food)
}

// The above program food is created by appending fruits to veggies..