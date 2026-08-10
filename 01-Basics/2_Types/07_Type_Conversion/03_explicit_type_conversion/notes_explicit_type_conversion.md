What does “explicit type conversion” mean?

👉 Explicit means: you must clearly tell Go what you want
👉 Go will never guess or convert types automatically

Your example (step by step)
i := 10
var j float64 = float64(i)

What is happening here?

i := 10

i is an int

Type inferred as int

var j float64

j is explicitly declared as float64

float64(i)

This is explicit type conversion

You are telling Go:

“Convert this int value into a float64 value”

Now assignment is allowed:

j = float64(i)

Why is this needed in Go?

Because Go is strongly typed and strict.

❌ Go does NOT do automatic conversions like:

var j float64 = i // ❌ NOT allowed

Even though:

10 → 10.0 looks obvious to humans

Go refuses to assume anything

What happens if you remove the conversion?
var j float64 = i

Compiler error:
cannot use i (variable of type int) as float64 value in assignment

💡 Go is basically saying:

“I see an int. You want a float64.
Tell me exactly how to convert it.”

What exactly is float64(i)?

It is not a function call.
It is a type conversion expression.

General syntax:
TargetType(value)

Examples:
float64(i)
int(b)
string(bytes)

What changes after conversion?
i := 10 // int
j := float64(i) // float64

Variable Type Value
i int 10
j float64 10.0
Why Go designers chose this rule (important)
Imagine this silent bug 👇
price := 99 // int
discount := 0.15 // float64

final := price - discount // what should happen?

Different languages behave differently → bugs

Go avoids this by saying:

“Be explicit. No surprises.”

Real-world backend example
items := 3
avgPrice := 199.99

total := float64(items) \* avgPrice

✔ Safe
✔ Clear
✔ No hidden conversion
✔ Production-grade code

One-line mental rule 🧠

If types are different → you MUST convert explicitly

Summary (keep this)

Go never auto-converts types

You must use:

float64(i)
int(f)

This avoids hidden bugs

This is why Go is trusted in backend systems
