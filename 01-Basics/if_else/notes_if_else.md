Here is a clear and simple explanation of your Go code:

if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}

✅ What this syntax means

This is a special Go feature:

🔹 You can declare a variable inside the if statement
🔹 That variable exists ONLY inside the entire if–else-if–else block
🔍 Step-by-step explanation
1. num := 9

This creates a new variable num only for this if-else chain.

It is equivalent to:

num := 9
if num < 0 {
    ...
}


But more compact.

🔍 2. First condition:
if num < 0


Is 9 less than 0? ❌ No
So this block is skipped.

🔍 3. Second condition:
else if num < 10


Is 9 less than 10? ✔ Yes
So it prints:

9 has 1 digit


And the entire if-else chain stops.

🔍 4. The final else only runs if ALL conditions above fail

In this case it does NOT run, because the else if was true.

🧠 Why use this pattern?

Using:

if x := <value>; condition {
    ...
}


is common in Go for:

Short variable lifetime

Cleaner code

Avoiding variables leaking into outer scope

Example use-case in real apps:

if user, err := db.GetUser(id); err == nil {
    fmt.Println("User found:", user.Name)
} else {
    fmt.Println("Error:", err)
}

⭐ Summary
Part	Meaning
num := 9	declare variable for this if-else chain
num < 0	first condition
num < 10	second condition
else	fallback

Output:

9 has 1 digit