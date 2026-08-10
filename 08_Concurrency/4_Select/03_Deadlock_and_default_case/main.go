package main

func main() {
	ch := make(chan string)

	select {
	case <-ch:

	}
}

// In the program above, we have created a channel ch in line no. 4. We try to read from this channel inside the select in line no. 6. The select statement will block forever since no other Goroutine is writing to this channel and hence will result in deadlock. This program will panic at runtime with the following message,

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         C:/Projects/go-learning/08_Concurrency/4_Select/03_Deadlock_and_default_case/main.go:7 +0x25
// exit status 2

// If a default case is present, this deadlock will not happen since the default case will be executed when no other case is ready.
