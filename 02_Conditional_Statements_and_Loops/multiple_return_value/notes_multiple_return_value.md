Go allows a function to return more than one value.

This is one of Go’s strongest features, widely used in real backend code (especially for returning result, error).

Let’s break it down.

🟢 Example Code
package main

import "fmt"

func vals() (int, int) {
return 3, 7
}

func main() {
a, b := vals()
fmt.Println(a)
fmt.Println(b)

    _, c := vals()
    fmt.Println(c)

}

🔍 What’s happening?
1️⃣ Function returning two values
func vals() (int, int) {
return 3, 7
}

This means:

The function returns two integers

The return statement return 3, 7 gives both values

2️⃣ Receiving both returned values
a, b := vals()

This means:

a gets the first returned value → 3

b gets the second returned value → 7

Then:

fmt.Println(a) // 3
fmt.Println(b) // 7

3️⃣ Ignoring a value using \_
\_, c := vals()

Here:

\_ means: ignore this value

c receives only the second returned value → 7

Output:

7

🧠 Why Go uses multiple return values? (Real Backend Use)
⭐ 1. Errors

Most Go backend functions return:

value, err := someFunction()

Example:

user, err := repo.GetUser(id)

Why?

value contains valid data

err tells if something went wrong

⭐ 2. Query results
count, total := calculateStats()

⭐ 3. HTTP handlers
status, message := validatePayload(data)

⭐ 4. Database operations
rows, err := db.Query(query)

🟦 Why use \_ blank identifier?

Because sometimes you don't need all returned values.

Example:
You only want the second value:

\_, c := vals()

Example in real world:

\_, exists := cache[userID]

🟢 Easy Summary
Code Meaning
(int, int) function returns 2 integers
a, b := vals() capture both return values
\_, c := vals() ignore one return value
return 3, 7 return two values
🎯 In Simple Words

Think of Go’s multiple returns like handing someone two items at once.

Example:

Take these: (3, 7)

You can accept both:

a = 3
b = 7

Or ignore the first:

c = 7

# ==========================

✅ 1. Function definition syntax
func vals() (int, int) {
return 3, 7
}

Meaning:

func → keyword to define a function

vals() → function name vals with no parameters

(int, int) → the function returns two values, both of type int

return 3, 7 → return value1 = 3, value2 = 7

📌 Go allows multiple return values, so you must list them in parentheses.

✅ 2. Capturing return values
a, b := vals()

Meaning:

vals() returns two values, so Go requires you to receive them using two variables

a receives the first return value

b receives the second return value

:= means "create variables and assign values"

📌 Number of variables on the left must match the number of values returned.

❓ Why not just do this?
a := vals()

Because the function returns two values, not one.
So Go will give an error:

multiple-value vals() in single-value context

✅ 3. Ignoring a value with \_
\_, c := vals()

Meaning:

\_ → blank identifier (throw-away variable)

c will store only the second returned value

📌 Go requires the number of receiving variables to match the number of returned values.
So if you want to ignore one, you must put \_.

🧠 Why does Go force this syntax?

Because Go designers wanted:

✔️ safer code

No accidental ignoring of values.

✔️ predictable error handling

Most Go functions return (value, error), like:

user, err := getUser(id)

✔️ clear visibility

If you ignore something, Go makes you explicitly say:

_ = something

So you KNOW you're ignoring it.

🔎 Example to help you understand
Function returns two things:
return height, width

Case 1: You need both
h, w := getSize()

Case 2: You need only width
_, w := getSize()

Case 3: You need only height
h, _ := getSize()

That’s why the syntax uses:

multiple variables

commas

blank identifier _

:= for declaring multiple variables at once

🎯 Final Summary (Super Simple)
Syntax Meaning
(int, int) function returns 2 ints
a, b := vals() capture 2 return values
_, c := vals() ignore first, keep second
_ discard this value
:= declare variables and assign values
