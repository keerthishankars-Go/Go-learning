✅ Full Explanation of the Functions Example
1️⃣ Function Definition: plus(a int, b int) int
func plus(a int, b int) int {
return a + b
}

What this means:
Part Meaning
func We are defining a function
plus Function name
(a int, b int) It takes two parameters, both integers
int after parentheses This function returns an integer
return a + b Explicit return value
✔ Why explicit return?

In some languages (Python, JavaScript), the last expression is returned automatically.
Go never does this.

You must write return.

💡 Where this is used in real-world Go?

Add totals in checkout

Add two amounts (tax + base price)

Combine two counts

Write small helper functions in bigger code

2️⃣ Shortened Parameter Syntax: plusPlus(a, b, c int)
func plusPlus(a, b, c int) int {
return a + b + c
}

Why we write:
a, b, c int

Instead of:

a int, b int, c int

👉 Because Go allows grouping parameters with the same type.

✔ Real-world use:

You will commonly see this when many values share the same type — like:

func createUser(name, email, password string) {}

3️⃣ Calling Functions in main()
res := plus(1, 2)
fmt.Println("1+2 =", res)

res = plusPlus(1, 2, 3)
fmt.Println("1+2+3 =", res)

What is happening?
Step 1 — Call plus(1,2)

Go passes 1 into a

Go passes 2 into b

Function returns 3

Stored in res

Step 2 — Call plusPlus(1,2,3)

Function receives 3 numbers

Returns sum → 6

🧠 Now let’s break down this concept deeper:
🔸 What is a function?

A reusable block of code that performs a specific job.

🔸 Why use functions?

Avoid repetition

Organize code

Make logic reusable

Easier debugging

Clean codebase

🛠 Practical real-world examples using this exact pattern
✔ Example 1 — Add two prices
func totalPrice(base int, tax int) int {
return base + tax
}

✔ Example 2 — Calculate total cart items
func totalItems(a, b, c int) int {
return a + b + c
}

✔ Example 3 — OTP validation helper
func validOTP(input string, real string) bool {
return input == real
}

✔ Example 4 — Check if user is allowed
func canAccess(role string) bool {
return role == "admin"
}

⭐ KEY TAKEAWAYS (Easy to Remember)
✔ 1. Functions start with func
✔ 2. Parameters need types
✔ 3. Return type is after parameters
✔ 4. You must use return
✔ 5. Group similar parameters:
func f(a, b, c int)

✔ 6. Call function like normal:
f(10, 20)

# ===============================

you MUST declare a function before you can use it.
You cannot write:

res := plusPlus(1, 2, 3)

unless somewhere in the same file/package you define:

func plusPlus(a, b, c int) int {
return a + b + c
}

Otherwise Go gives an error:

undefined: plusPlus

📝 Why Go requires you to declare the function?

Because Go is a compiled language, not interpreted.

When you compile:

Go checks all function names

All types must match

All function calls must refer to existing functions

If code calls a function that hasn’t been defined → error

This ensures:

No bugs

No missing code

No wrong types

No runtime surprises

🤔 But in JavaScript/Python we don’t declare first…?

Correct.
JS and Python interpret code line-by-line and allow implicit creation.

Go does not do that.

Go requires:

✔ Explicit declarations
✔ Explicit return types
✔ Exact parameter types

This makes Go safer for large backend systems.

🔍 But does the function need to be ABOVE main?

No.
You can define it before or after main().

This works:
func plus(a int, b int) int {
return a + b
}

func main() {
res := plus(1, 2)
fmt.Println(res)
}

This also works:
func main() {
res := plus(1, 2)
fmt.Println(res)
}

func plus(a int, b int) int {
return a + b
}

Go doesn’t care about order.
It cares only that the function exists somewhere in the same package.

❌ What you CANNOT do

You cannot call a function that is not defined anywhere:

func main() {
res := plusPlus(1, 2, 3) // ❌ ERROR
fmt.Println(res)
}

Error:

undefined: plusPlus

⭐ Why must we define it?

Because Go needs to know:

What parameters it takes (types)

What it returns (type)

How to generate machine code for it

Without a definition, Go cannot compile the program.
