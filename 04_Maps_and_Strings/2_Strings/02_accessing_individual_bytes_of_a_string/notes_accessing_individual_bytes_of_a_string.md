1️⃣ Overall structure (what goes where)
package main        // 1️⃣ Program entry package

import (
	"fmt"           // 2️⃣ Import fmt package
)

What this means

package main → tells Go: this is an executable program

import "fmt" → we need fmt.Printf for printing

2️⃣ Function definition: printBytes
func printBytes(s string) {

Syntax breakdown

func → keyword to define a function

printBytes → function name

(s string) → parameter

s → variable name

string → type

{ → function body starts

Inside the function
fmt.Printf("Bytes: ")


Calls Printf from fmt

Prints the text "Bytes: "

No newline here

for i := 0; i < len(s); i++ {

for loop syntax

i := 0 → declare & initialize i

i < len(s) → loop condition

i++ → increment after each loop

len(s) → number of bytes in string s

fmt.Printf("%x ", s[i])

What happens here

s[i] → accesses i-th byte of the string

Type of s[i] → byte (uint8)

%x → format specifier for hexadecimal

Prints each byte in hex form

}


Ends for loop

}


Ends printBytes function

3️⃣ main() function (program execution starts here)
func main() {


Every Go executable must have main()

Program starts executing from here

Step 1: Variable declaration
name := "Hello World"


:= → short variable declaration

name → variable

"Hello World" → string literal

Stored internally as bytes

Step 2: Print the string
fmt.Printf("String: %s\n", name)

Syntax meaning

"String: %s\n" → format string

%s → placeholder for string

\n → newline

name → value replacing %s

Step 3: Function call
printBytes(name)


Calls printBytes

Passes name into parameter s

4️⃣ Execution flow (VERY IMPORTANT)

This is the exact runtime flow:

main()
 ├── name := "Hello World"
 ├── fmt.Printf("String: %s\n", name)
 │      ↓
 │   prints: String: Hello World
 │
 ├── printBytes(name)
 │      ↓
 │   fmt.Printf("Bytes: ")
 │
 │   for i := 0; i < len(s); i++ {
 │        fmt.Printf("%x ", s[i])
 │   }
 │
 └── program ends

5️⃣ What each syntax piece did (one-line each)
Syntax	What it did
package main	Marks executable program
import "fmt"	Enables printing
func printBytes(s string)	Defines a function taking a string
len(s)	Counts bytes
s[i]	Gets one byte
%x	Prints hex
%s	Prints string
:=	Declare + assign
printBytes(name)	Calls function
6️⃣ Why they wrote it this way (logic)

main() → handles program start

printBytes() → separate responsibility

Loop → process string byte-by-byte

Formatting → shows internal representation

This is clean separation of logic.

7️⃣ Final mental model (remember this)

main() creates data
→ passes data to a function
→ function processes it using a loop
→ prints result

=========================================================================

1️⃣ Key idea (MOST IMPORTANT)

In Go:

A string is a read-only slice of bytes ([]byte)

So:

Each character you see in "Hello World"

Is stored internally as one or more bytes (UTF-8 encoding)

That’s why Go allows:

s[i]


→ it gives one byte, not a character.

2️⃣ Program flow (line by line, in order)
Step 1: main() starts
name := "Hello World"


A string literal is created

Internally stored as UTF-8 encoded bytes

Step 2: Print the string
fmt.Printf("String: %s\n", name)


%s tells Go: “treat this as a string”

Output:

String: Hello World

Step 3: Call printBytes(name)
printBytes(name)


The string is passed to the function as s string

Step 4: Inside printBytes
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}


Here’s what happens exactly:

len(s)

Returns number of bytes, not characters

"Hello World" → 11 bytes

s[i]

Accesses the i-th byte

Type: byte (uint8)

%x

Prints the byte in hexadecimal

3️⃣ What bytes are actually printed?

Let’s break "Hello World":

Character	ASCII / UTF-8 byte	Hex
H	72	48
e	101	65
l	108	6c
l	108	6c
o	111	6f
(space)	32	20
W	87	57
o	111	6f
r	114	72
l	108	6c
d	100	64

So output becomes:

Bytes: 48 65 6c 6c 6f 20 57 6f 72 6c 64


✔️ These are UTF-8 encoded byte values
✔️ For English letters, UTF-8 = ASCII (1 byte each)

4️⃣ VERY IMPORTANT: Why this can be dangerous 🚨

This works only because "Hello World" contains ASCII characters.

What if the string has non-English characters?

Example:

s := "Hi 👋"


👋 is 4 bytes in UTF-8

len(s) counts bytes, not characters

s[i] may print partial bytes of one character

That leads to broken output if you think in “characters”.

5️⃣ Correct way to iterate characters (Unicode-safe)
Use range
for _, r := range s {
    fmt.Printf("%c ", r)
}


range decodes UTF-8

r is a rune (Unicode code point)

Safe for emojis, Indian languages, etc.

6️⃣ One-line mental model (INTERVIEW GOLD)

🔹 string[i] → byte
🔹 len(string) → bytes count
🔹 range string → characters (runes)

7️⃣ Why Go is designed this way (WHY question)

Go is systems-friendly

Strings are byte-based for:

Performance

Network protocols

File I/O

Unicode decoding is explicit, not hidden

8️⃣ Short summary (as you asked earlier)

We first store the string.
We print it normally using %s.
Then we loop through its bytes using len() and indexing.
Each index gives one UTF-8 byte, which we print in hex using %x.

======================================================================
**********************************************************************
<><<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>><>
**********************************************************************
======================================================================

Good question. This is one of the most important Go concepts: **how function parameters receive values**.

Let's look at your code:

```go
func printBytes(s string) {

    fmt.Printf("Bytes: ")

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}
```

Here:

```go
func printBytes(s string)
```

means:

> "I am creating a function called `printBytes` which expects one input value. That input must be a string. Inside this function, I will call that input `s`."

Think of `s` as a temporary variable.

---

Now go to `main()`:

```go
func main() {

    name := "Keerthi Shankar S"

    fmt.Printf("String : %s\n", name)

    printBytes(name)
}
```

This line calls the function:

```go
printBytes(name)
```

At this moment, Go takes the value inside `name` and passes it into the function.

Before function call:

```
main()

name
 |
 |
 v
"Keerthi Shankar S"
```

Then:

```go
printBytes(name)
```

happens:

```
printBytes()

s
 |
 |
 v
"Keerthi Shankar S"
```

Now inside the function:

```go
fmt.Println(s)
```

is actually printing:

```
Keerthi Shankar S
```

---

## Visual flow

```
             main()

        name := "Hello"


              |
              |
              v


       printBytes(name)


              |
              |
              v


   func printBytes(s string)


              |
              |
              v


       s = "Hello"

```

---

## Important: The names do NOT need to match

You could write:

```go
func printBytes(x string) {

    fmt.Println(x)

}
```

and call:

```go
name := "Keerthi"

printBytes(name)
```

It still works.

Because:

```
name
 |
 | value copied
 v

x
```

The parameter name is only local to the function.

---

## Another example

```go
func add(a int, b int) {

    fmt.Println(a+b)

}


func main(){

    x := 10
    y := 20

    add(x,y)

}
```

Flow:

```
main()

x = 10
y = 20


add(x,y)


inside function:

a = 10
b = 20


output:

30

```

---

## What happens internally?

When you call:

```go
printBytes(name)
```

Go roughly does:

```go
s := name
```

So:

```
name ----------------+
                     |
                     v
              s receives value

```

For strings, Go copies the string header (not the entire underlying data).

---

## In your specific code:

```go
name := "Keerthi Shankar S"
```

Memory:

```
name
 |
 |
 v
String data:
K e e r t h i ...
```

Call:

```go
printBytes(name)
```

creates:

```
s
 |
 |
 v
same string value
```

Then:

```go
s[i]
```

accesses each byte:

```
K = 0x4b
e = 0x65
e = 0x65
r = 0x72
...
```

That is why output:

```
Bytes:
6b 65 65 72 74 68 69
```

---

### Remember this rule:

```go
func function(parameter datatype)
```

means:

> "When someone calls this function, give me a value of this datatype, and I will refer to it using this parameter name."

Example:

```go
func login(userID int)
```

Call:

```go
login(100)
```

Inside:

```
userID = 100
```

Same concept is used everywhere in your backend:

```go
func CreateBooking(ctx context.Context, req CreateBookingRequest)
```

`req` comes from wherever the function is called and becomes available inside the function.
