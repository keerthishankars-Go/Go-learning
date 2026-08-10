fallthrough is one of the most misunderstood switch features in Go, so let’s explain it slowly, at syntax + control-flow level, and also clarify why Go behaves this way.

1️⃣ First: How switch normally works in Go
Default Go behavior (VERY IMPORTANT)

In Go, switch automatically breaks after a case executes.

That means:

Only one case runs

Control exits the switch

No implicit fall-through (unlike C / Java)

Example (no fallthrough):

switch n {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
}


If n == 1, only "one" prints.

2️⃣ What is fallthrough (core idea)

fallthrough forces execution to continue into the next case, regardless of its condition.

Key points:

It ignores the next case condition

It jumps to the first statement of the next case

It must be the last statement in a case

3️⃣ Your program (reference)
func number() int {
	num := 15 * 5
	return num
}

func main() {
	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

4️⃣ Syntax explanation (line by line)
🔹 switch num := number(); {

This is expressionless switch with assignment.

Equivalent to:

num := number()
switch true {
	...
}


📌 num is:

Evaluated once

Scoped only inside the switch

5️⃣ Case conditions are runtime expressions
case num < 50:
case num < 100:
case num < 200:


These are:

NOT constants

Evaluated top to bottom

First matching case is executed

6️⃣ Step-by-step execution (IMPORTANT)

Assume:

num = 75

Case checks:

num < 50 → ❌ false

num < 100 → ✅ true
→ prints:

75 is lesser than 100


Now Go normally would exit the switch — but…

7️⃣ What fallthrough does here
fallthrough


This tells Go:

“Do NOT exit the switch — continue to the next case.”

So Go:

Does NOT evaluate case num < 200

Directly executes its body

Prints:

75 is lesser than 200

8️⃣ CRITICAL rule (memorize this)

fallthrough does NOT re-check the next case condition

This is why:

case num < 200:


executes even if it were false.

9️⃣ Why fallthrough must be last

❌ Invalid:

case num < 100:
	fallthrough
	fmt.Println("hello")


Why?

Because:

fallthrough is a control-transfer statement

Nothing is allowed after it

Go enforces clarity and predictability

📌 Compiler error:

fallthrough statement out of place

🔟 Why Go designers made it explicit

In C-like languages:

Fallthrough is implicit ❌

Causes many bugs

In Go:
✔ Fallthrough is explicit
✔ You must ask for it
✔ Safer and clearer

11️⃣ Correct mental model

Think of fallthrough as:

“Execute the next case body — no questions asked.”

12️⃣ When should you ACTUALLY use fallthrough?
✅ Valid use cases

Hierarchical classification

Range-based reporting

Cumulative conditions

Example:

switch {
case score >= 90:
	fmt.Println("Excellent")
	fallthrough
case score >= 75:
	fmt.Println("Good")
	fallthrough
case score >= 50:
	fmt.Println("Pass")
}

❌ When NOT to use fallthrough

Normal branching logic

Validation logic

Business rules with strict conditions

In most cases → prefer expressionless switch without fallthrough

13️⃣ Interview-level summary (save this)

“Go switch statements do not fall through by default. The fallthrough keyword explicitly transfers control to the next case without evaluating its condition.”

✅ Final rules to remember

Go switch auto-breaks

fallthrough forces next case execution

Next case condition is NOT checked

fallthrough must be last line

Use sparingly and intentionally