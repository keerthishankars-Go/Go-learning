package main  

// Similar to slices, maps are reference types. When a map is assigned to a new variable, they both point to the same underlying data structure. Hence changes made in one will reflect in the other.

import (
	"fmt"
)

func main () {
	employeeSalary := map[string]int {
		"Mike" : 60000,
		"Peter" : 90000,
		"Keerthi" : 160000,

	}
	fmt.Println("Original employee salary:", employeeSalary)
	modified:= employeeSalary
	modified["Keerthi"] = 200000
	fmt.Println("The modified employee salary is:", employeeSalary)
}

//  employeeSalary is assigned to modified. In the next line, the salary of Keerthi is changed to 200000 in the modified map. Keerthi’s salary will now be 200000 in employeeSalary too.