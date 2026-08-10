package main 

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.Name(), "Opened successfully")
}

// Errors indicate any abnormal condition occurring in the program. Let’s say we are trying to open a file and the file does not exist in the file system. This is an abnormal condition and it’s represented as an error.

// Errors in Go are plain old values. Just like any other built-in type such as int, float64, … error values can be stored in variables, passed as parameters to functions, returned from functions, and so on.

// Errors are represented using the built-in error type. 

//==========================================================================//

// In line no. 9 of the program above, we are trying to open the file at path /test.txt(which will obviously not exist in the playground). The Open function of the os package has the following signature,

// *func Open(name string) (File, error)

// If the file has been opened successfully, then the Open function will return the file handler and error will be nil. If there is an error while opening the file, a non-nil error will be returned.

// If a function or method returns an error, then by convention it has to be the last value returned from the function. Hence the Open function returns error as the last value.

// The idiomatic way of handling errors in Go is to compare the returned error to nil. A nil value indicates that no error has occurred and a non-nil value indicates the presence of an error. In our case, we check whether the error is not nil in line no. 10. If it is not nil, we simply print the error and return from the main function.

// Running this program will print

// open /test.txt: No such file or directory