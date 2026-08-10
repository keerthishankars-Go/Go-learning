package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		var pErr *os.PathError
		if errors.As(err, &pErr){
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return

	}
	fmt.Println(f.Name(), "opened successfully..")
}

// In the above program, we first check whether the error is not nil in line no. 11 and then we use the As function in line no. 13 to convert err to *os.PathError. If the conversion is successful, As will return true. Then we print the path using pErr.Path in line no. 14.

// If you are wondering why pErr is a pointer, the reason is, the error interface is implemented by the pointer of PathError and hence pErr is a pointer. The below code shows that *PathError implements the error interface.

// func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
// The As function requires the second argument to be a pointer to the type that implements the error. Hence we pass &perr.

// This program outputs,

// Failed to open file at path test.txt
// In case the underlying error is not of type *os.PathError, the control will reach line no. 17 and a generic error message will be printed.

