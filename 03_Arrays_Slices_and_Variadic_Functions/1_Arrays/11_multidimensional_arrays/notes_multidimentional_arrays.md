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

==============================================================================

This is a very good question because **nested `range` loops are where many Go beginners get confused**.

Let's break this line by line.

Your array:


a := [3][2]string{
    {"lion", "cat"},
    {"tiger", "rabbit"},
    {"pigeon", "peacock"},
}


Think of it as a table:


Index        Value

a[0]  --->  {"lion", "cat"}
a[1]  --->  {"tiger", "rabbit"}
a[2]  --->  {"pigeon", "peacock"}


Actually, this is an **array of arrays**.

The type:


[3][2]string


means:


3 rows
Each row contains 2 strings


---

Now look at your function:


func printarray(a [3][2]string) {

    for _, v1 := range a {

        for _, v2 := range v1 {

            fmt.Printf("%s ", v2)

        }

        fmt.Printf("\n")
    }
}


---

## First range loop


for _, v1 := range a


Here:


a = [3][2]string


The range gives:


index, value


But you ignore index using `_`.

So:


_ = index
v1 = each inner array


First iteration:


v1 = {"lion", "cat"}


Second iteration:


v1 = {"tiger", "rabbit"}


Third iteration:


v1 = {"pigeon", "peacock"}


So after the first loop:


v1 is [2]string


Not a string.

It is another array.

---

## Second range loop

Now:


for _, v2 := range v1


Remember:


v1 = {"lion", "cat"}


So now range goes inside that array.

Again range gives:


index, value


You ignore index:


_ = index
v2 = string value


First iteration:


v2 = "lion"


Second iteration:


v2 = "cat"


Then:


fmt.Printf("%s ", v2)


prints:


lion cat


---

Then first loop moves to next row:


v1 = {"tiger", "rabbit"}


Second loop:


v2 = "tiger"
v2 = "rabbit"


Output:


tiger rabbit


---

Third row:


v1 = {"pigeon", "peacock"}

v2 = "pigeon"
v2 = "peacock"


Output:


pigeon peacock


---

## Visual execution


a
|
|-- v1 = {"lion","cat"}
|       |
|       |-- v2 = lion
|       |-- v2 = cat
|
|-- v1 = {"tiger","rabbit"}
|       |
|       |-- v2 = tiger
|       |-- v2 = rabbit
|
|-- v1 = {"pigeon","peacock"}
        |
        |-- v2 = pigeon
        |-- v2 = peacock


---

## The easiest way to remember nested range:

For a 2D array:


for _, row := range matrix {
    for _, column := range row {
        
    }
}


Think:


outer loop = rows
inner loop = elements inside each row


Example:


matrix := [2][3]int{
    {1,2,3},
    {4,5,6},
}


Outer range gives:


row = {1,2,3}
row = {4,5,6}


Inner range gives:


1 2 3
4 5 6


---

For backend Go work, you will see this pattern mostly with:

* JSON arrays
* database query results
* slices of structs
* nested API responses

So the mental shortcut:

> `range` over a 2D structure → first range gives the inner structure, second range gives the actual values.

=============================================================================

If you are explaining this code in an interview or while learning, describe it like this:

---

**Description:**

> This program demonstrates how to work with a two-dimensional array in Go. A two-dimensional array is an array where each element is itself another array. Here, the array `a` is of type `[3][2]string`, which means it contains 3 rows and each row contains 2 string elements.

> The `printarray` function receives this 2D array as a parameter. Inside this function, two nested `for range` loops are used to iterate through the array.

> The outer `for range` loop iterates over the main array. Since each element of the outer array is a `[2]string` array, the variable `v1` represents one complete inner array (one row).

> The inner `for range` loop iterates over `v1`. Here, the variable `v2` represents each individual string element inside that row.

> The underscore (`_`) is used to ignore the index because only the values are required for printing.

