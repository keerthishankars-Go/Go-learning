package main

import "fmt"
✔ package main
Every runnable Go program must be in package main.

main package tells Go:

“This is an executable program. Build a binary.”

✔ import "fmt"
Imports the fmt package (formatting)

Lets us use fmt.Println, fmt.Print etc.

🟩 Start of the program
go
Copy code
func main() {
Execution starts here

Required for every Go executable

🟦 1️⃣ Declaring an empty array
go
Copy code
var a [5]int
fmt.Println("emp:", a)
✔ What this means:
Create an array of 5 integers

Since you did not give values, Go fills it with zero values

For integers, the zero value is 0

Array becomes:

csharp
Copy code
[0 0 0 0 0]
🟦 2️⃣ Setting and getting values
go
Copy code
a[4] = 100
fmt.Println("set:", a)
fmt.Println("get:", a[4])
✔ Explanation:
a[4] = 100 → sets 5th element (index 4)

a[4] → retrieve the element at index 4

Now array is:

csharp
Copy code
[0 0 0 0 100]
🟦 3️⃣ Getting array length
go
Copy code
fmt.Println("len:", len(a))
✔ Explanation:
len(a) gives the number of elements

Always 5 because arrays have fixed size

Output:

go
Copy code
len: 5
🟦 4️⃣ Using array literal values
go
Copy code
b := [5]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)
✔ Explanation:
Create array with exact length 5

Initial values provided explicitly

🟦 5️⃣ Letting Go infer array length using [...]
go
Copy code
b = [...]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)
✔ Explanation:
[...] means:

“Count elements and use that number as array length.”

So:

csharp
Copy code
[1 2 3 4 5]
Length = 5

🟦 6️⃣ Indexed initialization
go
Copy code
b = [...]int{100, 3: 400, 500}
fmt.Println("idx:", b)
This is the most confusing part — but easy if you follow step-by-step.

Step-by-step:
100 → goes to index 0
[100 _ _ _ _]

3:400 → puts 400 at index 3
[100 _ _ 400 _]

500 → next index after highest mentioned index (which is 3)
goes to index 4
[100 0 0 400 500]

Final:

csharp
Copy code
[100 0 0 400 500]
Go fills missing values with zero.

🟦 7️⃣ Creating a 2D array
go
Copy code
var twoD [2][3]int
for i := range 2 {
for j := range 3 {
twoD[i][j] = i + j
}
}
fmt.Println("2d: ", twoD)
✔ Explanation:
Step 1 — Create 2×3 matrix
[2][3]int means:

csharp
Copy code
[
[0 0 0]
[0 0 0]
]
Step 2 — Loops using Go 1.22 range <int>
range 2 → gives 0,1
range 3 → gives 0,1,2

So Go runs:

Copy code
twoD[0][0] = 0 + 0 = 0
twoD[0][1] = 0 + 1 = 1
twoD[0][2] = 0 + 2 = 2
twoD[1][0] = 1 + 0 = 1
twoD[1][1] = 1 + 1 = 2
twoD[1][2] = 1 + 2 = 3
Final:

csharp
Copy code
[
[0 1 2]
[1 2 3]
]
🟦 8️⃣ Declaring a 2D array with literal values
go
Copy code
twoD = [2][3]int{
{1, 2, 3},
{1, 2, 3},
}
fmt.Println("2d: ", twoD)
✔ Explanation:
This directly initializes the matrix:

csharp
Copy code
[
[1 2 3]
[1 2 3]
]
🟩 FULL CONCEPT SUMMARY
✔ Arrays in Go:
Fixed size

Size is part of the type

Automatically zero-filled

Stored in contiguous memory

Rarely used directly in real apps (slices are preferred)

✔ Literal initialization:
go
Copy code
[5]int{...}
[... ]int{...}
✔ Indexed initialization:
go
Copy code
3:400 // sets index 3
✔ Multidimensional arrays:
go
Copy code
[rows][cols]int
