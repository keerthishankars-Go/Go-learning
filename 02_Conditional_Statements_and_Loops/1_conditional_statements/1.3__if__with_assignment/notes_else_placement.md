Gotcha
The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.

Let’s understand this by means of a program.

package main

import (
"fmt"
)

func main() {
num := 10
if num % 2 == 0 { //checks if number is even
fmt.Println("the number is even")
}  
 else {
fmt.Println("the number is odd")
}
}
go
Run in playground

In the program above, the else statement does not start in the same line after the closing } of the if statement in line no. 11. Instead, it starts in the next line. This is not allowed in Go. If you run this program, the compiler will print the error,

./prog.go:12:5: syntax error: unexpected else, expected }
The reason is because of the way Go inserts semicolons automatically. You can read about the semicolon insertion rule here https://go.dev/ref/spec#Semicolons.

In the rules, it’s specified that a semicolon will be inserted after closing brace }, if that is the final token of the line. So a semicolon is automatically inserted after the if statement’s closing braces } in line no. 11 by the Go compiler.

So our program actually becomes

...
if num%2 == 0 {
fmt.Println("the number is even")
}; //semicolon inserted by Go Compiler
else {
fmt.Println("the number is odd")
}
after semicolon insertion. The compiler would have inserted a semicolon in line no. 4 of the above snippet.

Since if{...} else {...} is one single statement, a semicolon should not be present in the middle of it. Hence this program fails to compile. Therefore it is a syntactical requirement to place the else in the same line after the if statement’s closing brace }.
