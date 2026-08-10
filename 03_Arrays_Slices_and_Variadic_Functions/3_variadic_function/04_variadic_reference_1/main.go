package main 

import (
	"fmt"
)

func change (s ...string) {
	s[0] = "Go"
}

func main () {
	welcome := []string{"Hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

//In line no. 13 of the program above, we are using the syntactic sugar ... and passing the slice as a variadic argument to the change function.

// As we have already discussed, if ... is used, the welcome slice itself will be passed as an argument without a new slice being created. Hence welcome will be passed to the change function as argument.

// Inside the change function, the first element of the slice is changed to Go