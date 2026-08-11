When writing code live (especially in interviews), don't start by typing everything at once. Follow the **natural thought flow** of the program.

For this example, write in this order:

### 1. Start with package and imports

package main

import "fmt"

Reason: Every Go program needs these first. It sets the environment.

---

### 2. Write the function signature (the operation you want to demonstrate)

Before `main`, define the helper function:

func changeLocal(num [5]int) {
num[0] = 55
fmt.Println("inside function ", num)
}

Why first?

Because when someone reads your code, they immediately understand:

> "Okay, there is a function that receives an array and modifies its first element."

---

### 3. Write `main()` and create the original data

func main() {
num := [...]int{5, 6, 7, 8, 8}

}

Now the setup is clear.

---

### 4. Add the first observation (before function call)

fmt.Println("before passing to function ", num)

This establishes the initial state.

---

### 5. Call the function

changeLocal(num)

While typing, explain:

> "Arrays in Go are value types. Passing an array to a function creates a copy, so modifications inside the function won't affect the original array."

---

### 6. Print after the function call

fmt.Println("after passing to function ", num)

Now the proof is complete.

---

### Final live coding order:

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

---

For interviews, your mental sequence should be:

**Problem statement → Data structure → Function → Main setup → Function call → Verify result**

For Go specifically, avoid writing comments first. Explain concepts verbally while coding. Your code should show the behavior; your explanation should show the understanding.

=============================================================================

Yes, this is exactly the kind of question an interviewer **might ask**, because it tests whether you understand **execution flow** and not just "arrays are passed by value".

Your output:

Before going to pass an array into a function [1 3 4 5 5]

inside function [55 3 4 5 5]

After passing into func [1 3 4 5 5]

The reason **"inside function" appears in the middle** is because of the **order of execution**.

Let's trace your program:

### Step 1: `main()` starts

num := [...]int{1,3,4,5,5}

Memory:

main's num

[1 3 4 5 5]

---

### Step 2: First print executes

fmt.Println("Before going to pass an array into a function ", num)

Output:

Before going to pass an array into a function [1 3 4 5 5]

---

### Step 3: Function call happens

changeLocal(num)

Now Go creates a **copy** of the array.

Conceptually:

main num
+-------------+
|1 3 4 5 5 |
+-------------+

        copy

changeLocal num
+-------------+
|1 3 4 5 5 |
+-------------+

---

### Step 4: Inside function modifies its copy

num[0] = 55

Now:

main num
[1 3 4 5 5]

changeLocal num
[55 3 4 5 5]

Then:

fmt.Println("inside function ", num)

prints:

inside function [55 3 4 5 5]

This happens **before the function returns**.

---

### Step 5: Function finishes

Control goes back to `main()`.

Next line:

fmt.Println("After passing into func", num)

prints the original array:

After passing into func [1 3 4 5 5]

because the original was never modified.

---

### Interview answer:

If they ask:

> "Why is inside function printed in the middle?"

Say:

> "Because function execution is synchronous. When main calls changeLocal(), the control flow moves inside that function. The function modifies its local copy and prints it. Only after the function completes does execution return to main and continue with the next print statement."

This answer shows you understand **call stack execution**, not just pass-by-value.

A good follow-up they may ask:

> "If arrays are copied, what happens if the array is very large?"

That leads to why Go developers often use **slices**, because slices pass a small descriptor pointing to the underlying array instead of copying the entire array.

=============================================================================

The core difference is that a parameter is the variable defined in the function's declaration, while an argument is the actual value you pass to that function when you call it. [1, 2] 
## Quick Breakdown

* Parameter: The blueprint or placeholder (defined inside the function signature).
* Argument: The real data or value (passed into the function during execution). [3, 4, 5, 6] 

## Code Example

package main
// 'width' and 'height' are PARAMETERSfunc calculateArea(width int, height int) int {
    return width * height
}
func main() {
    // 10 and 5 are ARGUMENTS
    calculateArea(10, 5) 
}

## Key Differences

| Feature        |               Parameter |                Argument |
|---    |--- |---|
| When it exists | At declaration time.     | At runtime (execution). |
| What it is     | A variable name and type.| A concrete value or expression. |
| Location       | In the function signature.| In the function call. |
| Memory         | Defines the memory needed.| Holds the actual data in memory. |


