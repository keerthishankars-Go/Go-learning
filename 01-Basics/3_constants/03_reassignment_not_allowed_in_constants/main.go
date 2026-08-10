package main

func main() {
	const a = 55 // allowed

	//a = 45  // not allowed
}

// we are trying to assign another value 45 to a. This is not allowed since a is a constant. This program will fail to run with compilation error cannot assign to a (neither addressable nor a map index expression).
