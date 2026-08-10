package main 

//for loop is the only looping statement available in Go. It’s possible to use a variation of the for loop to achieve the functionality of a while loop.

// The program below prints all even numbers from 0 to 10.

import (
	"fmt"
)

func main () {
	i := 0
	for ; i <= 10 ; { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
}

// As we already know all the three components of the for loop namely initialisation, condition and post are optional. In the above program, initialisation and post are omitted. i is initialized to 0 outside the for loop. The loop will be executed as long as i <= 10. i is incremented by 2 inside the for loop. The above program outputs 0 2 4 6 8 10 .