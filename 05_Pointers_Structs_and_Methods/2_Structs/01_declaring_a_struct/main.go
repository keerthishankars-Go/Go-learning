package main

// A struct is a user-defined type that represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than having each of them as separate variables.

type Employee struct {
	firstName string
	lastName  string
	age       int
}

// The above snippet declares a struct type Employee with fields firstName, lastName and age. The above Employee struct is called a named struct because it creates a new data type named Employee using which Employee structs can be created.

// This struct can also be made more compact by declaring fields that belong to the same type in a single line followed by the type name. In the above struct firstName and lastName belong to the same type string and hence the struct can be rewritten as:

type Employeee struct {
	firstName, lastName string
	age                 int
}

// Although the above syntax saves a few lines of code, it doesn’t make the field declarations explicit. So I would not recommend using the above syntax.