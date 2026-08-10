********\*\*\*\*********If … else if … else statement**********\*\*\***********
The if statement also has optional else if and else components.

The syntax for the same is provided below

if condition1 {
...
} else if condition2 {
...
} else {
...
}

The condition is evaluated for the truth from the top to bottom.

In the above statement, if condition1 is true, then the block of code within if condition1 { and the closing brace } is executed.

If condition1 is false and condition2 is true, then the block of code within else if condition2 { and the next closing brace } is executed.

If both condition1 and condition2 are false, then the block of code in the else statement between else { and } is executed.

There can be any number of else if statements.

In general, whichever if or else if’s condition evaluates to true, it’s corresponding code block is executed. If none of the conditions are true then else block is executed.

============================================

Bus ticket booking program:

Why we start with price := 0
price := 0

Syntax + reasoning

:= → short-hand declaration

price → variable name

0 → initial value

📌 Why 0?

0 is the zero value for numbers

Ticket can be free, so 0 is a valid default

If age < 5, we don’t even need to change it

===========================================

if age < 5
if age < 5 {
price = 0
}

Syntax explanation

if → decision keyword

age < 5 → boolean expression

must evaluate to true or false

{} → block executed only if true

📌 Even though price is already 0, we write this for:

readability

clarity

future modifications

=============================================

else if age >= 5 && age <= 22
else if age >= 5 && age <= 22 {
price = 10
}

Why else if?

We only want this check if the previous condition failed

Prevents unnecessary checks

Guarantees only one price is assigned

Why &&?
age >= 5 && age <= 22

&& → logical AND

Means both conditions must be true

📌 This ensures age is inside a range, not just one side.

===========================================

Why else (not another else if)?
else {
	price = 15
}

Reasoning

At this point:

age is not < 5

age is not between 5 and 22

👉 Only remaining possibility:

age > 22


So:

else = default case

Cleaner

Less error-prone

=================================================

Why not write everything inside if without price?

You could write:

if age < 5 {
	fmt.Println(0)
} else if age <= 22 {
	fmt.Println(10)
} else {
	fmt.Println(15)
}

But this is ❌ NOT good practice

Because:

Business logic and output are mixed

Cannot reuse price

Harder to test

Harder to extend

💡 Professional rule:

Calculate first, print later

==============================================

Why else if instead of multiple if?

❌ Wrong approach:

if age < 5 {
	price = 0
}
if age <= 22 {
	price = 10
}
if age > 22 {
	price = 15
}


⚠️ This breaks logic:

Multiple conditions may execute

Price can be overwritten

✔ else if prevents this

9️⃣ Syntax rule summary (memorize this)
if / else if / else rules in Go

Condition must be boolean

No parentheses required

Braces {} are mandatory

else if avoids multiple matches

else is the fallback

🔟 Mental model (very important)

Think of it like:

“Start with free ticket → upgrade price only if age qualifies”

That’s why:

price := 0

✅ Final takeaway

Zero value is intentional, not random

else if enforces single decision

Initialize → evaluate → assign → print

This pattern is used everywhere in real Go projects