> Therefore, the outer loop moves row by row, and the inner loop moves element by element within each row, allowing all values of the two-dimensional array to be printed.

---

For your exact code, you can explain the flow:


a = [
      ["lion", "cat"],
      ["tiger", "rabbit"],
      ["pigeon", "peacock"]
    ]


Outer loop:


First iteration:
v1 = ["lion", "cat"]

Second iteration:
v1 = ["tiger", "rabbit"]

Third iteration:
v1 = ["pigeon", "peacock"]


Inner loop:


For v1 = ["lion", "cat"]

v2 = "lion"
v2 = "cat"


Then it moves to the next row.

---

A shorter interview-style explanation:

> "The outer range iterates through each inner array of the two-dimensional array, while the inner range iterates through each element inside that inner array. Since I only need values and not indexes, I use `_` to discard the index returned by range."

---

One important thing to remember:

In:


for _, v1 := range a


`v1` is **not a string**.

It is:


[2]string


because each item in `a` is another array.

Then:


for _, v2 := range v1


`v2` becomes:


string


because now you are iterating over the inner array elements.

This "type changes at every range level" is the key concept.

==============================================================================
******************************************************************************

This is an excellent question because you are touching the **core idea of nested range**. The confusion comes from thinking **v1 and v2 contain the same values**. They actually **do not**.

Let's look carefully.

Your array:


a := [3][2]string{
    {"lion", "cat"},
    {"tiger", "rabbit"},
    {"pigeon", "peacock"},
}


Memory representation:


a
|
|---- index 0 ---> [ "lion", "cat" ]
|
|---- index 1 ---> [ "tiger", "rabbit" ]
|
|---- index 2 ---> [ "pigeon", "peacock" ]


---

Now your outer loop:


for _, v1 := range a {


What does `v1` receive?

It receives **one complete inner array**.

First iteration:


v1 = ["lion", "cat"]


Second iteration:


v1 = ["tiger", "rabbit"]


Third iteration:


v1 = ["pigeon", "peacock"]


So `v1` is a **row**.

Its type is:


[2]string


---

Now inner loop:


for _, v2 := range v1 {


Here we are looping over the row.

Example:


v1 = ["lion", "cat"]


The range gives:

First iteration:


v2 = "lion"


Second iteration:


v2 = "cat"


So `v2` is an **individual element**.

Its type is:


string


---

## Why don't we print v1?

Because if you do:


fmt.Println(v1)


inside the outer loop:


for _, v1 := range a {
    fmt.Println(v1)
}


Output:


[lion cat]
[tiger rabbit]
[pigeon peacock]


You get each row as a group.

But your goal is:


lion cat
tiger rabbit
pigeon peacock


with each value formatted separately.

So you go one level deeper:


for _, v2 := range v1 {
    fmt.Printf("%s ", v2)
}


---

## Why do we need v1 then?

Because `v1` is the bridge between the two dimensions.

Without `v1`, you cannot reach the individual strings.

Think of it like a cupboard:


Kitchen cupboard (a)

    Shelf 1 (v1)
       |
       |-- Plate (v2)
       |-- Cup   (v2)

    Shelf 2 (v1)
       |
       |-- Plate (v2)
       |-- Cup   (v2)


You cannot directly pick the cup without opening the shelf first.

---

## Another way to write the same thing

You could use indexes:


func printArray(a [3][2]string) {

    for i := 0; i < len(a); i++ {

        for j := 0; j < len(a[i]); j++ {

            fmt.Printf("%s ", a[i][j])

        }

        fmt.Println()
    }
}


Here:


a[i]     == v1
a[i][j]  == v2


So:


v1 = a[i]      → complete row
v2 = a[i][j]   → individual value


---

The key sentence to remember:

> In a 2D array, the first range gives the inner array (row), and the second range gives the actual elements inside that row.

`v1` and `v2` are not duplicates.
`v1` = container
`v2` = item inside the container.
