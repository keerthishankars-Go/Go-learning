Good. This is another **very important production Go pattern**: `defer`.

You already learned:

* goroutine → run concurrently
* channel → communicate
* WaitGroup → wait
* Mutex → protect

Now:

> **defer → schedule something to run automatically when the current function finishes**

This example is about **measuring execution time**, but the same pattern is used everywhere in backend code.

---

# First understand the problem

Suppose you want to know:

> "How much time did this function take?"

You need:

1. Start timer
2. Execute function logic
3. Calculate difference

Normally:


func test() {

    start := time.Now()

    time.Sleep(2*time.Second)

    elapsed := time.Since(start)

    fmt.Println(elapsed)
}


Flow:


Start timer

      |
      v

Do work

      |
      v

Calculate time


But what if the function has many return paths?

Example:


func test() {

    start := time.Now()

    if error {
        return
    }

    if anotherError {
        return
    }

}


You would need to write timer calculation everywhere.

Go gives:


defer


---

# Your code


func test() {

    start := time.Now()

    defer totalTime(start)

    time.Sleep(2 * time.Second)

    fmt.Println("Sleep complete")
}


Let's break execution.

---

# Step 1

main:


func main() {
    test()
}


calls:


test()


Execution enters test.

---

# Step 2

Inside test:


start := time.Now()


Example:


start

10:00:00


Variable:


start
 |
 v
time.Time object


---

# Step 3 (important)


defer totalTime(start)


Many beginners misunderstand this.

It does NOT execute now.

It means:

> "Remember this function call. Execute it when test() is about to finish."

Timeline:


test starts

 |
 |
defer registered
 |
 |
do remaining work
 |
 |
function ending
 |
 |
execute deferred function


---

# Step 4

Now:


time.Sleep(2 * time.Second)


Program waits.

After 2 seconds:


fmt.Println("Sleep complete")


prints:


Sleep complete


---

# Step 5

Now test() is ending.

Before leaving:

Go checks:

"Are there deferred functions?"

Yes:


defer totalTime(start)


Execute it.

---

# Step 6

Now this runs:


func totalTime(start time.Time)


Inside:


time.Since(start)


means:


current time - start time


Example:


Current:

10:00:02.001


Start:

10:00:00.001


Difference:

2 seconds


Then:


.Seconds()


converts:


2s

to

2.000000


---

Output:


Sleep complete
Total time taken 2.001234 seconds


---

# Important syntax understanding

This line:


defer totalTime(start)


is a function call with a keyword before it.

Normally:


totalTime(start)


means:

> execute now

With defer:


defer totalTime(start)


means:

> execute later when this function returns

---

# Very important: arguments are evaluated immediately

Example:


x := 10

defer fmt.Println(x)

x = 20


What prints?

Many beginners think:


20


But output:


10


Why?

Because:


defer fmt.Println(x)


captures the current value.

At defer time:


x = 10


Go remembers:


Print 10 later


---

# Real backend examples

You will see this everywhere.

---

## 1. Closing files

Without defer:


file.Close()


You might forget.

Better:


file, err := os.Open("data.txt")

defer file.Close()


Meaning:

> When this function finishes, close the file.

---

## 2. Database connection

Example:


conn, err := db.Begin()

defer conn.Close()


---

## 3. HTTP response body

Very common:


resp, err := client.Do(req)

defer resp.Body.Close()


Meaning:

> After reading response, clean resources.

---

## 4. Mutex

You learned race conditions.

Normally:


mutex.Lock()

// critical section

mutex.Unlock()


Safer:


mutex.Lock()

defer mutex.Unlock()

// critical section


Why?

Even if an error occurs:


return


the unlock still happens.

---

# Defer execution order

If you have:


func test(){

    defer fmt.Println("first")

    defer fmt.Println("second")

    defer fmt.Println("third")

}


Output:


third
second
first


Why?

Because defer uses:

**LIFO**

Last In First Out.

Like stack.


defer first
      |
defer second
      |
defer third


Execute:

third
second
first


---

# How to recognize when to use defer

Ask:

> "Do I have something that must happen when this function ends?"

Examples:

| Resource      | Defer                    |
| ------------- | ------------------------ |
| File          | Close                    |
| DB connection | Close                    |
| HTTP response | Body.Close               |
| Mutex         | Unlock                   |
| Timer         | Stop                     |
| Transaction   | Rollback/Commit handling |

---

# Interview explanation

If asked:

> "What is defer in Go?"

Say:

> "`defer` schedules a function call to execute when the surrounding function returns. It is commonly used for cleanup operations like closing files, releasing locks, closing database connections, and measuring execution time. Deferred calls execute in LIFO order."

---

# Connect this with your payment client code

Remember your previous code:


resp, err := c.httpClient.Do(httpReq)

defer resp.Body.Close()


Now it makes sense.

The developer is saying:

> "I don't care where this function exits later. Success, error, validation failure — before leaving, always clean up the response body."

That is the production mindset behind `defer`.
