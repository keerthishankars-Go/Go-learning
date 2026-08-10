What is a function?

A function is a block of code that performs a specific task.

A function takes an input, performs some operations on the input and generates outputs.
For example, a function can take the radius as input and calculate the area and circumference as output.

Function declaration
The following is the syntax for declaring a function in Go

func functionname(parametername datatype) returntype {

//function body

}

The function declaration starts with the func keyword followed by the functionname. The parameters are specified between ( and ) followed by the returntype of the function.
The syntax for specifying a parameter is, parameter name followed by the type. Any number of parameters can be specified like (parameter1 datatype, parameter2 datatype). Then there is a block of code between { and } which is the body of the function.

When multiple parameters have the same type, write type once at the end.

The parameters and return type are optional in a function. Hence the following is also a valid function declaration.

func functionname() {
}

# ============================

Golden Rules (MEMORIZE THIS)

✅ main()

no parameters

no return value

✅ Business logic goes into separate functions

func calculateBill(...) int

Rule :If function returns something → must specify return type

If no return → omit it

✅ main() only:

prepares data

calls functions

prints results

🚀 Why Go enforces this (important for backend/devops)

Predictable startup

Simple binaries

Easy deployment

Clean architecture

This is why Go is loved in production systems.

# ===============================

Multiple assignment using :=
price, quantity := 90, 6

What happens here

Declares two variables

Assigns values in one line

Equivalent to:

price := 90
quantity := 6

📌 Rule

Left side count must match right side count

At least one variable must be new

# =====================================

🧠 FINAL RULES TO KEEP IN MIND (SAVE THIS)
🔹 Functions

func name(params) returnType

Same-type params → use comma

Explicit return is mandatory

🔹 main() rules

Only one main()

No arguments

No return value

Entry point of program

🔹 Variables

:= → inside functions

var → explicit or package-level

🔹 Function calling

Order matters

Types must match

Return value must be used or ignored

🔹 Package rule

All .go files in a folder must have same package name

# =====================================================

🔹 (float64, float64) ← multiple return types

This is the key new concept.

It means:

Function returns TWO values

Both are of type float64

Order matters

📌 Rule

When returning more than one value, return types must be wrapped in parentheses.

# =======================================================

🔹 return area, perimeter
return area, perimeter

Why comma , here?

Go allows returning multiple values

Order of values must match return types

This corresponds to:

(float64, float64)

📌 Rule

Number of returned values must exactly match function signature.
