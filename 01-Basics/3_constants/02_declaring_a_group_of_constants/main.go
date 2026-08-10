// There is also another syntax to define a group of constants using a single statement.

package main

import "fmt"

func main () {

	const (
		retryLimit = 4
		httpMethod = "GET"
		
	)

	fmt.Println(retryLimit)
	fmt.Println(httpMethod)
}

// We have declared 2 constants retryLimit and httpMethod.
// Constants, as the name indicate, cannot be reassigned again to any other value.