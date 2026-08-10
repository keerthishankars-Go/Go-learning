Great question. This line:

dchnl <- digit

is the **heart of channel communication**. Let's understand the syntax formation and how to say it.

---

## 1. First understand the pieces

Your function:

func digits(number int, dchnl chan int) {

means:

text
Function name: digits

Inputs:

number → int

dchnl → channel that carries int values

So inside the function:

text
dchnl

is a pipe

that can carry integers

---

## 2. The `<-` operator meaning

In , the arrow operator is the **channel operator**.

It has two directions.

---

### Sending

channel <- value

Read as:

> "Send value into the channel"

Example:

dchnl <- digit

Say it as:

> "Send digit to dchnl"

or

> "Write digit into the dchnl channel"

or

> "Push digit onto the channel"

All are acceptable.

---

### Receiving

value := <-channel

Read as:

> "Receive a value from the channel"

Example:

digit := <-dchnl

Say:

> "Receive from dchnl and store it in digit"

---

# 3. Now your exact code

func digits(number int, dchnl chan int) {

    for number != 0 {

        digit := number % 10

        dchnl <- digit

        number /= 10
    }

    close(dchnl)

}

Let's execute.

Input:

number = 589

---

## First loop

digit := number % 10

Calculation:

589 % 10 = 9

Now:

digit = 9

Memory:

number
589

digit
9

dchnl
(empty channel)

---

Now:

dchnl <- digit

Means:

Send 9 into dchnl

Visual:

digits routine

digit = 9

      |
      |
      v

+---------------+
| dchnl |
| |
| receives 9 |
+---------------+

---

## Second loop

Remove last digit:

number /= 10

Now:

589 / 10 = 58

Again:

digit := number % 10

58 % 10 = 8

Then:

dchnl <- digit

means:

Send 8 into dchnl

---

Third loop:

5 % 10 = 5

Send:

dchnl <- 5

---

# 4. What does the receiving side look like?

Somewhere else you will have:

for digit := range dchnl {
fmt.Println(digit)
}

This means:

> Keep receiving values from dchnl until the channel closes.

Flow:

digits routine main routine

digit = 9
|
|
v
dchnl -------------> receive 9

digit = 8
|
|
v
dchnl -------------> receive 8

digit = 5
|
|
v
dchnl -------------> receive 5

Output:

9
8
5

---

# 5. Why close(dchnl)?

At the end:

close(dchnl)

means:

> "I am done sending values. No more data will come."

Important:

Closing does NOT destroy the channel.

It only sends an EOF-like signal.

Example:

Before close:

dchnl

9
8
5

After close:

dchnl

9
8
5

CLOSED

Receiver knows:

No more digits are coming.
Stop the loop.

---

# 6. Very important syntax pattern

Remember this:

## Send

channel <- data

Read:

data es INTO channel

Visual:

data ---> channel

Example:

jobs <- task

means:

Send task into jobs channel

---

## Receive

data := <-channel

Visual:

channel ---> data

Example:

task := <-jobs

means:

Receive task from jobs channel

---

# 7. Why did they name it `dchnl`?

They could have written:

func digits(number int, x chan int)

and:

x <- digit

It would work.

Because parameter name is only local.

This:

dchnl <- digit

is just:

variable dchnl (which happens to be a channel)

- send operator <-
- value digit

  ***

# One sentence to remember

Whenever you see:

A <- B

ask:

> "Is A a channel?"

If yes:

Send B into channel A

So:

dchnl <- digit

means:

> **Send the current digit value into the integer channel named dchnl.**


============================================================================

Let's trace this program. This is a very good example because now you are combining:

* goroutines
* channels
* channel closing
* `range` over channel
* multiple goroutines

First let's calculate the **final output**, then understand the flow.

---

## Final output

Input:

go
number := 589


### calcSquares

It calls:

go
go digits(number, dch)


The `digits()` function sends:


589 → 9 → 8 → 5


through the channel.

So:

go
for digit := range dch {
    sum += digit * digit
}


calculates:


9² + 8² + 5²

81 + 64 + 25

= 170


So:


squares = 170


---

### calcCubes

Same digits:


9, 8, 5


Calculates:


9³ + 8³ + 5³

729 + 512 + 125

= 1366


So:


cubes = 1366


---

### Main:

go
squares + cubes


becomes:


170 + 1366

= 1536


## Output:


Final output 1536


---

# Now understand the execution flow

Your program actually creates **5 goroutines**.

Let's count.

---

## 1. Main goroutine starts

go
func main()


Creates:

go
sqrch := make(chan int)
cubech := make(chan int)


Memory:


main

sqrch  ---> channel<int>

cubech ---> channel<int>


---

## 2. Start calcSquares

go
go calcSquares(number, sqrch)


New goroutine:


Main

 |
 |
 +------ calcSquares


---

## 3. Start calcCubes

go
go calcCubes(number, cubech)


Now:


Main

 |
 +------ calcSquares
 |
 +------ calcCubes


---

# Inside calcSquares

Execution:

go
func calcSquares(number int, squareop chan int)


Creates another channel:

go
dch := make(chan int)


Now:


calcSquares

dch ---> channel<int>


Then:

go
go digits(number, dch)


Creates another goroutine:


calcSquares

 |
 |
 +------ digits()


Now digits starts sending.

---

# Inside digits()

First loop:

go
digit := number % 10


For:


589


gets:


9


Then:

go
dchnl <- digit


means:


Send 9 into dch


Flow:


digits goroutine

9
 |
 |
 v

dch channel

 |
 |
 v

calcSquares


---

Now this line in calcSquares:

go
for digit := range dch


means:

> Keep receiving values from dch until it is closed.

It receives:


9


and does:


sum = 0 + 9*9

sum = 81


---

Next digits sends:


8


calcSquares receives:


sum = 81 + 64

sum = 145


Next:


5


calcSquares:


sum = 145 + 25

sum = 170


---

Then digits finishes:

go
close(dchnl)


This tells:


calcSquares:

"No more digits are coming"


So:

go
for digit := range dch


stops.

Then:

go
squareop <- sum


sends:


170


to main.

---

# calcCubes is happening at the same time

It does the same thing:


digits()

9
8
5


but calculation:


9³ = 729

8³ = 512

5³ = 125


Total:


1366


Then:

go
cubeop <- sum


sends:


1366


to main.

---

# Main waits here

go
squares, cubes := <-sqrch, <-cubech


This line is very important.

It means:

Equivalent to:

go
squares := <-sqrch

cubes := <-cubech


Meaning:


Wait for square result

then

wait for cube result


Main is blocked until both values arrive.

---

# Complete picture


                 MAIN

        sqrch          cubech
          |              |
          |              |
          v              v

   calcSquares      calcCubes
        |               |
        |               |
       dch             dch
        |               |
        |               |
        v               v

     digits()        digits()

        |               |
        |               |
     9,8,5           9,8,5

        |               |
        v               v

      170             1366

        \             /
         \           /

          170 + 1366

               |

              1536


---

# One important new syntax here

This line:

go
for digit := range dch


is a special Go channel pattern.

It means:

go
receive values continuously from channel


until:

go
close(dch)


happens.

Equivalent mental model:

go
for {

    digit, ok := <-dch

    if !ok {
        break
    }

    // use digit
}


`range` hides this complexity.

---

# Why create another goroutine inside calcSquares?

This part:

go
dch := make(chan int)

go digits(number, dch)


creates a pipeline:


number
  |
  |
digits goroutine
  |
  |
channel
  |
  |
calculator


This is a common Go pattern:

**Producer → Channel → Consumer**

Examples in real backend systems:


Database reader
      |
      |
    channel
      |
      |
 Worker processors


or:


API requests
      |
      |
 Queue channel
      |
      |
 Workers


This program is basically a small version of a real Go worker pipeline.
