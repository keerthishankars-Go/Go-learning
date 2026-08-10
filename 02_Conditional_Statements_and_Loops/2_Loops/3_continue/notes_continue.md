🔹 Full program (reference)
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

1️⃣ What is continue (in one sentence)

continue skips the rest of the current loop iteration and jumps to the next iteration.

It does NOT:

stop the loop

exit the loop

It only skips what comes after it in the loop body.

2️⃣ The for loop syntax (why written like this)
for i := 1; i <= 10; i++ {


This has three parts, separated by ;

🔹 i := 1

Initialization

Runs once, before the loop starts

🔹 i <= 10

Condition

Checked before every iteration

If false → loop stops

🔹 i++

Post statement

Runs after each iteration

📌 Rule

for init; condition; post {} is Go’s classic loop form.

3️⃣ The if condition (key logic)
if i%2 == 0 {

Why % (modulus)?

% gives the remainder

i % 2 == 0 → number is even

Examples:

2 % 2 = 0 → even

3 % 2 = 1 → odd

📌 We want to skip even numbers, not print them.

4️⃣ Why continue is used here
continue

What happens when continue executes?

Go immediately stops executing the current loop body

Skips:

fmt.Printf("%d ", i)


Jumps to:

i++


Re-checks:

i <= 10


📌 Everything below continue in that iteration is ignored

5️⃣ Step-by-step execution (VERY IMPORTANT)

Let’s simulate:

🔹 i = 1

1 % 2 == 0 → false

continue not executed

fmt.Printf runs → prints 1

🔹 i = 2

2 % 2 == 0 → true

continue executes

fmt.Printf skipped

nothing printed

🔹 i = 3

odd → printed

🔹 i = 4

even → skipped

This repeats until 10.

6️⃣ Why output is: 1 3 5 7 9

Because:

Even numbers trigger continue

Print statement is skipped

Only odd numbers reach fmt.Printf

7️⃣ Why NOT use else instead?

You could write:

if i%2 != 0 {
	fmt.Printf("%d ", i)
}


This also works, but:

Why continue is often preferred

Reduces nesting

Clear “skip this case” intent

Cleaner for complex loops

📌 Professional rule

Use continue to skip unwanted cases early.

8️⃣ Difference between continue and break (don’t confuse)
Statement	What it does
continue	Skip current iteration
break	Exit the loop completely

Example:

if i == 5 {
	break
}


→ loop stops at 5

9️⃣ Real-world backend use cases (important)
🔹 Skip invalid data
for _, user := range users {
	if user.IsBlocked {
		continue
	}
	process(user)
}

🔹 Skip empty values
for _, v := range inputs {
	if v == "" {
		continue
	}
}

🔹 Skip failed validations
for _, req := range requests {
	if !isValid(req) {
		continue
	}
	handle(req)
}

🔟 Rules to remember (save this)

continue works only inside loops

It skips remaining statements in the loop body

Loop continues with next iteration

Commonly used to filter data

Helps keep code flat and readable

🧠 Mental model

“This iteration is not useful — skip it and move on.”