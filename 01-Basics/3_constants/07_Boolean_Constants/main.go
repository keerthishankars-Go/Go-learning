// Boolean constants are no different from string constants. They are two untyped constants true and false.

package main  

func main () {

	const trueConst = true 
	type myBool bool 
	var defaultBool = trueConst //allowed
	var customBool myBool = trueConst //allowed
	defaultBool = customBool // not allowed

}

