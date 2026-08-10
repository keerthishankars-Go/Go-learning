Mutex
A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.

Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.

mutex.Lock()
x = x + 1
mutex.Unlock()

In the above code, x = x + 1 will be executed by only one Goroutine at any point in time thus preventing race condition.

If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will be blocked until the mutex is unlocked.

=============================================================================

Good. This is the **next step after WaitGroup**. Now you are entering **shared memory problems**.

Before understanding the syntax, understand the problem:

> "What happens when multiple goroutines try to modify the same variable at the same time?"

That problem is called a **race condition**.

---

# 1. The problem in this program

Global variable:

```go
var x = 0
```

Think:

```text
x = 0
```

You create 1000 goroutines:

```go
for i := 0; i < 1000; i++ {
    go increment(&w, ch)
}
```

Meaning:

```
Goroutine 1  ---> x = x + 1
Goroutine 2  ---> x = x + 1
Goroutine 3  ---> x = x + 1
...
Goroutine 1000 ---> x = x + 1
```

Expected:

```
x = 1000
```

But there is a problem.

---

# 2. Why x = x + 1 is dangerous

Many beginners think:

```go
x = x + 1
```

is one operation.

Actually internally it is three steps:

```
Step 1:
Read x


Step 2:
Add 1


Step 3:
Write back x
```

Example:

Initial:

```
x = 0
```

Two goroutines:

```
Goroutine A              Goroutine B

Read x (0)               Read x (0)

Add 1                   Add 1

Write x=1               Write x=1
```

Final:

```
x = 1
```

But we expected:

```
x = 2
```

One update is lost.

That is a race condition.

---

# 3. What is the purpose of the channel here?

Look at:

```go
ch := make(chan bool, 1)
```

This is a **buffered channel**.

Capacity:

```
1
```

Meaning:

```
Only one goroutine can enter at a time.
```

It is being used like a **lock**.

---

# 4. This line

```go
ch <- true
```

means:

> "Acquire the lock."

Imagine:

```
Channel capacity = 1

Before:

[ empty ]
```

Goroutine comes:

```go
ch <- true
```

Now:

```
[ true ]
```

The channel is full.

Another goroutine tries:

```go
ch <- true
```

It must wait.

Why?

Because:

```
capacity = 1
```

Only one value can exist.

So:

```
Goroutine A
     |
     v
 [true]


Goroutine B
 waiting...
```

---

# 5. Critical section

This is the important part:

```go
ch <- true

x = x + 1

<- ch
```

This is the protected area.

Meaning:

```
Enter lock

      |
      v

Modify shared data

      |
      v

Release lock
```

Only one goroutine can execute:

```go
x = x + 1
```

at a time.

---

# 6. This line releases the lock

```go
<- ch
```

Means:

> Receive/remove the value from the channel.

Before:

```
channel:

[true]
```

After:

```
channel:

[empty]
```

Now another goroutine can enter.

---

# 7. Complete flow

Imagine 3 goroutines.

Initial:

```
x = 0

channel:
[]
```

---

### Goroutine 1

```go
ch <- true
```

Channel:

```
[true]
```

Now:

```go
x = x + 1
```

x:

```
1
```

Release:

```go
<-ch
```

Channel:

```
[]
```

---

### Goroutine 2

Enters.

x:

```
2
```

---

### Goroutine 3

Enters.

x:

```
3
```

Finally:

```
x = 1000
```

---

# 8. But this is actually a Mutex pattern

Go already provides this.

Instead of:

```go
ch := make(chan bool,1)
```

normally we write:

```go
var mutex sync.Mutex
```

Then:

```go
mutex.Lock()

x = x + 1

mutex.Unlock()
```

Much clearer.

---

# 9. Why call it Mutex?

Mutex means:

**Mutual Exclusion**

Meaning:

> Only one goroutine is allowed inside this section at a time.

Visual:

Without mutex:

```
G1
 |
 |----\
      \
G2 ----> x
      /
G3 ---/
```

Everyone touches x together.

---

With mutex:

```
G1 ---> x
 |
(wait)

G2 ---> x
 |
(wait)

G3 ---> x
```

One at a time.

---

# 10. Now understand the function syntax

```go
func increment(wg *sync.WaitGroup, ch chan bool)
```

Parameters:

| Name | Type                  |
| ---- | --------------------- |
| wg   | pointer to WaitGroup  |
| ch   | channel carrying bool |

Meaning:

"This function needs a WaitGroup and a channel lock."

---

Inside:

```go
ch <- true
```

Acquire lock.

---

```go
x = x + 1
```

Modify shared variable.

---

```go
<-ch
```

Release lock.

---

```go
wg.Done()
```

Tell WaitGroup:

"I finished."

---

# 11. Execution order in main

```go
var w sync.WaitGroup
```

Creates:

```
counter = 0
```

---

```go
ch := make(chan bool,1)
```

Creates:

```
lock channel
capacity = 1
```

---

Loop:

```go
for i:=0;i<1000;i++ {
```

Each iteration:

```go
w.Add(1)

go increment(&w,ch)
```

Counter:

```
0
 |
1000
```

1000 goroutines running.

---

Main:

```go
w.Wait()
```

Blocks.

Meaning:

"Wait until all 1000 goroutines call Done."

---

After all:

```go
fmt.Println("final value of x", x)
```

Expected:

```
final value of x 1000
```

---

# 12. Interview explanation

If asked:

> "Explain how you solved the race condition."

Say:

> "Multiple goroutines were updating the same shared variable x, which can cause inconsistent results because read-modify-write operations are not atomic. I used synchronization to ensure only one goroutine modifies x at a time. The channel with capacity one works like a binary semaphore, allowing one goroutine to enter the critical section while others wait. In production Go code, this is usually replaced with sync.Mutex using Lock and Unlock."

---

# The big picture you should remember

You now know three synchronization tools:

| Tool      | Purpose                      |
| --------- | ---------------------------- |
| Channel   | Send data between goroutines |
| WaitGroup | Wait until goroutines finish |
| Mutex     | Protect shared data          |

Think:

```
Need communication?
        |
        v
     Channel


Need waiting?
        |
        v
   WaitGroup


Need protect shared variable?
        |
        v
      Mutex
```

This race condition topic is the bridge between beginner Go concurrency and real backend-level Go.
