package main

// Channels can be thought of as pipes using which Goroutines communicate. Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the other end using channels.

// Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport. No other type is allowed to be transported using the channel.

// chan T is a channel of type T

// The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
	}
	a = make(chan int)
	fmt.Printf("Type of a is %T", a)
}


// The channel a declared in line no. 14 is nil as the zero value of a channel is nil. Hence the statements inside the if condition are executed and the channel is defined. a in the above program is a int channel. This program will output,

// channel a is nil, going to define it
// Type of a is chan int
// As usual, the short hand declaration is also a valid and concise way to define a channel.