package main  

// When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.

import (
	"fmt"
)

func main () {
	numa := [3]int{31, 44, 65}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 231
	fmt.Println("array after modification to slice nums2", numa)
}

// In numa[:] the start and end values are missing. The default values for start and end are 0 and len(numa) respectively. Both slices nums1 and nums2 share the same array.

// From the output, it’s clear that when slices share the same array. The modifications made to the slice are reflected in the array.