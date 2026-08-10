Go supports different switch patterns, and you're seeing all major types here.

✅ 1. Basic value switch
i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}

✔ How it works:

i = 2

The switch matches case 2

Prints: two

✔ Output:
Write 2 as two


This is exactly like switch(i) in other languages.

✅ 2. Multiple values in a single case
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}

✔ What this does:

Checks today’s weekday

If today is Saturday or Sunday → print "It's the weekend"

Otherwise → "It's a weekday"

✔ Feature:

A single case can check multiple values. Example:

case A, B, C:

✅ 3. Switch without a value (acts like if-else)
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}

✔ Key point:

This is a condition switch (no expression after switch).

It works like:

if t.Hour() < 12 {
    ...
} else {
    ...
}

✔ Output:

Based on your time, prints:

It's before noon


or

It's after noon

✅ 4. Type switch (very important in Go)
whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
}
whatAmI(true)
whatAmI(1)
whatAmI("hey")

✔ Explanation:

This is a type switch.

i.(type) asks: “What type is i?”

✔ Steps:
Call 1:
whatAmI(true)


→ true is a bool
→ prints: I'm a bool

Call 2:
whatAmI(1)


→ 1 is an int
→ prints: I'm an int

Call 3:
whatAmI("hey")


→ "hey" is a string
→ not bool, not int
→ matches default
→ prints: Don't know type string

✔ Why this is useful:

Allows checking runtime types of an interface{} value
— essential in APIs, JSON decoding, etc.

🟩 FINAL OUTPUT EXAMPLE (depending on time/day)
Write 2 as two
It's a weekday        (or weekend)
It's after noon       (or before noon)
I'm a bool
I'm an int
Don't know type string  




how real-world production code uses them.

🟦 1. Value Switch (Normal switch)
Code:
switch i {
case 1:
    ...
case 2:
    ...
}

🟢 How Go handles this internally

Go evaluates the expression only once (here: i)

Then compares it against each case using ==

Only the first matching case runs

Go automatically includes a break at the end (no fallthrough by default)

🟢 Why Go does this?

In C/C++, forgetting break causes bugs

Go avoids this by design → safer switch

🟢 When used in real projects

Convert numbers to string labels

Command handling (API codes, status codes)

Menu-like logic in CLIs

Mapping enums (iota) to behavior

🟦 2. Multi-value Case
Code:
case time.Saturday, time.Sunday:

🟢 How it works

You can list multiple matching values in one case.

This is equivalent to:

if day == Saturday || day == Sunday

🟢 Internally

Go does:

(day == Saturday) OR (day == Sunday)

🟢 Real-world usage

Categorizing HTTP status codes

Grouping error types

Matching multiple commands in CLI tools

Routing in state machines

Example:

switch status {
case 400, 404, 422:
    return "Client error"
case 500, 502, 503:
    return "Server error"
}

🟦 3. Switch Without an Expression (Condition Switch)
Code:
switch {
case t.Hour() < 12:
    ...
case t.Hour() < 17:
    ...
default:
    ...
}

🟢 How Go executes this

Go treats it like switch true

Meaning each case must be a boolean expression

It executes the first true condition

🟢 Equivalent to:
if t.Hour() < 12 {
} else if t.Hour() < 17 {
} else {
}

🟢 Benefit over if/else

Cleaner cascading conditions without repeating if, else if

🟢 Real-world usage

API validation

Handling ranges

Matching complex conditions cleanly

Routing based on request context

Example (real API validation):

switch {
case user == nil:
    return errors.New("not found")
case !user.IsActive:
    return errors.New("inactive")
case user.IsBanned:
    return errors.New("banned")
}


Much cleaner than writing 3 ifs.

🟦 4. Type Switch (Most powerful one)
Code:
switch t := i.(type) {
case int:
    ...
case string:
    ...
default:
    ...
}

🟢 How Go processes this internally

i is of type interface{}.

i.(type) asks the runtime:

“What is the actual underlying type inside this interface?”

The matched case binds the value to t with the right type.

So:

i := interface{}(1)


Under the hood:

i → (type=int, value=1)


Type switch allows extracting that exact type.

🟢 Why this is required?

Because an interface{} may hold ANY type.

🟢 Real-world usage (very common)
✔ JSON handling

After decoding JSON → values may be float64, string, bool.

✔ Error handling

Some APIs return different error types.

✔ Handling dynamic data

Maps of any

Configurations

External input

HTTP bodys

✔ Pattern matching in Go

Since Go doesn’t have Rust/Kotlin pattern matching, type switch is used.

Example from real backend code:

func handleValue(v any) {
    switch x := v.(type) {
    case string:
        saveString(x)
    case int:
        saveInt(x)
    case []byte:
        saveBlob(x)
    default:
        log.Printf("unknown type %T", x)
    }
}

🟩 BONUS: What happens if you use fallthrough?

Go does NOT fall through automatically.

But you can force it:

switch i {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
}


Output:

one
two


Rarely used — only when writing parsers or state machines.

🟩 BONUS: Why Go switch is more powerful than C, Java, Python

No fallthrough by default → safer

Cases can be expressions (not only constants)

Type switching

Multi-value cases

Switch without condition

Short variable declaration inside switch

⭐ Final Summary Table
Switch Type	Purpose	Used When
Value switch	simple matching	numbers, enums, commands
Multi-case	OR matching	grouping conditions
No expression	condition branching	replaces long if-else
Type switch	check variable type	JSON, dynamic data, interfaces