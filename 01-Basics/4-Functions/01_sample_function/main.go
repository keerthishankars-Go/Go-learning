// The parameters and return type are optional in a function.

// func functionname() {
// }

// Let’s write a function that takes the price of a single product and the quantity as input parameters and returns the total price by multiplying these two values.

package main // executable project: If you want to run using go run, package must be main

import "fmt" //Standard library package: Used for printing, formatting output

func calculateBill(price int, quantity int) int {

	var totalPrice = price * quantity //Declares a variable, multiplies inputs

	return totalPrice
}

func main() {

	price, quantity := 35, 50
	totalPrice := calculateBill(price, quantity) // calling a function : calculateBill is called

	//Returned value stored in totalPrice

	fmt.Println("Total price is", totalPrice)
}

//main is special : Program execution starts here , Go runtime calls it automatically

//STRICT RULE:

// func main() {
// }

// ❌ No parameters
// ❌ No return value


// Calling a function:

// totalPrice := calculateBill(price, quantity)

// What happens

// calculateBill is called

// price → goes into price

// quantity → goes into quantity

// Function returns an int

// Returned value stored in totalPrice

// 📌 Rule

// "Function call must match parameter order & types"

// 📌 Rule

// Use Println for simple debugging

// Use Printf for formatted output