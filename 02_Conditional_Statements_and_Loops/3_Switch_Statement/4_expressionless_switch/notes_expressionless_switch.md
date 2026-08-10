1️⃣ Normal switch (quick reminder)

Typical switch:

switch x {
case 1:
case 2:
}


Here:

x is evaluated once

Each case is compared with x

2️⃣ Expressionless switch — what changes?

In your code:

switch {
case hour >= 6 && hour < 12:
	...
}


👉 There is NO expression after switch

What Go does internally

Go treats this as:

switch true {
case hour >= 6 && hour < 12:
	...
}


📌 Key rule

If switch has no expression, Go assumes switch true.

3️⃣ Why switch true works

Each case must evaluate to a boolean expression.

Example:

case hour >= 12 && hour < 17:


This expression evaluates to:

true → case executes

false → Go checks next case

📌 This makes expressionless switch behave like:

a clean, structured if–else if–else chain

4️⃣ Your program (reference)
hour := 15

switch {
case hour >= 6 && hour < 12:
	fmt.Println("It's the morning shift.")
case hour >= 12 && hour < 17:
	fmt.Println("It's the afternoon shift.")
case hour >= 17 && hour < 21:
	fmt.Println("It's the evening shift.")
case (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6):
	fmt.Println("It's the night shift.")
default:
	fmt.Println("Invalid hour.")
}

5️⃣ Step-by-step execution (important)
hour = 15

Go evaluates cases top to bottom:

hour >= 6 && hour < 12

15 >= 6 → true

15 < 12 → false

Result → false ❌

hour >= 12 && hour < 17

15 >= 12 → true

15 < 17 → true

Result → true ✅

→ Go executes this block and stops checking further cases

📌 Switch stops at first matching case (no fallthrough by default)

6️⃣ Why this is better than if-else if
Equivalent if-else version
if hour >= 6 && hour < 12 {
	fmt.Println("It's the morning shift.")
} else if hour >= 12 && hour < 17 {
	fmt.Println("It's the afternoon shift.")
} else if hour >= 17 && hour < 21 {
	fmt.Println("It's the evening shift.")
} else if (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6) {
	fmt.Println("It's the night shift.")
} else {
	fmt.Println("Invalid hour.")
}

Why switch is preferred

✔ Cleaner
✔ Easier to read
✔ Less nesting
✔ Safer ordering
✔ Easy to add/remove cases

📌 Go style guideline

Use expressionless switch instead of long if-else chains.

7️⃣ Why parentheses in last case
case (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6):

Reason

&& has higher precedence than ||

Parentheses make logic explicit

Avoids subtle bugs

📌 Professional rule

Always group mixed && and || with parentheses.

8️⃣ Why default is needed
default:
	fmt.Println("Invalid hour.")


Catches unexpected values

Defensive programming

Ensures no silent failure

📌 Example:

hour := -1


Without default → nothing prints ❌

9️⃣ Real-world backend use cases (very common)
🔹 Request routing
switch {
case method == "GET":
case method == "POST":
}

🔹 Status classification
switch {
case code >= 200 && code < 300:
case code >= 400 && code < 500:
}

🔹 Permission checks
switch {
case user.IsAdmin:
case user.IsGuest:
}

🔟 Important rules to remember (save this)

switch expression is optional

Missing expression → switch true

Each case must return a boolean

First true case executes

No fallthrough by default

Cleaner alternative to if-else

🧠 Mental model (remember forever)

Expressionless switch = ordered decision table

✅ Final takeaway

Expressionless switch is idiomatic Go

Used heavily in production code

Best replacement for complex if-else

Improves readability and safety