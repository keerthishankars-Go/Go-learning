package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}

======================================================

🧠 BIG PICTURE FIRST (WHY THIS PROGRAM EXISTS)

This program demonstrates 3 core Go concepts together:

2D arrays (fixed size data)

Passing arrays to functions (by value)

Iterating nested data using range

Everything in the code supports one of these goals.

1️⃣ Why package main (VERY IMPORTANT)
package main

Why this is required

Go programs start execution from main()

main() must belong to package main

This tells Go:

“This file builds an executable program”

📌 If you change this to:

package utils


→ program will not run (no entry point)

2️⃣ Why import "fmt"
import "fmt"

Purpose

Needed for output (Printf, Println)

Go does not auto-import anything

Explicit imports = clarity

📌 If fmt is unused → compiler error
(Go enforces clean code)

3️⃣ Program execution always starts here
func main() {


Nothing above main() executes automatically.
Functions above are definitions only.

4️⃣ First thing main() does: create array a
a := [3][2]string{
	{"lion", "tiger"},
	{"cat", "dog"},
	{"pigeon", "peacock"},
}

Why [3][2]string

3 rows

2 columns

Fixed-size table

Memory created in main:

a →
[
 ["lion", "tiger"],
 ["cat", "dog"],
 ["pigeon", "peacock"]
]

Why trailing comma is REQUIRED

Go auto-inserts semicolons

Multi-line literals must end with commas

Prevents parsing ambiguity

📌 This is a language rule, not a style choice.

5️⃣ Call to printarray(a) — CRITICAL POINT
printarray(a)

What happens EXACTLY

a is a [3][2]string array

Arrays are value types

Go copies the entire array

Copy is passed to printarray

Memory now:

main stack:
a → original array

printarray stack:
a → COPY of original


📌 This ensures:

printarray cannot modify main’s data

Safe, predictable behavior

6️⃣ Entering printarray
func printarray(a [3][2]string) {

Why exact type is required

[3][2]string is a complete type

[4][2]string would NOT compile

Size is part of the type

7️⃣ Outer range loop (row-level)
for _, v1 := range a {

What range gives here

_ → index (ignored)

v1 → one row → type [2]string

Flow:

Iteration 1 → v1 = ["lion","tiger"]
Iteration 2 → v1 = ["cat","dog"]
Iteration 3 → v1 = ["pigeon","peacock"]


📌 _ is used to avoid unused-variable error.

8️⃣ Inner range loop (column-level)
for _, v2 := range v1 {


v1 is [2]string

v2 is string

Flow for one row:

"lion" → print
"tiger" → print

9️⃣ Why fmt.Printf("\n") is OUTSIDE inner loop
fmt.Printf("\n")


Purpose:

New line after each row

Keeps output formatted as a table

📌 Inner loop prints cells
📌 Outer loop controls rows

🔁 First printarray call finishes

Copied array is destroyed

Control returns to main()

Original array a still intact

10️⃣ Creating second array b (different style)
var b [3][2]string

Why this works

Arrays get zero values

string zero value = ""

Initial memory:

b →
[
 ["", ""],
 ["", ""],
 ["", ""]
]

11️⃣ Manual assignment (cell-by-cell)
b[0][0] = "apple"
b[0][1] = "samsung"
...

Why this approach exists

Used when:

Data comes dynamically

Values arrive later

You don’t know all values upfront

12️⃣ Blank line output
fmt.Printf("\n")


Purely for readability
Separates outputs of two arrays

13️⃣ Second printarray(b) call

Same flow again:

b is copied

Passed to function

Printed using nested range

Copy destroyed

🧠 COMPLETE EXECUTION FLOW (ONE LOOK)
Program starts
↓
main()
↓
create array a
↓
call printarray(a)
    ↓ copy array
    ↓ iterate rows
    ↓ iterate columns
    ↓ print
    ↓ return (copy destroyed)
↓
create array b (empty)
↓
fill array b
↓
call printarray(b)
    ↓ copy array
    ↓ iterate & print
    ↓ return
↓
program ends

❓ WHY THIS DESIGN (IMPORTANT)
Choice	Why Go does it
Arrays are copied	Prevent accidental mutation
Fixed sizes	Compile-time safety
range loops	Safer iteration
Trailing commas	Clean parsing
package main	Clear entry point
🧠 Mental Model (LOCK THIS IN)

Arrays are fixed, safe, copied tables
Functions get their own copy
range walks the structure cleanly