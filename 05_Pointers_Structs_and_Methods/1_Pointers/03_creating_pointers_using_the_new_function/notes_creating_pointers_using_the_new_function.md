1️⃣ What new does (one-line definition)

new(T) allocates memory for type T,
initializes it to zero value,
and returns a pointer to T (*T).

2️⃣ Your program (for reference)
func main() {
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
}

3️⃣ Syntax-level explanation (line by line)
🔹 Line 1
size := new(int)


Break this into parts:

new(int)

int → type

new(int) →

allocates memory for one int

sets it to zero value (0)

returns a pointer

So:

new(int)  →  *int

size :=

size becomes a variable of type *int

✔️ size does not store an int
✔️ size stores an address of an int

Memory state after this line
Memory:
0x414020 ──► 0   (int zero value)

size ──► 0x414020

🔹 Line 2
fmt.Printf(
  "Size value is %d, type is %T, address is %v\n",
  *size, size, size,
)


Let’s break each argument:

*size

* → dereference operator

Reads the value stored at the address

Value = 0

size

Prints the address

Example: 0x414020

%T

Prints type → *int

📌 Output so far:

Size value is 0, type is *int, address is 0x414020

🔹 Line 3
*size = 85

What this means

Go to the address stored in size

Replace the value there with 85

Memory after modification
Memory:
0x414020 ──► 85

size ──► 0x414020

🔹 Line 4
fmt.Println("New size value is", *size)


Dereferences pointer

Prints updated value

Output:

New size value is 85

4️⃣ Execution flow (step-by-step)
main()
 ├── size := new(int)
 │     └── allocate int → 0
 │     └── return pointer (*int)
 │
 ├── print *size (0), type (*int), address
 │
 ├── *size = 85
 │     └── modify value at address
 │
 └── print *size (85)

5️⃣ Compare new vs & (IMPORTANT)
Using &
b := 255
a := &b


Variable already exists

Pointer points to existing value

Using new
size := new(int)


No existing variable

Go allocates memory for you

Returns pointer to zero value

Side-by-side comparison
Method	Memory created?	Zero value?	Common use
&var	❌ No	❌ No	Reference existing variable
new(T)	✅ Yes	✅ Yes	Create pointer directly
6️⃣ VERY important clarification (common confusion)

❌ Thinking:

new(int) returns an int

✅ Reality:

new(int) returns *int

7️⃣ One-line mental model (remember forever)

new(T) → *allocate T, zero it, return T

8️⃣ When you SHOULD use new

When you want:

a pointer

with guaranteed zero value

without declaring a variable first

Example:

count := new(int)

9️⃣ When new is NOT commonly used

In real Go code:

Struct literals are preferred

p := &MyStruct{}


new is mainly used for:

simple types

teaching pointers

low-level allocations

🔟 Final summary (simple words)

The new function allocates memory for a value, initializes it to zero, and returns a pointer to it.
Dereferencing the pointer allows reading and modifying the allocated value.