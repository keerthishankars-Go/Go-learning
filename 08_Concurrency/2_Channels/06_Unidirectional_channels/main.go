package main  

// All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them. It is also possible to create unidirectional channels, that is channels that only send or receive data.

//import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {

	sendch := make(chan<- int)
	go sendData(sendch)
	//fmt.Println(<-sendch)

// invalid operation: cannot receive from send-only channel chan<- int sendch (variable of type chan<- int)compilerInvalidReceive
// var sendch chan<- int

}

// In the above program, we create send only channel sendch in line no. 13. chan<- int denotes a send only channel as the arrow is pointing to chan. We try to receive data from a send only channel in line no. 15. This is not allowed and when the program is run, the compiler will complain stating,

// ./prog.go:12:14: invalid operation: <-sendch (receive from send-only type chan<- int)

// All is well but what is the point of writing to a send only channel if it cannot be read from!

// This is where channel conversion comes into use. It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.