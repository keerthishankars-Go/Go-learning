package main

import "fmt"

// If a default case is present, this deadlock will not happen since the default case will be executed when no other case is ready. The program above is rewritten with a default case below.

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("Default case executed...")
	}
}

