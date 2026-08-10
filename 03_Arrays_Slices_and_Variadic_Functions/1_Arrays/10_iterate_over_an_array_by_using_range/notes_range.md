I’ll explain for range at syntax level, execution flow, why it exists, and why _ is used, so it really clicks.

🔹 Full program (reference)
package main

import "fmt"

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)

	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
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

Collection	range gives
array	index, value
slice	index, value
string	index, rune
map	key, value
channel	value
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


_ means: “I don’t care about this value”

Prevents unused variable errors

Ignore value
for i, _ := range a {
	fmt.Println(i)
}


Or idiomatically:

for i := range a {
	fmt.Println(i)
}


📌 If you use only one variable, Go assumes it is the index.

8️⃣ Why _ exists in Go

Go does not allow unused variables.

So instead of:

for i, v := range a {
	// use only v
}


(which would error on i)

You explicitly say:

for _, v := range a {
}


📌 This forces clean, intentional code.

9️⃣ Real-world usage (THIS is how Go code looks)
Sum values
for _, price := range prices {
	total += price
}

Process items
for _, user := range users {
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

_ is used to ignore unwanted values

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