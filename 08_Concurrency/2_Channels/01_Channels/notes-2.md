This is exactly how you should learn Go. Don't memorize programs—learn the **construction order**. Every Go program follows a pattern.

Let's build your program **from scratch**, exactly how an experienced Go developer would write it.

---

# Goal

We want:

```text
Main Program
      |
      |
Start another goroutine
      |
      |
Wait until it finishes
      |
      |
Continue main
```

So we ask ourselves:

> **What do I need?**

---

# Step 1 — Package

Every executable Go program starts with

```go
package main
```

Meaning

> This is an executable application.

---

# Step 2 — Import packages

We need printing.

```go
import "fmt"
```

Ask yourself:

> Am I using `fmt.Println()`?

Yes.

Therefore import fmt.

---

# Step 3 — Decide what the goroutine should do

Question:

> What work should happen in parallel?

Suppose:

```text
Print

Hello world goroutine
```

Create a function.

```go
func hello() {
    fmt.Println("Hello world goroutine")
}
```

At this point your program is

```go
package main

import "fmt"

func hello() {
    fmt.Println("Hello world goroutine")
}
```

Question:

Can this run in parallel?

Yes.

Can main know when it finishes?

No.

That's our next problem.

---

# Step 4 — We need communication

Ask yourself

```text
How will hello()

tell

main()

"I'm finished."
```

Need communication.

Go's answer:

```text
CHANNEL
```

---

# Step 5 — Add a channel parameter

Modify

```go
func hello() {
```

to

```go
func hello(done chan bool) {
```

Read this like English.

```text
hello()

needs

a channel named done

which carries bool values
```

Nothing else changed.

---

# Step 6 — Send a signal

When hello finishes

it should notify main.

Question:

How do we send into a channel?

Syntax

```go
channel <- value
```

So

```go
done <- true
```

Now

```go
func hello(done chan bool) {
    fmt.Println("Hello world goroutine")
    done <- true
}
```

Read it like English.

```text
Print

↓

Send

true

↓

Finish
```

---

# Step 7 — Now write main()

```go
func main() {

}
```

---

# Step 8 — Create the channel

Question

Where does the channel come from?

Need to create one.

Syntax

```go
make(chan bool)
```

General rule

```go
make(type)
```

Examples

```go
make(map[string]int)

make([]int,5)

make(chan bool)
```

Create

```go
done := make(chan bool)
```

Now

```text
Main

↓

creates

↓

done channel
```

---

# Step 9 — Start the goroutine

Question

How do we call a function normally?

```go
hello(done)
```

But that blocks.

Need concurrency.

One keyword.

```go
go hello(done)
```

Read it like English.

```text
Start hello()

in another goroutine

and continue immediately
```

---

# Step 10 — Wait

Question

How does main know hello finished?

Receive from channel.

Syntax

```go
<-done
```

Meaning

```text
Wait

until

someone sends

a value
```

---

# Step 11 — Continue

Now

```go
fmt.Println("main function")
```

---

# Final program

```go
package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)

	go hello(done)

	<-done

	fmt.Println("main function")
}
```

---

# The Thinking Process

This is how a Go developer thinks.

```text
Need another task

↓

Create function

↓

Need to run simultaneously

↓

Add

go

↓

Need communication

↓

Create channel

↓

Pass channel

↓

Send signal

↓

Receive signal

↓

Continue
```

---

# Syntax Formation Order

Instead of memorizing code, memorize this template.

### 1. Create worker function

```go
func worker(...) {

}
```

---

### 2. Add communication

```go
func worker(done chan bool) {

}
```

---

### 3. Do work

```go
func worker(done chan bool) {

    // work

}
```

---

### 4. Notify completion

```go
done <- true
```

---

### 5. Create channel

```go
done := make(chan bool)
```

---

### 6. Launch goroutine

```go
go worker(done)
```

---

### 7. Wait

```go
<-done
```

---

### 8. Continue

```go
fmt.Println(...)
```

---

# A Flow Diagram

```text
                   PROGRAM

                       │

                       ▼

            Define worker function

                       │

                       ▼

      Does worker need to tell main something?

                 Yes

                       │

                       ▼

          Add parameter

          done chan bool

                       │

                       ▼

          Worker completes

                       │

                       ▼

              done <- true

────────────────────────────────────────────

                   main()

                       │

                       ▼

         done := make(chan bool)

                       │

                       ▼

            go worker(done)

                       │

                       ▼

                <-done

                       │

                       ▼

          Continue remaining code
```

---

## The 5 patterns you'll write in almost every concurrent Go program

```go
// 1. Create a channel
ch := make(chan T)

// 2. Start a goroutine
go worker(ch)

// 3. Send a value or signal
ch <- value

// 4. Receive a value or wait for a signal
value := <-ch
// or simply
<-ch

// 5. Continue after synchronization
```

If you become comfortable with just these five patterns, you'll be able to understand most beginner and intermediate Go concurrency examples, and you'll be well prepared to learn buffered channels, `select`, worker pools, and pipelines.
