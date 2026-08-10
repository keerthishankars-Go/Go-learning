1️⃣ First: Core rule in Go

Go has only ONE loop keyword: for

There is:

❌ no while

❌ no do-while

Everything is done using variations of for.

2️⃣ Standard for loop form (recap)
for init; condition; post {
	// body
}


Example:

for i := 0; i <= 10; i++ {
	fmt.Println(i)
}


All three parts are optional.

3️⃣ First “while-like” version (with semicolons)
Code
i := 0
for ; i <= 10; {   // init and post omitted
	fmt.Printf("%d ", i)
	i += 2
}

Syntax-level explanation

i := 0
→ Initialization happens outside the loop

for ; i <= 10; {
→ First ; means no initialization
→ Condition controls the loop
→ Second ; means no post statement

i += 2
→ Increment is done inside the loop body

📌 This behaves exactly like:

while (i <= 10) {
	i += 2
}


in other languages.

4️⃣ Why semicolons can be removed

Go allows this syntax:

for condition {
	// body
}


So this:

for ; i <= 10; {


can be rewritten as:

for i <= 10 {


Both mean the same.

5️⃣ Second version (pure while-style)
Code
i := 0
for i <= 10 {
	fmt.Printf("%d ", i)
	i += 2
}

Why this is called a “while loop”

Because:

Only a condition is present

Loop runs while condition is true

No init or post inside for

📌 Mental translation

for i <= 10 { }   // while (i <= 10)

6️⃣ Why Go designers did this (important reasoning)

Instead of adding:

for

while

do-while

Go keeps:
✔ One keyword
✔ Multiple forms
✔ Less syntax to learn
✔ Clear behavior

This reduces language complexity.

7️⃣ What happens if you omit the condition too?
for {
	fmt.Println("infinite loop")
}


This is equivalent to:

while (true) { }


Used in:

servers

event loops

workers

8️⃣ Execution flow (step-by-step)

For this code:

i := 0
for i <= 10 {
	fmt.Printf("%d ", i)
	i += 2
}

Steps

i = 0

Check i <= 10 → true

Print 0

i = 2

Check again

Repeat until i = 12

Condition fails → loop exits

9️⃣ Why increment is inside the loop

Because:

There is no post section

You must control loop progression manually

📌 Rule

When using while-style for, ensure progress happens inside the loop
Otherwise → infinite loop ❌

🔟 Real-world backend use cases
🔹 Retry logic
for retries <= maxRetries {
	retries++
}

🔹 Polling
for !jobDone {
	checkStatus()
}

🔹 Reading until EOF
for err == nil {
	data, err = read()
}

✅ Rules to remember (save this)

Go has no while

for is the only loop keyword

for condition {} = while loop

Semicolons are optional when only condition exists

Ensure loop variable changes inside loop

Infinite loop = for {}

🧠 Mental model

Go uses for everywhere — just change how much you write inside it.


========================================================================

This is a fantastic question because it teaches you **how the order of statements inside a loop changes the output**.

Let's compare them step by step.

---

# Program 1

```go
package main

import "fmt"

func main() {
	i := 0

	for i <= 10 {

		i += 2

		fmt.Printf("%d ", i)
	}
}
```

### Execution

Initially

```
i = 0
```

### First iteration

Condition

```
0 <= 10 ✔
```

First statement

```go
i += 2
```

Now

```
i = 2
```

Print

```
2
```

---

### Second iteration

```
i = 2
```

Condition

```
2 <= 10 ✔
```

Increase

```
i = 4
```

Print

```
4
```

This continues...

Output

```
2 4 6 8 10 12
```

Notice something interesting.

Even though the loop condition is

```go
i <= 10
```

it still prints **12**.

Why?

Because the condition is checked **before** entering the loop body.

When `i` is 10:

```
10 <= 10 ✔
```

So Go enters the loop.

Then inside the loop

```
i += 2
```

becomes

```
12
```

and **then** it prints 12.

---

# Program 2

```go
package main

import "fmt"

func main() {

	i := 0

	for i <= 10 {

		fmt.Printf("%d ", i)

		i += 2
	}
}
```

---

Initially

```
i = 0
```

### First iteration

Condition

```
0 <= 10 ✔
```

Print

```
0
```

Increase

```
2
```

---

Next

Print

```
2
```

Increase

```
4
```

...

Output

```
0 2 4 6 8 10
```

When

```
i = 10
```

Prints

```
10
```

Then

```
i += 2
```

becomes

```
12
```

Loop goes back.

Condition

```
12 <= 10 ❌
```

Loop stops.

So 12 is **never printed**.

---

# Side-by-side

### Program 1

```go
i += 2
fmt.Println(i)
```

Flow

```
0
 ↓
+2
 ↓
2 printed
```

Output

```
2 4 6 8 10 12
```

---

### Program 2

```go
fmt.Println(i)
i += 2
```

Flow

```
0 printed
 ↓
+2
 ↓
2
```

Output

```
0 2 4 6 8 10
```

---

# Visual Timeline

### Program 1

```
Start

i=0

Condition ✔

Increase → 2

Print → 2

----------------

Condition ✔

Increase → 4

Print → 4

...

Condition (10<=10) ✔

Increase → 12

Print → 12

Condition (12<=10) ✘
```

---

### Program 2

```
Start

i=0

Condition ✔

Print → 0

Increase → 2

----------------

Condition ✔

Print → 2

Increase → 4

...

Condition ✔

Print →10

Increase →12

Condition ✘
```

---

## The key lesson

Inside a loop, **statements execute in the exact order you write them**.

* If you **update first, then print**, you see the updated value.
* If you **print first, then update**, you see the current value before it changes.

This idea applies to all loops (`for`, `range`) and is a common source of off-by-one errors, so it's worth getting comfortable tracing the execution step by step.
