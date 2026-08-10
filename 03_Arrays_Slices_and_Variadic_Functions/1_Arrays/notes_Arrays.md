What is an Array?
An array is a collection of elements. For example, the collection of integers 5, 8, 9, 79, 76 constitute an array.

Declaration
An array belongs to the type [n]T.

n represents the number of elements in an array and T represents the type of each element. The number of elements n is also a part of the type.

The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized. Don’t worry about this restriction since slices exist to overcome this.

package main

func main() {
a := [3]int{5, 78, 8}
var b [5]int
b = a //not possible since [3]int and [5]int are distinct types
}

In line no. 6 of the program above, we are trying to assign a variable of type [3]int to a variable of type [5]int which is not allowed and hence the compiler will print the following error

./prog.go:6:7: cannot use a (type [3]int) as type [5]int in assignment

========================================================================

1️⃣ Full program (reference)
package main

import "fmt"

func changeLocal(num [5]int) {
num[0] = 55
fmt.Println("inside function ", num)
}

func main() {
num := [...]int{5, 6, 7, 8, 8}
fmt.Println("before passing to function ", num)
changeLocal(num)
fmt.Println("after passing to function ", num)
}

2️⃣ Syntax-level breakdown (line by line)
🔹 Function signature
func changeLocal(num [5]int)

This means:

num is a parameter

Its type is array of exactly 5 integers

[5]int is a complete value type

📌 Rule (important)

In Go, [5]int and [6]int are different types.

🔹 Array declaration in main
num := [...]int{5, 6, 7, 8, 8}

... tells Go to infer the size

Go counts elements → size = 5

Actual type becomes [5]int

So both:

num := [...]int{...}

and

func changeLocal(num [5]int)

match exactly.

3️⃣ What happens when calling the function (core concept)
changeLocal(num)

What Go does internally

Takes array num from main

Copies all 5 elements

Creates a new array for changeLocal

Passes the copy

Memory picture:

main() num → [5 6 7 8 8]
changeLocal num → [5 6 7 8 8] (COPY)

📌 This is pass-by-value

4️⃣ Why modification doesn’t affect original
num[0] = 55

This modifies:

✅ the copied array

❌ NOT the original one

That’s why:

inside function [55 6 7 8 8]
after passing [5 6 7 8 8]

5️⃣ Why Go chose this design (WHY we take it like this)

Go treats arrays like:

int

float64

struct

All are value types.

Design goals:

✔ No hidden side effects
✔ Safer code
✔ Predictable behavior
✔ Easier reasoning

📌 You always know:

“If I pass an array, it will NOT be modified unless I explicitly allow it.”

6️⃣ What is the use of this behavior?
✅ Use case 1: Read-only safety
func checksum(data [32]byte) int {
// safe: data cannot be modified outside
}

Used in:

crypto

validation

hashing

✅ Use case 2: Immutable logic
func sortedCopy(arr [5]int) [5]int {
// modify copy and return
}

Caller’s data stays intact.

✅ Use case 3: Small fixed data

Arrays are useful when:

Size is small

Fixed at compile time

Must not be changed accidentally

Example:

var directions = [4]string{"N", "S", "E", "W"}

7️⃣ Why arrays are rare in real Go code

Because:

❌ Copying large arrays is expensive
❌ Size must match exactly
❌ Not flexible

That’s why Go introduced slices.

8️⃣ The real-world alternative (IMPORTANT)
Using slice (reference behavior)
func changeLocal(num []int) {
num[0] = 55
}

Call:

changeLocal(num[:])

Now:

No copy

Shared underlying array

Original changes

📌 Slices = shared, mutable
📌 Arrays = copied, immutable by default

9️⃣ Interview-level explanation (short)

“Arrays in Go are value types. When passed to functions, they are copied. This prevents unintended mutations. Slices are used instead when shared mutation is required.”

🔟 Rules to remember (save this)

Arrays are value types

Passing array copies all elements

[n]int is a distinct type

Modifications affect only the copy

Use arrays for fixed, safe data

Use slices for real applications

🧠 Mental model (lock this in)

Array = photocopy
Slice = shared view

=============================================================

I’ll show exact execution flow, what memory is created, what gets copied, what gets destroyed, and why output looks like this.

🧠 Program we are tracing
func changeLocal(num [5]int) {
num[0] = 55
fmt.Println("inside function ", num)
}

func main() {
num := [...]int{5, 6, 7, 8, 8}
fmt.Println("before passing to function ", num)
changeLocal(num)
fmt.Println("after passing to function ", num)
}

🔁 FULL EXECUTION FLOW (STEP BY STEP)
🟢 STEP 1: Program starts

Go runtime always starts here:

func main()

Nothing else runs before main().

🟢 STEP 2: Array creation in main
num := [...]int{5, 6, 7, 8, 8}

