package main

// It is possible to declare structs without creating a new data type. These types of structs are called anonymous structs.

import (
	"fmt"
)

func main() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Keerthi",
		lastName:  "S",
		age:       27,
		salary:    25,
	}

	fmt.Println("Employee 3 is", emp3)

}

// In line no 10. of the above program, an anonymous struct variable emp3 is defined. As we have already mentioned, this struct is called anonymous because it only creates a new struct variable emp3 and does not define any new struct type like named structs.
