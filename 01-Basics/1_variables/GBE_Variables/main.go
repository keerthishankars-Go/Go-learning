// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
package variables

import "fmt"

func GBE() {

	var a = "initial" //var declares 1 or more variables.
	fmt.Println(a)

	var b, c int = 1, 2 //You can declare multiple variables at once.
	fmt.Println(b, c)

	var d = true //Go will infer the type of initialized variables.
	fmt.Println(d)

	var e int //Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	fmt.Println(e)

	f := "apple" //The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
	fmt.Println(f)
}
