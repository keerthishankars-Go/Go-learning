🔢 The code (for reference)
func change(s ...string) {
	s[0] = "Go"
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

1️⃣ First: what s ...string REALLY means
func change(s ...string)


➡️ Inside the function:

s is of type []string


So mentally rewrite it as:

func change(s []string)


⚠️ Important:
... does NOT mean “copy values”.
It only controls how arguments are passed.

2️⃣ Memory picture BEFORE function call

In main:

welcome := []string{"hello", "world"}


Memory looks like this:

Underlying array:
+---------+--------+
| "hello" | "world"|
+---------+--------+
    ↑
  welcome slice


The slice welcome contains:

pointer → underlying array

length = 2

capacity = 2

3️⃣ What change(welcome...) does (THIS IS THE KEY)
... is called slice expansion
change(welcome...)


Means:

“Take the elements of welcome and pass them as variadic arguments”

But here is the compiler trick 👇

Since welcome is already a slice of string,

Go does NOT create a new slice.

Instead, it passes the same slice header.

So internally this happens:

s := welcome   // SAME slice header, SAME array


Memory now:

+---------+--------+
| "hello" | "world"|
+---------+--------+
    ↑        ↑
 welcome     s


👉 Both welcome and s point to the same underlying array

4️⃣ Inside change() – step-by-step
func change(s ...string) {
	s[0] = "Go"
}

What executes?
s[0] = "Go"


This modifies:

Underlying array[0]


Memory becomes:

+-------+--------+
| "Go"  | "world"|
+-------+--------+
    ↑        ↑
 welcome     s


⚠️ Since slices share memory, this change affects everyone

5️⃣ Back to main
fmt.Println(welcome)


welcome still points to the same array, which is now:

["Go", "world"]

✅ Final output
[Go world]

6️⃣ Why no new slice was created?

Because of this rule 👇

🔑 Rule of variadic + slice expansion
How you call	What happens
change("a","b")	Compiler creates new slice
change(slice...)	SAME slice is passed
change(append(slice,"x")...)	New slice
7️⃣ What if you DIDN’T use ...?

This is invalid ❌:

change(welcome) // compile-time error


Because:

change expects individual strings, not a slice

8️⃣ Why this is dangerous (and powerful)
⚠️ Dangerous

Because:

You might accidentally mutate caller’s data

✅ Powerful

Because:

Very efficient (no copying)

Used heavily in:

append

fmt.Println

logging libraries

9️⃣ How to make it SAFE (copy before modifying)

If you want to avoid modifying original slice:

func change(s ...string) {
	tmp := make([]string, len(s))
	copy(tmp, s)
	tmp[0] = "Go"
}


Now:

welcome stays unchanged

🔟 One-line mental model (remember this forever)

slice... passes the same slice, not a copy — so modifications affect the original.

🧠 Interview-ready explanation

“When a slice is passed to a variadic function using ..., the slice header is reused and no new slice is created. Since slices reference an underlying array, modifying the slice inside the function modifies the original data.”