🔴 ONE LINE FIRST (READ THIS)

Length = how many elements the slice currently has
Capacity = how many elements it can grow into before needing a new array

Keep this in mind while reading.

🔹 Your code (reference)
fruitarray := [...]string{
	"apple", "orange", "grape", "mango",
	"water melon", "pine apple", "chikoo",
}

fruitslice := fruitarray[1:3]

fmt.Printf("length of slice %d capacity %d",
	len(fruitslice), cap(fruitslice))

1️⃣ Step 1: Understand the array first (REAL DATA)
fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}


Array index map:

Index:  0        1        2        3        4            5             6
Value: apple | orange | grape | mango | water melon | pine apple | chikoo


Array length = 7

📌 This array owns all the data

2️⃣ Step 2: Slice creation (MOST IMPORTANT LINE)
fruitslice := fruitarray[1:3]

What [1:3] means

Start at index 1

Stop before index 3

So slice contains:

Index 1 → orange
Index 2 → grape

3️⃣ Slice contents (what it sees)
fruitslice = ["orange", "grape"]


So:

len(fruitslice) == 2


✔ Makes sense — 2 elements

4️⃣ Now CAPACITY (this is the confusing part)
Rule (IMPORTANT)

Capacity = elements available from slice start index till end of array

Slice starts at index 1

Array length = 7

So capacity = 7 - 1 = 6

Elements it could use:

orange, grape, mango, water melon, pine apple, chikoo


That’s 6 elements

5️⃣ Visual diagram (THIS MAKES IT CLICK)
fruitarray:
[ apple | orange | grape | mango | water melon | pine apple | chikoo ]
            ↑
        fruitslice starts here

fruitslice length = 2
fruitslice capacity = 6

6️⃣ Why slice capacity is NOT 2

Because:

Slice can grow

It is allowed to extend forward in the same array

As long as capacity allows

Example:

fruitslice = append(fruitslice, "mango")


This works without new allocation because capacity exists.

7️⃣ Why capacity exists at all (WHY Go designed this)

Without capacity:

Go wouldn’t know when to reallocate

append() would be inefficient

Capacity allows:

Fast appends

Fewer memory allocations

Predictable performance

📌 This is a performance feature, not a complexity.

8️⃣ What happens when capacity is exceeded

When you append beyond capacity:

Go creates a new array

Copies data

Slice now points elsewhere

This is how Go breaks sharing safely

9️⃣ Real-world meaning (VERY PRACTICAL)
Example: API response building
users := make([]User, 0, 100)


Meaning:

length = 0 (empty list)

capacity = 100 (can append 100 users efficiently)

This avoids repeated memory allocations.

🔟 Summary table (SAVE THIS)
Term	Meaning
len(slice)	How many elements it has now
cap(slice)	How many elements it can grow into
Capacity starts from	slice start index
Capacity tied to	underlying array
Used mainly by	append()
🧠 Mental model (LOCK THIS IN)

Length = what you see
Capacity = how much more you can add

✅ Final takeaway

Slices are views into arrays

Length is current window size

Capacity is how far the window can slide

Capacity enables fast growth

This is why slices replace arrays in real apps