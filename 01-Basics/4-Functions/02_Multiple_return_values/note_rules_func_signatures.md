🔒 RULES TO MEMORIZE (SHORT LIST)
✅ Function signature
func name(params) (type1, type2)

✅ Return statement
return value1, value2

✅ Function call
a, b := function()

✅ Ignore values
a, \_ := function()

❌ NOT allowed
return area // missing one value ❌
area := rectProps() // missing variables ❌

# =================================================

1️⃣ Function signature — MOST IMPORTANT PART
func rectProps(length, width float64) (float64, float64)

Let’s split it:

🔹 rectProps

Function name.

🔹 length, width float64

Same rule you already know:

length float64
width float64

Shortened using comma.

📌 Rule

Multiple parameters with same type → write type once at the end.

🔹 (float64, float64) ← multiple return types

This is the key new concept.

It means:

Function returns TWO values

Both are of type float64

Order matters

📌 Rule

When returning more than one value, return types must be wrapped in parentheses.

2️⃣ Function body
var area = length _ width
var perimeter = (length + width) _ 2

Nothing special here:

Calculations using parameters

Stored in local variables

🔹 return area, perimeter
return area, perimeter

Why comma , here?

Go allows returning multiple values

Order of values must match return types

This corresponds to:

(float64, float64)

📌 Rule

Number of returned values must exactly match function signature.

3️⃣ Calling a multi-return function
area, perimeter := rectProps(10.8, 5.6)

What happens step by step

rectProps runs

Returns (areaValue, perimeterValue)

First value → assigned to area

Second value → assigned to perimeter

📌 Rule

Left side variables must match number and order of returned values.

4️⃣ Ignoring return values (VERY IMPORTANT)

If you don’t want one value:

area, \_ := rectProps(10.8, 5.6)

\_ (blank identifier)

Tells Go: “I don’t care about this value”

Prevents unused variable error

📌 Rule

Use \_ to ignore unwanted return values.

5️⃣ Printing with Printf
fmt.Printf("Area %f Perimeter %f", area, perimeter)

Why %f

%f → float values

Better readable version:

fmt.Printf("Area %.2f, Perimeter %.2f\n", area, perimeter)

Output:

Area 60.48, Perimeter 32.80


===================================================
===================================================

🧠 WHY Go supports multiple returns (REAL REASON)
🔹 1. Error handling (MOST COMMON)
value, err := doSomething()


Instead of exceptions → Go uses explicit returns.

🔹 2. Avoid structs for small results

Instead of:

type Result struct {
	Area float64
	Perimeter float64
}


You just return:

(area, perimeter)

🔹 3. Clean APIs (backend style)

You’ll see this everywhere:

user, found := getUserByID(id)
data, ok := cache.Get(key)
n, err := file.Read(buf)