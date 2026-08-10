// In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.
package main

import "fmt"

func main() {
	var a [5]int //Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
	fmt.Println("emp:", a)

	a[4] = 100 //We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) //The builtin len returns the length of an array.

	b := [5]int{1, 2, 3, 4, 5} //Use this syntax to declare and initialize an array in one line.
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} //You can also have the compiler count the number of elements for you with ...
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} //If you specify the index with :, the elements in between will be zeroed.
	fmt.Println("idx:", b)

	var twoD [2][3]int //Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{ //You can create and initialize multi-dimensional arrays at once too.
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)

}

//Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.
