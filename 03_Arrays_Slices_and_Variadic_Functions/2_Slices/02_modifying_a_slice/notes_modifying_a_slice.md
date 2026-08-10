🧠 BIG IDEA (one line)

A slice is NOT data — it is a VIEW into an array.

That sentence explains everything that follows.

🔹 Program we are tracing
package main

import "fmt"

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)

	for i := range dslice {
		dslice[i]++
	}

	fmt.Println("array after", darr)
}

1️⃣ Step 1: Array creation (real data is allocated)
darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}

What Go does

Allocates one array in memory

Stores 9 integers

Memory (conceptually):

Index:  0   1   2   3   4   5   6   7   8
darr = [57, 89, 90, 82,100, 78, 67, 69, 59]


📌 This array owns the data

2️⃣ Step 2: Slice creation (NO data copy happens)
dslice := darr[2:5]

Syntax meaning

Start index = 2

End index = 5 (exclusive)

Slice covers indices: 2, 3, 4

What Go actually creates

A slice header with 3 fields:

dslice:
- pointer → &darr[2]
- length  → 3
- capacity→ 7 (from index 2 to end of array)


📌 IMPORTANT

No new array is created
No data is copied

Visual memory layout
darr (real data):
[57, 89, 90, 82, 100, 78, 67, 69, 59]
          ↑    ↑     ↑
        dslice elements

3️⃣ Step 3: Print before modification
fmt.Println("array before", darr)


Output:

array before [57 89 90 82 100 78 67 69 59]


Nothing changed yet.

4️⃣ Step 4: Iterating over the slice
for i := range dslice {
	dslice[i]++
}

What range dslice gives

i → 0, 1, 2

These are slice indexes, NOT array indexes

Mapping:

dslice index	darr index
0	2
1	3
2	4
5️⃣ Step 5: Modification happens (CRITICAL)
dslice[i]++


This means:

darr[2]++
darr[3]++
darr[4]++


Because:

dslice points directly into darr

Same memory is shared

Memory after modification
Index:  0   1   2   3   4   5   6   7   8
darr = [57, 89, 91, 83,101, 78, 67, 69, 59]


📌 Array changed because slice modified it

6️⃣ Step 6: Print after modification
fmt.Println("array after", darr)


Output:

array after [57 89 91 83 101 78 67 69 59]


Matches exactly what we expect.

🔥 CORE RULE (MEMORIZE THIS)

Slice changes data → array changes
Array changes data → slice sees it

They are two views of the same memory.

7️⃣ Why Go designed slices like this (WHY)

Go needed:

Efficient data handling

No unnecessary copying

Fast function arguments

Flexible data structures

So Go chose:

Feature	Arrays	Slices
Own data	✅	❌
Fixed size	✅	❌
Passed by value	✅	header copied
Underlying data shared	❌	✅
Used in real apps	❌	✅
8️⃣ Why slices are passed to functions (real-world)
func process(data []int) {
	data[0] = 999
}


Only slice header copied (cheap)

Data shared

Changes visible everywhere

📌 This is why almost all Go APIs use slices.

9️⃣ Mental model (VERY IMPORTANT)
Array

🧱 Owns the building

Slice

🪟 Window into the building

If you change what you see through the window, the building changes.

🔟 Common beginner mistake ⚠️

❌ Thinking slice is a copy
❌ Thinking modifying slice is safe
❌ Forgetting slices share memory

✔ Always assume slices mutate shared data

✅ FINAL TAKEAWAY

Slice does NOT own data

Slice references array memory

No copy when slicing

Modifying slice modifies array

This is intentional and powerful

This is why slices dominate Go code

======================================================================

🔴 ONE SENTENCE FIRST (READ THIS)

An array stores the data.
A slice is just a pointer (handle) that looks at part of that array.

That’s it. Everything else comes from this.

1️⃣ Start with the code (only this part)
darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}

What REALLY happens

Go creates ONE array

This array owns the numbers

Think of it like this (real memory):

Index:  0   1   2   3   4   5   6   7   8
Value: [57, 89, 90, 82,100, 78, 67, 69, 59]


👉 This array is the only place where numbers live

2️⃣ Now this line (MOST IMPORTANT)
dslice := darr[2:5]

What you might THINK ❌

“Go created a new smaller array”

What ACTUALLY happens ✅

NO new array is created

Instead, Go creates a slice which only remembers:

Where to start (index 2)

How many items (3)

Which array it belongs to (darr)

So dslice is like:

dslice → points to darr[2]


Visual:

darr:
[57, 89, 90, 82, 100, 78, 67, 69, 59]
          ↑    ↑     ↑
        dslice sees these


👉 dslice does NOT own data
👉 dslice does NOT copy data

3️⃣ This print (nothing special yet)
fmt.Println("array before", darr)


Output:

[57 89 90 82 100 78 67 69 59]


Everything is unchanged.

4️⃣ Now the loop (THIS is where confusion happens)
for i := range dslice {
	dslice[i]++
}


Let’s break it into tiny pieces.

5️⃣ What is range dslice doing?

dslice has 3 elements, so:

i = 0
i = 1
i = 2


⚠️ IMPORTANT:

i is slice index

NOT array index

6️⃣ What does dslice[i]++ REALLY mean?

This is the key.

When i = 0
dslice[0]++


But remember:

dslice[0] → darr[2]


So Go actually does:

darr[2]++

When i = 1
dslice[1]++


This is actually:

darr[3]++

When i = 2
dslice[2]++


This is actually:

darr[4]++

7️⃣ After loop finishes (REAL memory)

Original array is now:

[57, 89, 91, 83, 101, 78, 67, 69, 59]


Because:

Slice modified the SAME memory

No copy exists

8️⃣ Final print (this now makes sense)
fmt.Println("array after", darr)


Output:

[57 89 91 83 101 78 67 69 59]

🔥 WHY THIS HAPPENS (VERY SIMPLE REASON)
Arrays:

Own data

Copy on function call

Heavy

Slices:

Just point to arrays

Light

Fast

Shared

Go wants performance, so slices are designed to share memory.

🧠 THE ONLY MENTAL MODEL YOU NEED
Array = 📦 Box with items
Slice = 🪟 Window into the box

If you change something through the window →
the box changes

❗ COMMON CONFUSION (YOU ARE NOT ALONE)

❌ “Slice is a smaller array”
❌ “Slice has its own memory”

✅ Slice is just reference + length

9️⃣ One last ultra-simple example
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]
s[0] = 100
fmt.Println(arr)


Output:

[1 100 3 4 5]


Because:

s[0] → arr[1]

✅ FINAL RULE (REMEMBER THIS)

If you modify a slice, assume the array is modified.

Always.