What Go does internally

Sees ... → counts elements → size = 5

Allocates memory for [5]int

Stores values

Memory now looks like:
main stack:
num → [5, 6, 7, 8, 8]

🟢 STEP 3: First print
fmt.Println("before passing to function ", num)

Output:

before passing to function [5 6 7 8 8]

No copies yet. Only one array exists.

🟡 STEP 4: Function call happens
changeLocal(num)

This is the most important moment.

🟡 STEP 5: ARRAY IS COPIED (CRITICAL)

Because function signature is:

func changeLocal(num [5]int)

Go must do this:

Internal steps

Allocate NEW memory for parameter num

Copy all 5 values from main.num

Pass the COPY into the function

Memory now:
main stack:
num → [5, 6, 7, 8, 8]

changeLocal stack:
num → [5, 6, 7, 8, 8] ← COPY

📌 These two arrays are independent.

🟡 STEP 6: Function body executes
num[0] = 55

What changes?

ONLY this array:

changeLocal.num → [55, 6, 7, 8, 8]

Main’s array is untouched.

🟡 STEP 7: Print inside function
fmt.Println("inside function ", num)

Output:

inside function [55 6 7 8 8]

Correct — because we printed the copied array.

🔴 STEP 8: Function returns
}

What happens now?

changeLocal finishes

Its stack frame is destroyed

Copied array is discarded

Memory becomes:

main stack:
num → [5, 6, 7, 8, 8]

📌 The modified copy no longer exists.

🟢 STEP 9: Back to main()

Execution resumes right after:

changeLocal(num)

🟢 STEP 10: Final print
fmt.Println("after passing to function ", num)

Since original array was never changed:

Output:

after passing to function [5 6 7 8 8]

🔁 COMPLETE FLOW SUMMARY (ONE GLANCE)
main() starts
↓
array created in main
↓
array printed
↓
array COPIED into function
↓
copy modified
↓
copy printed
↓
function exits → copy destroyed
↓
main resumes
↓
original array printed (unchanged)

🧠 WHY Go DOES THIS (IMPORTANT)

Go follows explicit behavior:

Type Passing behavior
int copied
struct copied
array copied
slice shared
map shared
pointer shared

📌 Arrays behave like numbers and structs.

🧪 THINK OF IT LIKE THIS
Array = 📄 Photocopy

You edit your copy

Original stays safe

Slice = 🔗 Shared link

Everyone sees changes

❗ WHY THIS MATTERS IN REAL PROJECTS
❌ Bug prevention
func validate(data [32]byte) { } // safe

❌ Performance concern

Large arrays = expensive copies

✅ Why slices are preferred
func process(data []byte) { } // shared, efficient

✅ FINAL RULES (MEMORIZE)

Arrays are copied when passed

Function works on its own copy

Original is unchanged

Copy dies when function exits

Slices exist to avoid this copying

=======================================

In Go, the array size is an integral part of its type.

This means that [5]int and [10]int are considered two entirely different and incompatible types by the compiler.
This design choice has several important implications:
Fixed Size: Once an array is declared, its size cannot be changed at runtime. Resizing an array would technically mean changing its type, which is not allowed.
Type Safety: The compiler can perform size checks at compile time, which helps catch potential errors like out-of-bounds access before the program runs.
Value Semantics: Arrays are value types. When an array is assigned to a new variable or passed to a function, a complete copy of the underlying data is created. This is different from many other languages (like C or Java) where arrays are reference types.
Memory Layout: The fixed size allows the compiler to know the exact memory footprint of the array at compile time, leading to efficient memory access and optimization of operations.
If you need a dynamic, variable-sized collection, Go provides slices, which are a more flexible data structure built on top of arrays. Slices are generally preferred in typical Go code for their versatility. The type of a slice, []int, does not include the length, allowing it to grow or shrink dynamically using functions like append.

=============================================

In Go, arrays are value types. 

This means that when an array is assigned to a new variable or passed to a function, a complete copy of the array is created, rather than a reference to the original one. 
This behavior distinguishes them from slices, which are reference types (or more accurately, the slice header, which contains a pointer to an underlying array, is copied, resulting in shared underlying memory). 
Key Implications of Arrays Being Value Types:
Copying: Assigning one array to another (e.g., b := a) creates a new, independent copy. Changes to b will not affect the original array a.
Function Parameters: When an array is passed as an argument to a function, the function receives a copy. Any modifications made to the array inside the function remain local to that function and do not impact the caller's original array.
Performance: For large arrays, this copying behavior can be memory-intensive and may impact performance. For most use cases requiring dynamic size or shared memory, Go slices are typically preferred. 
The fixed length is part of an array's type, which is why [5]int and [10]int are considered two different, incompatible types. 
