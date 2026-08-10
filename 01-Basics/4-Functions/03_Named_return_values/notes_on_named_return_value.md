1️⃣ What are named return values (syntax-level)

Your example:

func rectProps(length, width float64) (area, perimeter float64) {
area = length _ width
perimeter = (length + width) _ 2
return
}

What happens internally

When Go sees this:

(area, perimeter float64)

It automatically does this behind the scenes:

func rectProps(length, width float64) (area float64, perimeter float64) {
var area float64
var perimeter float64
// function body
}

So:

area and perimeter are real variables

They live for the entire function

return means → “return current values of these variables”

2️⃣ Why return works without values
return

This is called a naked return.

📌 Rule

Naked return is allowed only when return values are named.

Equivalent to:

return area, perimeter

3️⃣ Is this the standard way?
❌ NOT the default style

Most Go code does NOT use named returns for simple functions.

This is the most common and recommended style:

func rectProps(length, width float64) (float64, float64) {
area := length _ width
perimeter := (length + width) _ 2
return area, perimeter
}

Why?

Clear

Explicit

Easier to read

Easier to debug

4️⃣ When ARE named return values used? (REAL WORLD)

Named returns are used, but only in specific cases.

✅ Case 1: Error handling with deferred logic (VERY COMMON)
func readFile(name string) (data []byte, err error) {
file, err := os.Open(name)
if err != nil {
return
}
defer file.Close()

    data, err = io.ReadAll(file)
    return

}

Here:

defer can modify err

Named return makes this clean

📌 This is the #1 valid use case

✅ Case 2: Documentation clarity (rare)
func stats(nums []int) (min int, max int) {
// logic
return
}

Helps readers understand intent without comments.

❌ Case 3: Small functions (NOT recommended)
func add(a, b int) (sum int) {
sum = a + b
return
}

This is considered unnecessary and less clear.

5️⃣ Hidden danger of named returns ⚠️

Named returns can silently return wrong values.

Example:

func divide(a, b int) (result int) {
if b == 0 {
return // returns result = 0 ❌ silently
}
result = a / b
return
}

This can cause bugs that are hard to notice.

6️⃣ Go community best practice (IMPORTANT)
✔ Preferred (most of the time)
return area, perimeter

⚠ Use named returns only when:

Function is long

defer modifies return values

Error handling is involved

It improves clarity (rare)

7️⃣ Simple rule to follow (remember this)

If you are learning Go or writing simple functions → DO NOT use named returns.

If you see them in real projects → understand them, but don’t overuse them.

8️⃣ Your rectangle example — final recommendation
Best version for you right now ✅
func rectProps(length, width float64) (float64, float64) {
area := length _ width
perimeter := (length + width) _ 2
return area, perimeter
}

9️⃣ Interview-ready answer (save this)

“Go supports named return values, but they are mainly used in functions involving defer and error handling. For most functions, explicit return values are preferred for clarity.”
