package main 

// The length of the slice is the number of elements in the slice. The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.

import(
	"fmt"
)

func main () {
	fruitarray := [...]string{"apple", "orange", "Butterfruit", "Pineapple", "grapes", "mango", "banana"}
	fruitslice := fruitarray[1:3]
	fmt.Println("The slice is", fruitslice)
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
}

// In the above program, fruitslice is created from indexes 1 and 2 of the fruitarray. Hence the length of fruitslice is 2.

// The length of the fruitarray is 7. fruiteslice is created from index 1 of fruitarray. Hence the capacity of fruitslice is the no of elements in fruitarray starting from index 1 i.e from orange and that value is 6. Hence the capacity of fruitslice is 6. The program prints length of slice 2 capacity 6.