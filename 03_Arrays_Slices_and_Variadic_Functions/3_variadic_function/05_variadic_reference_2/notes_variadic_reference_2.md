This example is very important, because it shows both behaviors together:

Modification affects caller

append does NOT affect caller

Let’s walk through this slowly, step-by-step, with exact flow and memory changes.

🔢 The code (reference)
func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	welcome := []string{"Hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

🧠 ONE-LINE SUMMARY (READ FIRST)

Direct modification affects the original slice.
append may create a new array, so its changes may NOT affect the original slice.

Now let’s prove this.

1️⃣ Step 1: Slice creation in main
welcome := []string{"Hello", "world"}


Memory:

Underlying array (A):
+--------+--------+
| Hello  | world  |
+--------+--------+

welcome slice:
ptr → array A
len = 2
cap = 2

2️⃣ Step 2: Function call with ...
change(welcome...)

What ... does here

No new slice is created

Same slice header is reused

So inside change:

s slice:
ptr → array A
len = 2
cap = 2


Memory now:

array A:
+--------+--------+
| Hello  | world  |
+--------+--------+
   ↑        ↑
welcome     s

3️⃣ Step 3: Modify element directly
s[0] = "Go"


This modifies array A directly.

Memory becomes:

array A:
+----+--------+
| Go | world  |
+----+--------+


Because:

s and welcome share the same array

4️⃣ Step 4: The IMPORTANT line — append
s = append(s, "playground")


Let’s slow this down.

Before append:
len = 2
cap = 2


Slice is full.

🔥 What append does internally

Checks capacity → ❌ no space

Allocates a new array

Copies existing elements

Appends new element

Returns a new slice

So Go creates:

New array (B):
+----+--------+-------------+
| Go | world  | playground  |
+----+--------+-------------+

s slice NOW:
ptr → array B
len = 3
cap = 4 (or more, implementation-dependent)


⚠️ IMPORTANT:

welcome still points to array A

Only s now points to array B

5️⃣ Step 5: Print inside function
fmt.Println(s)


Prints:

[Go world playground]


This is from array B.

6️⃣ Step 6: Function returns

Local variable s is destroyed

Array B may be garbage-collected later

welcome is untouched by append

7️⃣ Step 7: Back in main
fmt.Println(welcome)


What does welcome point to?

array A:
+----+--------+
| Go | world  |
+----+--------+


So output is:

[Go world]

🧠 FINAL OUTPUT (COMPLETE)
[Go world playground]
[Go world]

🔑 WHY THIS HAPPENS (KEY RULES)
Rule 1

Modifying elements of a slice modifies the underlying array

s[0] = "Go"  // affects caller

Rule 2

append may create a new array

s = append(s, "playground") // affects only local s

Rule 3

You must reassign append

s = append(s, ...)


Because it may return a new slice.

🧠 Mental model (lock this in)
Before append:
welcome ──┐
          ├──► array A
s ────────┘

After append:
welcome ──► array A
s ────────► array B

⚠️ This is why bugs happen in real Go code

Developers assume:

append modifies original slice ❌

Reality:

append modifies only the returned slice

✅ How to make behavior predictable
Option 1: Return the slice
func change(s ...string) []string {
	s[0] = "Go"
	s = append(s, "playground")
	return s
}

Option 2: Avoid append inside variadic functions
🎯 Interview-ready explanation

“Slices share underlying arrays, so element modifications are visible to the caller. However, append may allocate a new array if capacity is exceeded, so changes after append affect only the local slice unless the returned slice is used.”