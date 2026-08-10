Great. You have already understood channels, so **WaitGroup will be much easier**. The key difference:

* **Channel** → communication between goroutines (send/receive data)
* **WaitGroup** → synchronization (wait until goroutines finish)

Think:

```
Channel:
"Here is some data"

WaitGroup:
"I am done, you can continue"
```

Let's deeply understand the syntax and execution.

---

# 1. Why do we need WaitGroup?

Imagine this:

```go
func main() {

    go process(1)
    go process(2)
    go process(3)

    fmt.Println("main finished")
}
```

Execution:

```
main
 |
 +---- process 1
 |
 +---- process 2
 |
 +---- process 3
```

Problem:

Main goroutine does not wait.

It may finish immediately:

```
main finished
(program exits)
```

Before:

```
process 1
process 2
process 3
```

complete.

So we need:

> "Main, wait until all goroutines finish."

That is WaitGroup.

---

# 2. Import sync

```go
import "sync"
```

Why?

Because WaitGroup is inside Go's `sync` package.

Think:

```
sync package

    |
    |
    +---- WaitGroup
    +---- Mutex
    +---- Once
```

---

# 3. Create WaitGroup

Inside main:

```go
var wg sync.WaitGroup
```

Break syntax:

General:

```go
var variableName Type
```

Examples:

```go
var age int

var name string
```

Here:

```go
var wg sync.WaitGroup
```

means:

```
variable name:

wg

type:

sync.WaitGroup
```

Memory:

```
wg

counter = 0
```

Important:

WaitGroup internally maintains a counter.

---

# 4. The process function

```go
func process(i int, wg *sync.WaitGroup)
```

Let's break.

Parameters:

| Name | Type            |
| ---- | --------------- |
| i    | int             |
| wg   | *sync.WaitGroup |

The second one:

```go
wg *sync.WaitGroup
```

means:

> A pointer to a WaitGroup.

Why pointer?

Because we want every goroutine to access the SAME WaitGroup.

---

Imagine without pointer:

```go
func process(i int, wg sync.WaitGroup)
```

When calling:

```go
go process(i, wg)
```

Go copies:

```
Main:

wg
counter = 3


Goroutine 1:

copy of wg
counter = 3


Goroutine 2:

copy of wg
counter = 3
```

Now when goroutine calls:

```go
wg.Done()
```

it decreases its own copy.

Main's counter never changes.

Bad.

---

With pointer:

```go
go process(i, &wg)
```

Memory:

```
             same object


Main
 |
 |
 v

wg
counter = 3
 ^
 |
 |
Goroutine 1

Goroutine 2

Goroutine 3
```

Everyone modifies the same counter.

---

# 5. Adding goroutines

Main:

```go
no := 3
```

We want 3 goroutines.

---

Loop:

```go
for i := 0; i < no; i++ {
```

Runs:

```
i=0
i=1
i=2
```

Three times.

---

Inside:

```go
wg.Add(1)
```

This increases WaitGroup counter.

Initially:

```
counter = 0
```

First loop:

```
Add(1)

counter = 1
```

Second:

```
Add(1)

counter = 2
```

Third:

```
Add(1)

counter = 3
```

Now Go knows:

```
3 goroutines are running
```

---

# 6. Start goroutine

```go
go process(i, &wg)
```

Example first iteration:

```
process(0, address of wg)
```

Second:

```
process(1, address of wg)
```

Third:

```
process(2, address of wg)
```

Now:

```
Main

      |
      |
 -----------------
 |       |       |
 v       v       v

P0      P1      P2
```

---

# 7. Inside process()

First:

```go
fmt.Println("started Goroutine ", i)
```

Example:

```
started Goroutine 0
```

---

Sleep:

```go
time.Sleep(2 * time.Second)
```

Meaning:

Pause this goroutine for 2 seconds.

Important:

It does NOT pause all goroutines.

Only that goroutine.

Example:

```
Goroutine 0 sleeping
Goroutine 1 sleeping
Goroutine 2 sleeping
```

They sleep independently.

---

# 8. The most important line

```go
wg.Done()
```

This decreases the counter.

Remember:

Before:

```
wg counter = 3
```

Goroutine 0 finishes:

```
Done()

counter = 2
```

Goroutine 1:

```
Done()

counter = 1
```

Goroutine 2:

```
Done()

counter = 0
```

---

# 9. Wait()

Main:

```go
wg.Wait()
```

Meaning:

> Block here until WaitGroup counter becomes zero.

At this moment:

```
counter = 3
```

So:

```
main waits
```

Timeline:

```
MAIN

wg.Wait()
 |
 |
 | waiting...
 |
 |
 +----------------+
                  |
                  v

             counter = 0

                  |
                  v

             continue
```

Then:

```go
fmt.Println("All go routines finished executing")
```

runs.

---

# Complete execution

```
main starts

wg counter = 0


Loop starts


Iteration 1:

Add(1)

counter = 1

start goroutine 0



Iteration 2:

Add(1)

counter = 2

start goroutine 1



Iteration 3:

Add(1)

counter = 3

start goroutine 2



main reaches:

wg.Wait()


BLOCK


--------------------------------


Goroutine 0 finishes

Done()

counter = 2


Goroutine 1 finishes

Done()

counter = 1


Goroutine 2 finishes

Done()

counter = 0


--------------------------------


Wait releases


main continues


print:

All go routines finished executing
```

---

# How to explain in an interview

You can say:

> "WaitGroup is used for goroutine synchronization. It maintains an internal counter. Before starting each goroutine, I increment the counter using Add(). Each goroutine calls Done() when its execution completes, which decrements the counter. The main goroutine calls Wait(), which blocks until the counter becomes zero. I pass the WaitGroup as a pointer because all goroutines must update the same counter instead of working on copies."

---

# Relation with channels

You now know two synchronization patterns:

## Channel

```
Goroutine A
      |
      |
      v
   channel
      |
      |
      v
Goroutine B
```

Purpose:

> Transfer data

---

## WaitGroup

```
Goroutine 1
Goroutine 2
Goroutine 3

      |
      |
      v

 WaitGroup counter

      |
      |
      v

 main continues
```

Purpose:

> Wait for completion

In real Go backend systems you often use both together:

```
API request
    |
    |
start 10 workers
    |
    |
channels send results
    |
    |
WaitGroup waits for workers
    |
    |
return response
```

That is exactly where worker pools come next.
