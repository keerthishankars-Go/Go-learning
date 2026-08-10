1️⃣ The REAL problem (in simple words)

A slice keeps the entire array alive in memory — even if you only need a small part of it.

So:

Big array ❌ stays in memory

Small slice ✅ uses only a few elements

Garbage Collector ❌ cannot free the array

This is not a bug — it’s how slices are designed.

2️⃣ Why this happens (key idea)

A slice does not own data.

A slice only has:

pointer → array

length

capacity

As long as any slice points to an array, Go’s garbage collector says:

“Someone might still use this array → DON’T delete it”

3️⃣ Let’s walk your code step-by-step (VERY IMPORTANT)
Code
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy
}

4️⃣ Step-by-step MEMORY FLOW
🔵 Step 1: Create the original slice
countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}


Memory:

Array: [USA Singapore Germany India Australia]
Slice countries:
    ptr → array[0]
    len = 5
    cap = 5

🔵 Step 2: Create a sub-slice
neededCountries := countries[:len(countries)-2]


This becomes:

neededCountries = ["USA", "Singapore", "Germany"]


BUT in memory:

neededCountries:
    ptr → same array[0]
    len = 3
    cap = 5   ← IMPORTANT


⚠️ The entire array is STILL referenced

So memory is still:

[USA Singapore Germany India Australia]

🔴 If you returned neededCountries here (BAD CASE)
return neededCountries


Then:

Caller keeps a slice pointing to the original array

Garbage collector ❌ cannot free the array

Even "India", "Australia" stay in memory

👉 Memory leak risk when arrays are huge

5️⃣ Now the FIX: copy
🔵 Step 3: Create a NEW slice with make
countriesCpy := make([]string, len(neededCountries))


Memory:

New array (empty): [ "", "", "" ]
countriesCpy:
    ptr → new array
    len = 3
    cap = 3

🔵 Step 4: Copy data
copy(countriesCpy, neededCountries)


What happens:

Values are copied element-by-element

NO shared memory

Now:

Original array: [USA Singapore Germany India Australia]
New array:      [USA Singapore Germany]

🔵 Step 5: Return the new slice
return countriesCpy


At this point:

countries and neededCountries go out of scope

No references remain to the original array

✅ Garbage Collector can now FREE it

6️⃣ Final memory state after function returns
Returned slice:
    ["USA", "Singapore", "Germany"]

Original array:
    ❌ Garbage collected

7️⃣ Why copy solves the problem
Without copy	With copy
Slice points to big array	Slice points to new small array
Big array stays in memory	Big array can be GC’d
Risk of memory leak	Memory efficient
Faster initially	Slight copy cost
8️⃣ Rule to remember (VERY IMPORTANT)

If a slice must live long, and it comes from a large array → COPY IT

9️⃣ When should YOU use this pattern?

Use copy when:

You sliced a large buffer

You return a slice from a function

You store slice globally / long-lived

You want to avoid memory leaks

❌ Don’t bother copying for:

Small arrays

Short-lived slices

Tight loops where performance matters more than memory

🔟 One-line mental model

Slicing is cheap but keeps memory alive.
Copying costs a little but frees memory.

🧠 Interview-ready answer

“Slices retain references to the entire underlying array. To avoid unnecessary memory retention, we copy the required portion into a new slice so the original array can be garbage collected.”