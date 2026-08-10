🔹 Your code (reference)
currencyCode := map[string]string{
	"USD": "US Dollar",
	"GBP": "Pound Sterling",
	"EUR": "Euro",
}

currency := "USD"
currencyName := currencyCode[currency]

fmt.Println("Currency name for currency code", currency, "is", currencyName)

1️⃣ Syntax-level explanation
Map type
map[string]string


key type → string

value type → string

So:

currencyCode["USD"] → string

Retrieving a value
currencyName := currencyCode[currency]


Break it down:

Part	Meaning
currencyCode	the map
[currency]	lookup using key
currencyName	receives the value

📌 This is lookup by key, not index.

2️⃣ Execution flow (step-by-step)
Step 1: Map creation
currencyCode := map[string]string{...}


Internally:

Go creates a hash table

Stores key–value pairs

Step 2: Key variable
currency := "USD"


Just a normal string variable.

Step 3: Map lookup
currencyName := currencyCode[currency]


What Go does internally:

Hash "USD"

Look up the bucket

Find matching key

Return its value → "US Dollar"

So now:

currencyName == "US Dollar"

Step 4: Print
fmt.Println(...)


Output:

Currency name for currency code USD is US Dollar

3️⃣ IMPORTANT: What if the key does NOT exist?

This is where many beginners get confused.

Example:

currency := "INR"
currencyName := currencyCode[currency]
fmt.Println(currencyName)

What happens?

No panic

No error

Go returns zero value of the value type

Since value type is string → zero value is "" (empty string)

So:

currencyName == ""


⚠️ This can be dangerous if you don’t check.

4️⃣ Correct way: the “comma ok” idiom (VERY IMPORTANT)
currencyName, ok := currencyCode[currency]


Meaning:

currencyName → value (or zero value)

ok → true if key exists, false otherwise

Example:

if currencyName, ok := currencyCode[currency]; ok {
	fmt.Println("Currency name:", currencyName)
} else {
	fmt.Println("Currency code not found")
}


This is the standard Go way.

5️⃣ Why Go behaves this way (design reason)

Go designers wanted:

Fast lookups

No exceptions

Simple control flow

So instead of throwing errors, Go gives:

zero value

optional ok flag

6️⃣ Real-world usage (VERY PRACTICAL)
Config lookup
timeout, ok := config["TIMEOUT"]

Cache lookup
value, found := cache[key]

HTTP headers
contentType, ok := headers["Content-Type"]

7️⃣ Mental model (lock this in)
map[key] → value OR zero value
map[key], ok → safe lookup

8️⃣ Rules to remember (SAVE THIS)

map[key] retrieves a value

Missing key → zero value

No panic on missing key

Use value, ok := map[key] when key may not exist

Always use ok in production code

✅ One-line takeaway

Map lookup never fails in Go — it returns a value and optionally tells you whether the key existed.