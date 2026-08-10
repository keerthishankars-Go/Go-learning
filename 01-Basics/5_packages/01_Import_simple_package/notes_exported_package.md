Exported Names

We capitalized the function Calculate in the Simple interest package.
This has a special meaning in Go.

Any variable or function which starts with a capital letter are exported names in go.

Only exported functions and variables can be accessed from other packages. In our case, we want to access Calculate function from the main package. Hence this is capitalized.

If the function name is changed from Calculate to calculate in simpleinterest.go, and if we try to call the function using simpleinterest.calculate(p, r, t) in main.go, the compiler will error

# learnpackage

./main.go:13:8: undefined: simpleinterest.calculate
Hence if you want to access a function outside of a package, it should be capitalized.

==============================================================

init function

Each package in Go can contain an init function.
The init function must not have any return type and it must not have any parameters.
The init function cannot be called explicitly in our source code.

It will be called automatically when the package is initialized. The init function has the following syntax

func init() {
}

---

The order of initialization of a package is as follows

Package level variables are initialised first

init function is called next.
A package can have multiple init functions (either in a single file or distributed across multiple files) and they are called in the order in which they are presented to the compiler.
If a package imports other packages, the imported packages are initialized first.

A package will be initialized only once even if it is imported from multiple packages.

================================================

PACKAGE LEVEL VARIABLES

package main

import (
"fmt"
"learnpackage/simpleinterest" //importing custom package
"log"
)

var p, r, t = 5000.0, 10.0, 1.0 //package level variables

/\*

- init function to check if p, r and t are greater than zero
  \*/
  func init() {
  fmt.Println("Main package initialized")
  if p < 0 {
  log.Fatal("Principal is less than zero")
  }
  if r < 0 {
  log.Fatal("Rate of interest is less than zero")
  }
  if t < 0 {
  log.Fatal("Duration is less than zero")
  }
  }

func main() {
fmt.Println("Simple interest calculation")
si := simpleinterest.Calculate(p, r, t)
fmt.Println("Simple interest is", si)
}

The following are the changes made to main.go:

p, r and t variables are moved to package level from the main function level.
An init function has been added. The init function prints a log and terminates the program execution if either the principal, rate of interest or time duration is less than zero using log.Fatal function.
The order of initialisation of the is as follows,

The imported packages are initialized first. Hence simpleinterest package is initialized first and it's init method is called.
Package level variables p, r and t are initialized next.
init function is called in main package
main function is called at last.


====================

in main.go to

var p, r, t = -5000.0, 10.0, 1.0
We have initialised p to negative.

Now if you run the application, you will see

Simple interest package initialized
Main package initialized
2024/04/01 02:58:32 Principal is less than zero

p is negative. Hence when the init function runs, the program terminates after printing Principal is less than zero.

====================================================================

******************Use of blank identifier***********************

It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do so. The reason for this is to avoid bloating of unused packages which will significantly increase the compilation time. Replace the code in main.go with the following,

package main
  
import (
        "learnpackage/simpleinterest"
)

func main() {

}

The above program will error

# learnpackage
./main.go:4:2: imported and not used: "learnpackage/simpleinterest"
But it is quite common to import packages when the application is under active development and use them somewhere in the code later if not now. The _ blank identifier saves us in those situations.

The error in the above program can be silenced by the following code,

package main
  
import (
        "learnpackage/simpleinterest"
)

var _ = simpleinterest.Calculate

func main() {

}


The line var _ = simpleinterest.Calculate mutes the error. We should keep track of these kinds of error silencers and remove them including the imported package at the end of application development if the package is not used. Hence it is recommended to write error silencers in the package level just after the import statement.

Sometimes we need to import a package just to make sure the initialization takes place even though we do not need to use any function or variable from the package. For example, we might need to ensure that the init function of the simpleinterest package is called even though we plan not to use that package anywhere in our code. The _ blank identifier can be used in this case too as shown below.

package main

import (
	_ "learnpackage/simpleinterest"
)

func main() {

}

Running the above program will output Simple interest package initialized. We have successfully initialized the simpleinterest package even though it is not used anywhere in the code.