// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.
package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string // Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make([]string, 3)                                  //To create a slice with non-zero length, use the builtin make.
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) //Here we make a slice of strings of length 3 (initially zero-valued). By default a new slice’s capacity is equal to its length;
	//if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.

	//We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d") //the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) //Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5] //Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
	fmt.Println("sl1:", l)

	l = s[:5] //This slices up to (but excluding) s[5].
	fmt.Println("sl2:", l)

	l = s[2:] //And this slices up from (and including) s[2].
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} //We can declare and initialize a variable for slice in a single line as well.
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	//Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
