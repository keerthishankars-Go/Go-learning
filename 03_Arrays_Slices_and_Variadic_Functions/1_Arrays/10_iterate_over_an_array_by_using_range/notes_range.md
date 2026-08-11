I’ll explain for range at syntax level, execution flow, why it exists, and why \_ is used, so it really clicks.

🔹 Full program (reference)
package main

import "fmt"

func main() {
a := [...]float64{67.7, 89.8, 21, 78}
sum := float64(0)

    for i, v := range a {
    	fmt.Printf("%d th element of a is %.2f\n", i, v)
    	sum += v
    }

    fmt.Println("\nsum of all elements of a", sum)

}

1️⃣ Why range exists in Go

Before range, you would write:

for i := 0; i < len(a); i++ {
v := a[i]
}

Problems:

More code

Easy to make off-by-one errors

Less readable

📌 Go introduced range to make iteration safer and clearer.

2️⃣ Syntax of for range (VERY IMPORTANT)
for index, value := range collection {
}

General rule:

Collection range gives
array index, value
slice index, value
string index, rune
map key, value
channel value
3️⃣ Syntax-level breakdown of your loop
for i, v := range a {

What Go does internally

Each iteration:

i → index (0, 1, 2, …)

v → value at that index

Equivalent to:

for i := 0; i < len(a); i++ {
v := a[i]
}

But safer and cleaner.

4️⃣ Execution flow (step-by-step)
Array:
a := [...]float64{67.7, 89.8, 21, 78}

Iteration 1

i = 0

v = 67.7

Print

sum = 0 + 67.7

Iteration 2

i = 1

v = 89.8

sum = 67.7 + 89.8

Iteration 3

i = 2

v = 21

sum = 178.5 + 21

Iteration 4

i = 3

v = 78

sum = 199.5 + 78

Loop ends

Final sum = 256.5

5️⃣ Why sum := float64(0)
sum := float64(0)

Why not just 0?

Because:

v is float64

sum += v requires same type

📌 Go does not do implicit type conversion.

This avoids hidden bugs.

6️⃣ Why range returns BOTH index and value

Because sometimes you need:

only value

only index

both

Go gives both by default.

7️⃣ Ignoring values using _ (blank identifier)
Ignore index
for _, v := range a {
sum += v
}

\_ means: “I don’t care about this value”

Prevents unused variable errors

Ignore value
for i, \_ := range a {
fmt.Println(i)
}

Or idiomatically:

for i := range a {
fmt.Println(i)
}

📌 If you use only one variable, Go assumes it is the index.

8️⃣ Why \_ exists in Go

Go does not allow unused variables.

So instead of:

for i, v := range a {
// use only v
}

(which would error on i)

You explicitly say:

for \_, v := range a {
}

📌 This forces clean, intentional code.

9️⃣ Real-world usage (THIS is how Go code looks)
Sum values
for \_, price := range prices {
total += price
}

Process items
for \_, user := range users {
sendEmail(user)
}

Index-based logic
for i, item := range items {
fmt.Println(i, item)
}

🔟 Important range rules to remember

range works on arrays, slices, maps, strings, channels

It returns index + value

Order is guaranteed for arrays & slices

\_ is used to ignore unwanted values

Cleaner than traditional for

Preferred in real Go code

🧠 Mental model (lock this in)

“range walks the collection and hands you what you need.”

✅ Final takeaway

for range is idiomatic Go

Reduces bugs

Improves readability

Used everywhere in real projects

Next best follow-ups:
1️⃣ range with slices (memory behavior)
2️⃣ range with maps (order warning!)
3️⃣ range with strings (runes vs bytes)
4️⃣ Common range mistakes

============================================================================

This confusion is very common when learning Go because `range` looks like a special syntax. The trick is: **don't memorize `range`; understand what it gives you.**

Think of `range` as:

> "Go, take this collection and give me each item's position and value one by one."

---

Your code:


sum := float64(0)

for i, v := range a {
    fmt.Printf("%d th element of a is %.2f\n", i, v)
    sum += v
}


Assume:


a := [...]float64{1.89, 32.10, 12.43, 12.45}


Go internally does this:

### First iteration


i = 0
v = 1.89


Your code executes:


fmt.Printf("%d th element of a is %.2f\n", i, v)
sum += v


Now:


sum = 1.89


---

### Second iteration


i = 1
v = 32.10


Now:


sum = 1.89 + 32.10
     = 33.99


---

### Third iteration


i = 2
v = 12.43


Now:


sum = 46.42


---

### Fourth iteration


i = 3
v = 12.45


Now:


sum = 58.87


Loop ends.

---

## The mental model you should have

Whenever you see:


for i, v := range a


immediately translate it in your head:


for every element in a:
       give me index -> i
       give me value -> v


Like:


a = [10,20,30]

range gives:

iteration 1:
i = 0
v = 10

iteration 2:
i = 1
v = 20

iteration 3:
i = 2
v = 30


---

## Three forms of range you must remember

### 1. Need index + value

Most common:


for i, v := range a {
    fmt.Println(i, v)
}


Output:


0 10
1 20
2 30


---

### 2. Need only value

Ignore index using `_`


for _, v := range a {
    fmt.Println(v)
}


Output:


10
20
30


Here:


_ = index (throw it away)
v = value


---

### 3. Need only index


for i := range a {
    fmt.Println(i)
}


Output:


0
1
2


---

"Is it `for i,v := range`? Is range after variable? What comes first?"

Instead remember this sentence:

> **"I am looping over a collection, so I write `for` → variables I want → `range` → collection."**

Template:


for <what I want> := range <collection> {

}


Examples:


for index,value := range array



for _,value := range slice



for key,value := range map


---

For backend Go interviews, you will use this pattern constantly:


for _, user := range users {
    process(user)
}


So your muscle memory should become:

**collection after range always.**

`for i, v := range a` = "give me index and value from a."
