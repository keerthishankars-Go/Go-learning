Sending and receiving from a channel
The syntax to send and receive data from a channel is given below,

data := <- a // read from channel a
a <- data // write to channel a
The direction of the arrow with respect to the channel specifies whether the data is sent or received.

In the first line, the arrow points outwards from a and hence we are reading from channel a and storing the value to the variable data.

In the second line, the arrow points towards a and hence we are writing to channel a.

Sends and receives are blocking by default
Sends and receives to a channel are blocking by default. What does this mean? When data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly, when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages.

============================================================================

Excellent! This is your **first synchronization program** in Go. It introduces **goroutines** and **channels**. Like we did with interfaces, let's understand **how to think about writing this syntax**, not just what it does.

---

# First, the complete program

```go
package main

import (
	"fmt"
)

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

# High-Level Picture

Think of two people communicating.

```text
          MAIN GOROUTINE
                 │
                 │
      Creates a channel
                 │
                 ▼
           done channel
                 ▲
                 │
                 │
       HELLO GOROUTINE
```

The channel is like a **pipe** connecting two goroutines.

---

# Step 1: Program starts

Only one goroutine exists initially.

```text
Runtime

┌────────────────────────┐
│ Main Goroutine         │
└────────────────────────┘
```

Execution begins:

```go
func main() {
```

---

# Step 2: Create a channel

```go
done := make(chan bool)
```

Let's understand this syntax.

## General Syntax

Creating things in Go often follows:

```go
variable := make(Type)
```

Examples:

```go
make([]int, 5)      // slice
make(map[string]int) // map
make(chan bool)     // channel
```

Here:

```go
make(chan bool)
```

means

> Create a channel that carries values of type `bool`.

---

Memory now:

```text
Main Goroutine

done
 │
 ▼

Channel<bool>

(empty)
```

Nothing is inside yet.

---

# Step 3: Launch a goroutine

```go
go hello(done)
```

This syntax confuses many beginners.

Let's split it.

Normally you call:

```go
hello(done)
```

which means

```text
Call hello()

↓

Wait until hello finishes

↓

Continue
```

Timeline:

```text
main

hello()

wait

continue
```

---

Now add one word.

```go
go hello(done)
```

That one keyword changes everything.

Meaning:

> Start executing `hello(done)` in a **new goroutine**, and don't wait for it.

Timeline:

```text
Main Goroutine

continue immediately
|
|
v

New Goroutine

hello(done)
```

Now there are two goroutines.

```text
               Runtime

     ┌──────────────────────┐
     │ Main Goroutine       │
     └──────────────────────┘
                │

     ┌──────────────────────┐
     │ Hello Goroutine      │
     └──────────────────────┘
```

---

# Step 4: hello() begins

The new goroutine runs:

```go
func hello(done chan bool) {
```

Now understand this parameter.

```go
done chan bool
```

Read it like English.

> done is a channel carrying bool values.

Exactly like

```go
age int
```

means

```text
age

Type

int
```

Similarly

```go
done chan bool
```

means

```text
done

Type

channel of bool
```

---

# Step 5: Print

```go
fmt.Println("Hello world goroutine")
```

Output:

```text
Hello world goroutine
```

Nothing special.

---

# Step 6: The most important line

```go
done <- true
```

This introduces the **channel send operator**.

Read it as:

> Send `true` into the channel `done`.

Think of the arrow direction.

```text
true

   │

   ▼

done channel
```

Or

```text
true -----> done
```

The arrow points **towards the channel**.

---

Memory:

Before:

```text
done

(empty)
```

After send:

```text
done

true
```

Actually, because this is an **unbuffered channel**, the value is **not stored**. The sender and receiver meet directly. We'll come back to that.

---

# Step 7: Meanwhile, main continues

After starting the goroutine:

```go
go hello(done)
```

main immediately executes

```go
<-done
```

This is the receive operator.

Read it as

> Receive a value from the channel.

Notice the direction.

Sending:

```go
done <- true
```

Arrow points **toward** channel.

Receiving:

```go
<-done
```

Arrow points **away from** channel.

Think:

```text
Channel

↓

receive
```

or

```text
done -----> true
```

---

# Visualize both together

Sending:

```go
done <- true
```

```text
true

 │

 ▼

CHANNEL
```

Receiving:

```go
<-done
```

```text
CHANNEL

 │

 ▼

receiver
```

---

# What happens internally?

Main reaches

```go
<-done
```

Channel is empty.

So main **blocks**.

```text
Main

Waiting...

Waiting...

Waiting...
```

The scheduler pauses it.

---

Meanwhile

Hello goroutine reaches

```go
done <- true
```

Now

```text
Sender waiting

Receiver waiting
```

Both are ready.

Go connects them.

```text
Hello

true
 │
 ▼
CHANNEL
 │
 ▼
Main
```

The value transfers instantly.

No storage.

---

Then

Hello goroutine ends.

Main resumes.

```go
fmt.Println("main function")
```

Output:

```text
main function
```

---

# Why use bool?

Actually, the value isn't important.

You're using the channel only as a **signal**.

Like saying

```text
"I'm finished."
```

Could have been

```go
done <- true
```

or

```go
done <- false
```

Main doesn't even use it.

It simply waits.

---

# Better way (common Go)

Instead of

```go
chan bool
```

many programs use

```go
chan struct{}
```

because nothing needs to be transferred.

Only the signal matters.

---

# Complete Timeline

```text
Program Starts

↓

Main Goroutine

↓

make(chan bool)

↓

Channel created

↓

go hello(done)

↓

Hello Goroutine starts

↓

Main reaches

<-done

↓

Main blocks

──────────────────────────────

Hello prints

Hello world goroutine

↓

Hello executes

done <- true

↓

Main receives signal

↓

Main resumes

↓

Print

main function
```

Output:

```text
Hello world goroutine
main function
```

---

# The syntax you'll use in almost every Go backend

Whenever you see these patterns:

### Create a channel

```go
ch := make(chan T)
```

> "Create a communication pipe that carries values of type `T`."

### Send

```go
ch <- value
```

> "Send `value` into channel `ch`."

### Receive

```go
value := <-ch
```

> "Wait until a value arrives from `ch`, then assign it to `value`."

### Launch a goroutine

```go
go someFunction()
```

> "Run `someFunction()` concurrently in a new goroutine and continue immediately."

---

## Mental model to remember

Think of channels as a **door** between goroutines.

```text
                 Main Goroutine
                       │
              waits at the door
                       │
                    <- done
                       │
════════════════════════════════════
               CHANNEL (door)
════════════════════════════════════
                       │
                 done <- true
                       │
             Hello Goroutine
```

The sender opens the door by sending, the receiver opens the door by receiving, and the value passes through. With an **unbuffered channel** like this one, neither side proceeds until both are present—this is why `<-done` is such a common way to wait for a goroutine to finish.

========================================================================


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
