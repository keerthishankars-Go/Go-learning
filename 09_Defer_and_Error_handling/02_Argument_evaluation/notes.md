This is a **very important `defer` concept**. This example is specifically showing:

> **Arguments passed to a deferred function are evaluated immediately when `defer` is executed, not when the deferred function actually runs.**

Let's execute step by step.

---

## Code


func displayValue(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	a := 5

	defer displayValue(a)

	a = 10

	fmt.Println("value of a before deferred function call", a)
}


---

# Step 1

Program enters:


func main()


Creates:


a := 5


Memory:


a
|
5


---

# Step 2 (most important)

Now:


defer displayValue(a)


Many beginners think:

> "It will execute later, so it will use the latest value of a."

But no.

At this exact moment:


a = 5


Go evaluates the argument immediately.

It is equivalent to:


defer displayValue(5)


Go remembers:


Call displayValue with value 5 later


The deferred function is now waiting.

---

# Step 3

Next line:


a = 10


Now memory changes:


a
|
10


But the deferred function already captured:


5


It does not know about this change.

---

# Step 4

Now:


fmt.Println("value of a before deferred function call", a)


Current value:


a = 10


Output:


value of a before deferred function call 10


---

# Step 5

main function is about to return.

Before returning, Go executes deferred functions.

Remember:


defer displayValue(5)


So it calls:


displayValue(5)


Not:


displayValue(10)


---

Output:


value of a before deferred function call 10
value of a in deferred function 5


---

# Visual timeline


main starts

a = 5


defer displayValue(a)

        |
        |
        v

Go stores:

displayValue(5)


a = 10


Print a

Output:
10


main ends


Run deferred function:

displayValue(5)


Output:
5


---

# Why does Go do this?

Because function arguments are normally evaluated immediately.

Example:


x := 100

print(x)

x = 200


When `print(x)` happens, it receives:


100


not:


200


`defer` follows the same rule.

The only difference is **execution time**.

---

# Compare these two cases

## Case 1: Passing value


a := 5

defer displayValue(a)

a = 10


Output:


5


Because value was copied.

---

## Case 2: Passing pointer

Now:


func displayValue(a *int) {
	fmt.Println(*a)
}

func main() {

	a := 5

	defer displayValue(&a)

	a = 10
}


Output:


10


Why?

Because you passed the address.

Timeline:


a
|
10


pointer
|
v

same memory location


The deferred function reads the latest value.

---

# Real backend connection

This matters in production code.

Example:


start := time.Now()

defer func() {
	fmt.Println(time.Since(start))
}()


Here we use a closure.

Why?

Because we want the **latest value of start** when defer executes.

---

Another example:


mutex.Lock()

defer mutex.Unlock()


The function call itself is delayed, but the `mutex` object reference is already captured.

---

# Interview answer

If asked:

**"When are deferred function arguments evaluated?"**

Answer:

> "Arguments passed to a deferred function are evaluated immediately when the defer statement executes, but the function call itself is delayed until the surrounding function returns."

This one concept prevents many subtle Go bugs.
