One question might be bothering you though. If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic length. Well what happens under the hood is, when new elements are appended to the slice, a new array is created. The elements of the existing array are copied to this new array and a new slice reference for this new array is returned. The capacity of the new slice is now twice that of the old slice.

====================================================================================

I’ll explain append(), with what happens first, then next, and why Go does this

🔴 ONE-LINE IDEA (READ FIRST)

A slice can grow.
If there is no space left, Go creates a new array behind the scenes.

That’s it. Now let’s see how.

🔹 Your program (reference)
cars := []string{"Ferrari", "Honda", "Ford"}
fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))

cars = append(cars, "Toyota")

fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

1️⃣ Step 1: Slice creation (IMPORTANT)
cars := []string{"Ferrari", "Honda", "Ford"}

What Go creates

Behind the scenes:

An array of size 3

A slice pointing to it

Memory:

Underlying array:
[ Ferrari | Honda | Ford ]

cars slice:

- length = 3
- capacity = 3

📌 Slice is full — no extra space

2️⃣ Step 2: Print initial state
len(cars) = 3
cap(cars) = 3

✔ Length = how many elements exist
✔ Capacity = how many elements can fit

3️⃣ Step 3: This is the critical line
cars = append(cars, "Toyota")

Let’s break this into tiny steps.

4️⃣ What append() checks first

append() asks:

“Is there room in the existing array?”

Answer:

❌ NO

length = 3

capacity = 3

So Go cannot add more elements to the same array.

5️⃣ What Go does next (THIS IS THE MAGIC)

Since there is no space:

Go creates a new array

New array size = double the old capacity

old capacity = 3
new capacity = 6

Copies existing elements into new array

Adds "Toyota" to the new array

Returns a new slice pointing to the new array

📌 Old array still exists but is no longer used by cars

6️⃣ Memory after append
New underlying array:
[ Ferrari | Honda | Ford | Toyota | _ | _ ]

cars slice now:

- length = 4
- capacity = 6

The \_ slots are unused but reserved.

7️⃣ Why we MUST reassign (cars = append(...))
cars = append(cars, "Toyota")

Why this is required

Because:

append() returns a new slice

That slice may point to a different array

❌ This is WRONG:

append(cars, "Toyota") // result ignored

✔ This is CORRECT:

cars = append(cars, "Toyota")

8️⃣ Final print (now it makes sense)
len(cars) = 4
cap(cars) = 6

Output:

cars: [Ferrari Honda Ford Toyota] has new length 4 and capacity 6

🔥 WHY Go doubles capacity (DESIGN REASON)

If Go increased capacity by 1 every time:

Too many allocations

Very slow

Doubling:

Fewer reallocations

Faster append

Predictable performance

📌 This is a performance optimization

9️⃣ Important clarification (VERY IMPORTANT)

Append does NOT always create a new array

If capacity is available:

s := make([]int, 0, 10)
s = append(s, 1) // no new array

New array is created only when capacity is exceeded.

10️⃣ Variadic syntax explained (simple)
func append(s []T, x ...T) []T

...T means:

Accept one or many values

Examples:

append(s, 1)
append(s, 1, 2, 3)
append(s, otherSlice...)

🧠 Mental model (LOCK THIS IN)
Slice growth
Enough capacity → grow in same array
No capacity → new array + copy

✅ Final rules to remember

Slices grow using append

Arrays never grow

append may allocate new memory

Capacity usually doubles

Always reassign the result of append

Slice may stop sharing its old array

🧩 Why this matters in real apps

Efficient API response building

High-performance data processing

Safe memory management

Predictable behavior

🎯 ONE-LINE SUMMARY

append() grows a slice; if needed, Go silently creates a new array.
