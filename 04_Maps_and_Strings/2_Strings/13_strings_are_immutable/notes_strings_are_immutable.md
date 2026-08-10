1️⃣ What “strings are immutable” really means

Immutable = once a string is created, its contents cannot be changed

You can:

read a string

compare it

copy it

convert it

❌ You cannot modify any character inside it.

2️⃣ First program (WHY it fails)
Code
func mutate(s string) string {
	s[0] = 'a'
	return s
}

Syntax-level explanation
🔹 s string

s is a string

Internally: a read-only byte sequence

🔹 s[0]

Indexing a string gives a byte

But that byte is read-only

🔹 'a'

Single quotes → rune

Rune is an alias for int32

Represents Unicode character 'a'

🔹 s[0] = 'a'

This line tries to:

change the first byte of the string

which Go does not allow

Compiler error (key meaning)
cannot assign to s[0]
(neither addressable nor a map index expression)


This means:

s[0] does not have a writable memory address

Strings are immutable by design

🚫 Compile-time error (not runtime)

Execution flow (first program)
main()
 ├── h := "hello"
 ├── mutate(h)
 │     ├── attempt s[0] = 'a'
 │     └── ❌ compilation fails


The program never runs.

3️⃣ WHY Go makes strings immutable (important WHY)

Go strings are immutable because:

Strings may be shared internally

Safe for concurrency

Faster comparisons & memory reuse

Prevents accidental data corruption

👉 This is a design decision, not a limitation.

4️⃣ Correct workaround: convert to []rune

Now let’s see the working version.

Second program (correct way)
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

Syntax-level explanation
🔹 s []rune

s is a slice

Slices are mutable

Each element is a rune (Unicode character)

🔹 s[0] = 'a'

Valid assignment

You are modifying a slice element

Completely allowed

🔹 string(s)

Converts rune slice → new string

This creates a new immutable string

✔️ Original string is untouched
✔️ New string is returned

main() in second program
h := "hello"
fmt.Println(mutate([]rune(h)))

What happens step-by-step

"hello" → string

[]rune(h) → rune slice: ['h','e','l','l','o']

mutate modifies index 0 → 'a'

Convert back to string → "aello"

Print result

Execution flow (second program)
main()
 ├── h := "hello"
 ├── []rune(h) → ['h','e','l','l','o']
 ├── mutate(...)
 │     ├── s[0] = 'a'
 │     └── return "aello"
 └── print "aello"

5️⃣ Why rune slice, not byte slice?

You could use []byte, but:

❌ Breaks Unicode characters
✅ []rune safely handles Unicode

Example:

"ñ" → []byte = 2 bytes
"ñ" → []rune = 1 rune

6️⃣ Key rules (remember these)
Type	Mutable?	Purpose
string	❌ No	Text storage
[]byte	✅ Yes	Raw data
[]rune	✅ Yes	Unicode-safe text manipulation
7️⃣ One-line mental model (VERY IMPORTANT)

Strings are read-only views of bytes.
To change text, convert → modify → create a new string.

8️⃣ Final summary (simple words)

Go strings cannot be modified once created.
Attempting to change a character causes a compile-time error.
To modify a string, convert it to a rune slice, make changes, and convert it back to a new string.