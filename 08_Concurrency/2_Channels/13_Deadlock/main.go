package main  

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "Keerthi"
	ch <- "shankar"
	ch <- "S"

	fmt.Println(<- ch)
	fmt.Println(<- ch)
}

// In the program above, we write 3 strings to a buffered channel of capacity 2. When the control reaches the third write in line no. 11, the write is blocked since the channel has exceeded its capacity. Now some Goroutine must read from the channel in order for the write to proceed, but in this case, there is no concurrent routine reading from this channel. Hence there will be a deadlock and the program will panic at run time with the following message,

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//         C:/Projects/go-learning/08_Concurrency/16_Deadlock/main.go:9 +0x58
// exit status 2 