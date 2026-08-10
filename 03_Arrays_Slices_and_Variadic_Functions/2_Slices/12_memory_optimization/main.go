package main 

import (
	"fmt"
)

func countries() []string {
	countries := []string{"India", "Japan", "USA", "China", "Russia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy

}

func main () {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

// Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected. This might be of concern when it comes to memory management. Let’s assume that we have a very large array and we are interested in processing only a small part of it. Henceforth we create a slice from that array and start processing the slice. The important thing to be noted here is that the array will still be in memory since the slice references it.

// One way to solve this problem is to use the copy function func copy(dst, src []T) int to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.

//  Now countries array can be garbage collected since neededCountries is no longer referenced.