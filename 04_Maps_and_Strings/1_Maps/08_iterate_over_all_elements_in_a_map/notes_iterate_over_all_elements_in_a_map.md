🔹 Your code (reference)
for code, name := range currencyCode {
	fmt.Printf("Currency Name for currency code %s is %s\n", code, name)
}

1️⃣ Syntax-level explanation
range on a map
for key, value := range mapName {
    // use key and value
}


Applied to your code:

Variable	Meaning
code	map key (string)
name	map value (string)
currencyCode	map being iterated

📌 range returns two values for a map:

key

value

2️⃣ Execution flow (step-by-step)
Step 1: Map exists in memory
currencyCode := map[string]string{
	"USD": "US Dollar",
	"GBP": "Pound Sterling",
	"EUR": "Euro",
}


Internally (unordered):

USD → US Dollar
GBP → Pound Sterling
EUR → Euro

Step 2: range starts iterating
for code, name := range currencyCode {


What Go does:

Picks any key-value pair

Assigns:

code = key
name = value


Executes loop body

Repeats for remaining entries

⚠️ Order is NOT guaranteed

Step 3: Loop body executes
fmt.Printf("Currency Name for currency code %s is %s\n", code, name)


Printed once per key-value pair.

3️⃣ Why output order is random (VERY IMPORTANT)

Example output:

GBP → Pound Sterling
EUR → Euro
USD → US Dollar


Another run may print:

USD → US Dollar
GBP → Pound Sterling
EUR → Euro

Why?

Because:

Maps are hash tables

Hash iteration order is randomized

Go intentionally prevents relying on order

📌 Never depend on map iteration order

4️⃣ Ignoring key or value (common patterns)
Ignore value
for code := range currencyCode {
	fmt.Println(code)
}

Ignore key
for _, name := range currencyCode {
	fmt.Println(name)
}


_ = blank identifier (ignore variable)

5️⃣ Real-world usage (VERY PRACTICAL)
Logging
for k, v := range headers {
	log.Println(k, v)
}

JSON building
for field, value := range payload {
	result[field] = value
}

Validation
for key := range input {
	if !allowed[key] {
		return error
	}
}

6️⃣ How to iterate in a fixed order (IMPORTANT)

Maps cannot be ordered.
If order matters → use a slice of keys.

keys := make([]string, 0, len(currencyCode))
for k := range currencyCode {
	keys = append(keys, k)
}

sort.Strings(keys)

for _, k := range keys {
	fmt.Println(k, currencyCode[k])
}

7️⃣ Mental model (lock this in)
map = unordered key-value store
range = visit every pair once
order = undefined

8️⃣ Rules to remember (SAVE THIS)

range iterates over all map entries

Returns key, value

Order is random

Never rely on iteration order

Use slices if order matters

✅ One-line takeaway

range lets you iterate over all key–value pairs in a map, but the order is undefined.