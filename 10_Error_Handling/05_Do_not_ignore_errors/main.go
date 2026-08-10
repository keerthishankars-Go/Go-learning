package main

// Never ever ignore an error. Ignoring errors is inviting for trouble. Let me rewrite the example which lists the name of all files that match a pattern ignoring errors.

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("[")

	fmt.Println("matched files", files)
}

// We already know from the previous example that the pattern is invalid. I have ignored the error returned by the Glob function by using the _ blank identifier in line no. 9. I simply print the matched files in line no. 10. This program will print,

// matched files []

// Since we ignored the error, the output seems as if no files have matched the pattern but actually the pattern itself is malformed. So never ignore errors.
