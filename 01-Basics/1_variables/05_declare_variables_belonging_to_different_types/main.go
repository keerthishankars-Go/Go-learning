// declare variables belonging to different types in a single statement. The syntax for doing that is:

// var (
//       name1 = initialvalue1
//       name2 = initialvalue2
// )

package main

import "fmt"

func main() {
	//Here we declare a variable name of type string, age and height of type int
	var (
		name   = "keerthi"
		age    = 27
		height int
	)

	fmt.Println("Name is", name)
	fmt.Println("Age is", age)
	fmt.Println("Height is", height)
}
