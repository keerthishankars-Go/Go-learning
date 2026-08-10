✅ 1. What is a Variadic Function?

A variadic function accepts any number of arguments.

Example:

sum(1)
sum(1, 2)
sum(1, 2, 3, 4, 5)

All of these are valid if sum is variadic.

✅ 2. Syntax: func sum(nums ...int)
func sum(nums ...int) {

Meaning:

nums is the parameter name.

...int means the function accepts 0 or more integers.

Inside the function, nums behaves like a slice: []int

📌 So inside the function:

nums ...int → nums []int

That’s why you can do:

len(nums)

range nums

nums[0]

append (not recommended directly because nums is read-only slice copy)

✅ 3. Full Code
func sum(nums ...int) {
fmt.Print(nums, " ")
total := 0

    for _, num := range nums {
        total += num
    }

    fmt.Println(total)

}

Line-by-line explanation:
🔹 fmt.Print(nums, " ")

Prints the slice of numbers passed in.

🔹 total := 0

A counter variable.

🔹 for \_, num := range nums {

Loop through each number in the slice.

\_ → ignore index

num → the value at each index

Example:
If input was sum(1, 2, 3) → nums = []int{1,2,3}

Loop runs as:

num = 1

num = 2

num = 3

🔹 total += num

Adds each number.

🧪 4. Using the function
sum(1, 2)
sum(1, 2, 3)

Output:

[1 2] 3
[1 2 3] 6

🧠 5. Passing a slice into a variadic function

If you already have a slice:

nums := []int{1, 2, 3, 4}
sum(nums)

This will not work ❌:

cannot use nums (value of type []int) as type int in argument to sum

Because the function expects individual ints, not a slice.

❗ Solution: use ...
sum(nums...)

The ... expands the slice into individual arguments:

sum([]int{1,2,3,4}...)
→ sum(1,2,3,4)

✔ This is the only time you write ... while calling a function.

🔥 6. WHY do we use variadic functions?

In backend work:

✔ Logging
log.Println("User:", id, "Status:", status)

✔ Validation
validateFields(name, email, phone)

✔ Summation / Analytics
average(scores...)

✔ Utility functions
max(nums...)
min(nums...)

✔ String building
fmt.Println("a", "b", "c") // Println is variadic!

💡 7. Why does the type of nums become []int?

Because internally Go uses a slice to hold all the passed arguments.

nums ...int → Go converts → nums []int

Slices are dynamic lists. Perfect fit for this feature.

🎉 Final Summary
Syntax Meaning
nums ...int accept any number of ints
nums inside function becomes []int
for \_, num := range nums loop through all
sum(nums...) expand a slice into arguments

# =============================================

✅ 1. What does nums ...int mean?

Think of:

nums ...int

as:

👉 “This function can take ANY number of int values AND store them into one slice called nums.”

Example:

sum(1, 2, 3, 10)

Inside the function, Go converts it to:

nums = []int{1, 2, 3, 10}

So nums is just a slice automatically created by Go.

✅ 2. Why must it be written as ...int (no spaces, no extra dots)?

Because in Go:

... and int together make one complete type:
👉 "variadic integers"

So the only correct way is:

...int

Go reads it as:

"A variable number of ints"

If you write:

nums....int
nums... int
nums .. .int

Go cannot understand the type, so it throws an error.

✅ 3. Why do we write func sum(nums ...int)?

Because the function must know two things:

The variable name → nums

The type → ...int (variadic integers)

So the full meaning:

func sum(nums ...int)

=
👉 “Create a function named sum that accepts unlimited integers, call them nums.”

✅ 4. What does range nums mean inside the function?

Since nums becomes a slice, like:

[]int{1, 2, 3}

You can loop it:

for \_, num := range nums {
total += num
}

\_ → ignore index

num → each value in the slice

✅ 5. Why do we use nums... when calling the function with a slice?

If you already have a slice:

values := []int{1, 2, 3}

To pass it to a variadic function, write:

sum(values...)

This means:

👉 “Expand the slice into individual integers.”

So Go does:

sum(1, 2, 3)

📦 Summary in One Line
✔ nums ...int

Means → “Accept unlimited integers, store them in slice nums.”

✔ Only ...int is valid

No extra dots or spaces.

✔ Inside the function

nums becomes a slice → you can loop, access length, etc.

✔ When passing a slice

Use slice... to expand.

# ===============================================

✅ 1. Variadic Arguments for Logging

WHY?
Logs often need to accept any number of values.

Sometimes you log one value, sometimes five:

log("User logged in")
log("Payment success", orderId, amount)
log("Error", err, "User:", userID)

A regular function cannot accept unlimited values.

⭐ Variadic function solution:
func log(values ...any) {
fmt.Println(values...)
}

Now the logger can accept 1 value, 3 values, or 100 values.

✔ Practical Example Output
log("User logged in")
// ["User logged in"]

log("Payment success", 123, 499.00)
// ["Payment success" 123 499]

log("Error", err)
// ["Error" <error value>]

Why backend developers use this?

Logging events

Debugging

Sending logs to monitoring tools (Grafana, Loki, Cloud Logging)

✅ 2. Variadic Args for Combining Database Query Filters

WHY?
In backend APIs, users may filter in many ways:

Example API:

GET /users?age=20&country=IN&active=true

You don’t know how many filters will be applied.

⭐ Variadic Functions help you add dynamic SQL filters:
func buildQuery(base string, filters ...string) string {
for \_, f := range filters {
base += " AND " + f
}
return base
}

✔ Using it:
query := buildQuery(
"SELECT \* FROM users WHERE 1=1",
"age = 20",
"country = 'IN'",
"active = true",
)

fmt.Println(query)

Output:
SELECT \* FROM users WHERE 1=1 AND age = 20 AND country = 'IN' AND active = true

Why backend devs use this?

Building flexible search filters

Adding conditions only if the user requested them

SQL query builders (GORM, Squirrel, Ent use this idea)

✅ 3. Variadic Args for Building API Responses

WHY?
Sometimes API responses include optional fields.

Example Response:
Some APIs return:

{ "message": "success" }

Others:

{
"message": "success",
"user": { ... },
"token": "abc"
}

⭐ Variadic function solution:
func response(msg string, data ...map[string]any) map[string]any {
res := map[string]any{"message": msg}

    for _, obj := range data {
        for k, v := range obj {
            res[k] = v
        }
    }
    return res

}

✔ Using it:
fmt.Println(response("OK"))

Output:

{"message":"OK"}

Add extra data:

fmt.Println(response("OK", map[string]any{
"user": user,
}))

Output:

{
"message": "OK",
"user": { ... }
}

Add even more fields:

response("OK",
map[string]any{"token": token},
map[string]any{"profile": profile},
)

Why backend devs use this?

Flexible JSON responses

Reduce duplicated code

Send optional data cleanly

🎯 SUPER SIMPLE SUMMARY
Use Case Why Variadic? Example
Logging Logs take any number of values log("err", user, ip)
DB Filters Queries need dynamic conditions buildQuery("SELECT *", "age=20", "country='IN'")
API Responses API may return optional fields response("OK", map1, map2...)
