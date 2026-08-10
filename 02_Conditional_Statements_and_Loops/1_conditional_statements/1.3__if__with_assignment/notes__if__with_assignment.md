This if with assignment syntax is very Go-specific and used a lot in real code, so let’s explain it purely at syntax + reasoning level, slowly.

1️⃣ The special if syntax (core idea)
if assignment-statement; condition {
    // block
}


This means:

First, Go executes the assignment

Then, it evaluates the condition

If condition is true → enter block

📌 The semicolon ; is mandatory here
It separates:

setup part

decision part

2️⃣ Your example (cleaned)
ticketPrice := 0

if age := 10; age < 5 {
	ticketPrice = 0
} else if age >= 5 && age <= 22 {
	ticketPrice = 10
} else {
	ticketPrice = 15
}

fmt.Printf("Ticket price is $%d", ticketPrice)


Now let’s break why it is written like this.

3️⃣ age := 10 inside if
if age := 10; age < 5 {

What this does (step by step)

age := 10

Declares a variable age

Assigns value 10

This happens before checking age < 5

Equivalent to:

age := 10
if age < 5 {
    ...
}


But with one big difference 👇

4️⃣ Scope of age (VERY IMPORTANT)

The variable age declared here:

if age := 10; ...


👉 Exists ONLY inside this if–else chain

Valid scope:

if age := 10; age < 5 {
	// age is accessible
} else if age <= 22 {
	// age is accessible
} else {
	// age is accessible
}


❌ Invalid (compiler error):

fmt.Println(age) // age is undefined here

5️⃣ WHY Go allows this syntax (reasoning)

This syntax exists to:

✔ Avoid polluting outer scope
✔ Keep variables short-lived
✔ Prevent accidental reuse
✔ Improve readability

Think like a professional:

“This variable is only needed to make a decision — nowhere else.”

So Go forces it to die after the if.

6️⃣ Why ticketPrice is declared outside
ticketPrice := 0

Why not inside if?

Because:

We need ticketPrice after the if

Scope must be wider

📌 Rule:

Declare variables in the smallest scope that still allows usage

So:

age → decision-only → inside if

ticketPrice → used later → outside

7️⃣ Why initialize ticketPrice to 0
ticketPrice := 0

Syntax + logic reason

0 is the zero value

Free ticket case already satisfied

Avoids uninitialized state

Defensive programming

Even if no branch runs (hypothetically), value is safe.

8️⃣ Why else if works with same age
else if age >= 5 && age <= 22 {


Even though age was declared in if, it is still accessible because:

📌 Scope rule

Variables declared in if initialization are visible in:

if

all else if

else

But nowhere else.

9️⃣ Why semicolon ; is required
if age := 10; age < 5 {


Because Go grammar needs to know:

Where assignment ends

Where condition begins

Without ;, this would be ambiguous.

❌ This is invalid:

if age := 10 age < 5 { }

🔟 Real-world backend examples (you WILL see this)
Example 1: Error handling
if err := doSomething(); err != nil {
	return err
}

Example 2: Map lookup
if value, ok := cache[key]; ok {
	fmt.Println(value)
}

Example 3: File open
if file, err := os.Open(name); err == nil {
	defer file.Close()
}


These are idiomatic Go.

🧠 Mental model (remember forever)

“Create → check → discard”

Create variable

Use it for decision

Let it die immediately

✅ Rules to remember (short & strict)

if assignment; condition {} is valid Go

Assignment runs before condition

Variable scope is limited to the if–else chain

Semicolon ; is mandatory

Use this when variable is only for decision logic