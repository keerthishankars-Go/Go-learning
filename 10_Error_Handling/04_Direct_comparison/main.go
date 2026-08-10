package main

import (
	"errors"
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("[")
	if err != nil {
		if errors.Is(err, filepath.ErrBadPattern) {
			fmt.Println("Bad pattern error: ", err)
			return
		}
		fmt.Println("Generic error: ", err)
		return

	}
	fmt.Println("matched files: ", files)
}

// In the program above we search for files of pattern [ which is a malformed pattern. We check whether the error is not nil. To get more information about the error, we directly compare it to filepath.ErrBadPattern in line. no 11 using the Is function. Similar to As, the Is function works on an error chain. We will learn more about this in our next tutorial. For the purposes of this tutorial, the Is function can be thought of as returning true if both the errors passed to it are the same.

// The Is returns true in line no. 12 since the error is due to a malformed pattern. This program will print,

// Bad pattern error: syntax error in pattern
