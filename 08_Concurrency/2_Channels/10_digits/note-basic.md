For a **live demo/interview**, don't start by typing the full complex version immediately. Build it progressively so the interviewer can see your thinking.

For this program, I would write in this order:

---

## Step 1: First show the basic problem (without goroutines)

Start with the simplest logic:

> "First I will write a function to extract digits."

```go
package main

import "fmt"

func digits(number int) {
	for number != 0 {
		digit := number % 10
		fmt.Println(digit)
		number /= 10
	}
}

func main() {
	digits(589)
}
```

Output:

```
9
8
5
```

Explain:

> "This extracts digits one by one."

---

# Step 2: Introduce channel

Now the question:

> "Instead of printing, I want to send these digits to another function."

Change:

```go
func digits(number int)
```

to:

```go
func digits(number int, dchnl chan int)
```

Now:

```go
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}

	close(dchnl)
}
```

Explain:

> "The producer sends digits through the channel and closes it when finished."

---

# Step 3: Add consumer

Now create the receiver.

```go
func calcSquares(number int, squareop chan int) {

	sum := 0

	dch := make(chan int)

	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit
	}

	squareop <- sum
}
```

Explain:

Flow:

```
digits()
    |
    |
    v
 channel
    |
    |
    v
calcSquares()
```

---

# Step 4: Add cube calculation

Now duplicate the pattern:

```go
func calcCubes(number int, cubeop chan int) {

	sum := 0

	dch := make(chan int)

	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit * digit
	}

	cubeop <- sum
}
```

Explain:

> "Same digit stream, different calculation."

---

# Step 5: Finally write main()

Now connect everything.

```go
func main() {

	number := 589

	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)

	squares := <-sqrch
	cubes := <-cubech

	fmt.Println("Final output", squares+cubes)
}
```

---

# The order in your head should be:

```
1. What work needs to happen?
        |
        v
2. Create functions
        |
        v
3. Need communication?
        |
        v
4. Add channels
        |
        v
5. Need parallel execution?
        |
        v
6. Add go keyword
        |
        v
7. Receive results
        |
        v
8. Print output
```

---

# For interview/demo, I would NOT start with this:

```go
go calcSquares()
go calcCubes()
```

because the interviewer cannot see where the logic came from.

Start with:

```
digits()
      |
      v
channel
      |
      v
calculator
      |
      v
main
```

Then add concurrency.

---

# The final demo explanation (30 seconds)

You can say:

> "I have separated this into producer and consumer pattern. The digits function is a producer that extracts digits and sends them through a channel. The square and cube calculators consume those digits using range over the channel. Once the producer finishes, it closes the channel. The main goroutine starts both calculations concurrently and waits for both results before printing the final output."

That explanation sounds like real Go backend thinking, not just code memorization.

============================================================================

This is one of the **most common interview questions**. The interviewer may ask:

> **"Why do we use `% 10` and `/= 10`?"**

Don't memorize it. Understand **the mathematics behind it**.

---

# Imagine the number

```text
589
```

Think of it as three boxes.

```text
+-----+-----+-----+
|  5  |  8  |  9  |
+-----+-----+-----+
```

Suppose I ask:

> **How can you get only the last digit?**

Go gives us the modulo operator `%`.

---

# Step 1: `% 10` gives the last digit

```go
digit := number % 10
```

Suppose

```go
number = 589
```

Calculation:

```text
589 ÷ 10

Quotient  = 58
Remainder = 9
```

So

```go
589 % 10
```

returns

```text
9
```

Why?

Because

```text
589 = (58 × 10) + 9
```

The remainder is always the last digit.

---

## More examples

```text
123 % 10 = 3

4567 % 10 = 7

908 % 10 = 8

20 % 10 = 0
```

So a rule emerges:

```text
number % 10

↓

Always gives the LAST digit
```

---

# Step 2: `/= 10` removes the last digit

Now suppose we've already processed **9**.

We don't want it again.

How do we remove it?

```go
number /= 10
```

This means

```go
number = number / 10
```

Since `number` is an integer:

```text
589 / 10

↓

58
```

The decimal part is discarded.

So

```text
589

↓

58
```

The last digit is gone.

---

# Why do we need both?

Let's see what happens **without `/= 10`**.

```go
number := 589

for number != 0 {

    digit := number % 10

    fmt.Println(digit)

}
```

Iteration 1

```text
589 % 10

↓

9
```

Iteration 2

Still

```text
number

589
```

Again

```text
589 %10

↓

9
```

Again

```text
9
```

Infinite loop.

Because:

```text
number

never changes.
```

---

Now add

```go
number /= 10
```

Iteration 1

```text
number = 589

digit = 9

number = 58
```

Iteration 2

```text
number = 58

digit = 8

number = 5
```

Iteration 3

```text
number = 5

digit = 5

number = 0
```

Loop stops.

---

# Visualize it

Initial

```text
589
```

Take last digit

```text
589 %10

↓

9
```

Remove last digit

```text
589 /10

↓

58
```

Now repeat

```text
58 %10

↓

8
```

Remove

```text
58 /10

↓

5
```

Repeat

```text
5 %10

↓

5
```

Remove

```text
5 /10

↓

0
```

Finished.

---

# Think of peeling an onion 🧅

Imagine every loop removes one layer.

```text
589

↓

9 removed

↓

58

↓

8 removed

↓

5

↓

5 removed

↓

0
```

---

# This is why the loop works

```go
for number != 0 {

    digit := number % 10

    // Process digit

    number /= 10
}
```

Each iteration does **two jobs**:

### Job 1

```go
digit := number % 10
```

Gets the last digit.

---

### Job 2

```go
number /= 10
```

Removes the last digit so the next iteration can process the next one.

---

# Interview answer (30 seconds)

If the interviewer asks:

> **Why do you use `%10` and `/10`?**

You can answer:

> "The modulo operator `%10` extracts the last digit of the number because dividing by 10 leaves the last digit as the remainder. After processing that digit, I use integer division `/=10` to remove the last digit from the number. By repeating these two operations inside a loop, I can process every digit from right to left until the number becomes zero."

---

## The mental formula to remember

Whenever you need to process the digits of an integer:

```go
for number != 0 {

    digit := number % 10   // Extract the last digit

    // Process the digit

    number /= 10           // Remove the last digit
}
```

This is a standard algorithm used in many interview problems such as:

* Sum of digits
* Reverse a number
* Palindrome number
* Armstrong number
* Sum of squares/cubes of digits
* Counting digits

Once you recognize this pattern, you'll immediately know **`% 10` means "extract the current last digit"** and **`/ 10` means "move on to the next digit."**
