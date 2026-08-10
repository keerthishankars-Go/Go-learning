This is one of the best examples to understand **Go's syntax, interfaces, method sets, and runtime execution**. Let's not just explain **what** happens, but **how the Go compiler and runtime think**.

---

# Phase 1: The compiler reads your file (Top to Bottom)

The compiler first scans all the type and function declarations.

## Step 1: Package

```go
package main
```

Meaning:

> This is an executable program.

---

## Step 2: Import

```go
import "fmt"
```

Compiler loads the `fmt` package.

Memory (conceptually):

```text
Program

|
+-- fmt package
      |
      +-- Println()
      +-- Printf()
      +-- Sprintf()
```

---

# Phase 2: Interface declarations

First interface:

```go
type SalaryCalculator interface {
    DisplaySalary()
}
```

This creates a **new interface type**.

Think of it like this:

```text
SalaryCalculator

Required Methods

+--------------------+
| DisplaySalary()    |
+--------------------+
```

Nothing is executed.

No objects are created.

Just a **type definition**.

---

Second interface:

```go
type LeaveCalculator interface {
    CalculateLeavesLeft() int
}
```

Compiler stores:

```text
LeaveCalculator

Required Methods

+------------------------------+
| CalculateLeavesLeft() int    |
+------------------------------+
```

Again...

Nothing runs.

---

# Phase 3: Struct declaration

```go
type Employee struct {
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}
```

Compiler creates a new concrete type.

Think:

```text
Employee

Fields

firstName
lastName
basicPay
pf
totalLeaves
leavesTaken
```

Still...

No Employee exists yet.

---

# Phase 4: Method declarations

Now:

```go
func (e Employee) DisplaySalary()
```

Let's understand this syntax deeply.

---

## The syntax

```go
func
```

means

> I'm defining a function.

---

Then

```go
(e Employee)
```

This is called the **receiver**.

Read it like English.

> This function belongs to Employee.

Equivalent mentally to:

```text
Employee

Methods

DisplaySalary()
```

Not

```text
Global Functions

DisplaySalary()
```

This is why later you write:

```go
e.DisplaySalary()
```

instead of

```go
DisplaySalary(e)
```

---

Same here:

```go
func (e Employee) CalculateLeavesLeft() int
```

Compiler adds another method.

Now Employee looks like:

```text
Employee

Fields

firstName
lastName
basicPay
pf
totalLeaves
leavesTaken

-----------------------

Methods

DisplaySalary()

CalculateLeavesLeft()
```

---

# Phase 5: Compiler checks interfaces

Now compiler asks:

Does Employee satisfy SalaryCalculator?

SalaryCalculator says:

```text
Need:

DisplaySalary()
```

Employee has:

```text
DisplaySalary()
```

YES

---

Compiler asks again.

Does Employee satisfy LeaveCalculator?

Need:

```text
CalculateLeavesLeft() int
```

Employee has:

```text
CalculateLeavesLeft() int
```

YES

---

So compiler records

```text
Employee

implements

SalaryCalculator

LeaveCalculator
```

Automatically.

You never write:

```java
implements
```

like Java.

---

# Phase 6: main() starts

Now runtime begins.

Execution starts here.

```go
func main() {
```

Nothing above was executed.

Everything above was definitions.

---

# Step 1

```go
e := Employee{
```

Memory:

```text
Stack

e

+---------------------------+
| firstName = Naveen        |
| lastName = Ramanathan     |
| basicPay = 10000          |
| pf = 200                  |
| totalLeaves = 30          |
| leavesTaken = 5           |
+---------------------------+
```

Now a real Employee exists.

---

# Step 2

Now this line:

```go
var s SalaryCalculator = e
```

Let's break the syntax.

General Go declaration syntax is

```go
var variable Type = value
```

Examples

```go
var age int = 20
```

```go
var name string = "Keerthi"
```

```go
var e Employee = Employee{}
```

Exactly the same idea.

Now

```go
var s SalaryCalculator = e
```

means

> Create a variable named **s** whose type is **SalaryCalculator**.

Initially

```text
s

SalaryCalculator Interface

(empty)
```

Then

```go
= e
```

Go asks

Does Employee satisfy SalaryCalculator?

YES.

So Go creates an interface value.

Internally

```text
s

+--------------------------------+
| Dynamic Type = Employee        |
| Dynamic Value = Employee{...}  |
+--------------------------------+
```

This is the famous

```text
(Type, Value)
```

pair.

---

# Step 3

Now

```go
s.DisplaySalary()
```

Compiler looks at

```text
What is s?

↓

SalaryCalculator
```

SalaryCalculator guarantees

```text
DisplaySalary()
```

Therefore

```go
s.DisplaySalary()
```

is valid.

Runtime then does

```text
Interface

↓

Dynamic Type = Employee

↓

Call Employee.DisplaySalary()

↓

Pass stored Employee value

↓

Execute
```

Equivalent to

```go
Employee.DisplaySalary(e)
```

behind the scenes.

---

# Step 4

Next

```go
var l LeaveCalculator = e
```

Exactly the same process.

Go asks

Does Employee satisfy LeaveCalculator?

Need

```text
CalculateLeavesLeft() int
```

Employee has

```text
CalculateLeavesLeft() int
```

YES.

Interface created.

```text
l

+--------------------------------+
| Dynamic Type = Employee        |
| Dynamic Value = Employee{...}  |
+--------------------------------+
```

---

# Step 5

Now

```go
l.CalculateLeavesLeft()
```

Compiler thinks

```text
What is l?

↓

LeaveCalculator

↓

What methods exist?

↓

CalculateLeavesLeft()

↓

Allowed
```

Runtime

```text
Open interface

↓

Find Employee

↓

Execute

Employee.CalculateLeavesLeft()

↓

Returns 25
```

Then

```go
fmt.Println(...)
```

prints

```text
Leaves left = 25
```

---

# Understanding the Receiver Syntax

Many beginners find this strange:

```go
func (e Employee) DisplaySalary()
```

Think of it like attaching functions to a type.

Without methods:

```go
func DisplaySalary(e Employee)
```

You call:

```go
DisplaySalary(e)
```

With methods:

```go
func (e Employee) DisplaySalary()
```

You call:

```go
e.DisplaySalary()
```

The compiler internally treats it similarly to:

```text
DisplaySalary(e)
```

The method receiver is simply syntactic sugar that associates the function with the `Employee` type.

---

# Why can `s.DisplaySalary()` work but `s.CalculateLeavesLeft()` cannot?

Because the **variable's type** is the interface.

```go
var s SalaryCalculator = e
```

The compiler only sees:

```text
s

Type = SalaryCalculator
```

SalaryCalculator only promises:

```text
DisplaySalary()
```

So

```go
s.DisplaySalary()
```

✅ Allowed

But

```go
s.CalculateLeavesLeft()
```

❌ Compile Error

Even though the actual object inside is an `Employee`, the compiler only allows methods guaranteed by the interface's method set.

---

# Complete Runtime Flow

```text
Program Starts
      │
      ▼
Compiler reads interfaces
      │
      ▼
Compiler reads Employee
      │
      ▼
Compiler attaches methods
      │
      ▼
Compiler verifies Employee satisfies both interfaces
      │
      ▼
main()
      │
      ▼
Create Employee object (e)
      │
      ▼
Assign e to SalaryCalculator (s)
      │
      ▼
Interface stores:
(Type = Employee, Value = e)
      │
      ▼
s.DisplaySalary()
      │
      ▼
Runtime dispatches to Employee.DisplaySalary()
      │
      ▼
Assign e to LeaveCalculator (l)
      │
      ▼
Interface stores:
(Type = Employee, Value = e)
      │
      ▼
l.CalculateLeavesLeft()
      │
      ▼
Runtime dispatches to Employee.CalculateLeavesLeft()
      │
      ▼
Returns 25
      │
      ▼
Print:
Leaves left = 25
```

---

## A mental model that will help you write this syntax yourself

Whenever you see a line like:

```go
variable.Method()
```

train yourself to ask two questions:

1. **What is the static type of `variable`?** (`Employee`, `SalaryCalculator`, `LeaveCalculator`, etc.)
2. **Does that type's method set include `Method()`?**

If the answer is yes, the call is valid.

That single habit explains not only this example, but nearly every method call you'll write in Go—from interfaces in your backend services to `http.Client.Do()`, `db.Query()`, `gin.Context.JSON()`, and the standard library's `io.Reader`/`io.Writer` interfaces.
