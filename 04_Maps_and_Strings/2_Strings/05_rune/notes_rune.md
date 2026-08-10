1️⃣ What is a rune (syntax-level first)
type rune = int32


rune is a built-in alias

Stores one Unicode code point

Size: 4 bytes (always)

Represents a character, not raw bytes

Even if a character uses 2–4 bytes in UTF-8,
one rune represents it as a single value

2️⃣ New function: printChars
func printChars(s string) {


Defines a function

Takes a string

Purpose: print characters (not bytes)

Step inside printChars
fmt.Printf("Characters: ")


Prints label text

runes := []rune(s)

MOST IMPORTANT LINE 🔥
What happens here

Converts string → []rune

Go decodes UTF-8

Each Unicode character becomes one rune

Example:
"Señor"
↓
[]rune{'S','e','ñ','o','r'}

for i := 0; i < len(runes); i++ {


Loop over characters

len(runes) = number of characters

fmt.Printf("%c ", runes[i])


%c → prints a character

runes[i] → one Unicode code point

3️⃣ printBytes (unchanged but now meaningful)
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}


len(s) → bytes

s[i] → byte

%x → hex representation

4️⃣ Execution flow (VERY CLEAR)
First run: "Hello World"
name := "Hello World"
printChars(name)
printBytes(name)

Internally:

All characters are ASCII

Each character = 1 byte

Bytes == Characters

✔️ Output matches exactly

Second run: "Señor"
name = "Señor"
printChars(name)
printBytes(name)

What Go sees internally
UTF-8 encoding:
Character	UTF-8 bytes
S	53
e	65
ñ	c3 b1
o	6f
r	72
printChars
Characters: S e ñ o r


ñ → one rune

Correct character count = 5

printBytes
Bytes: 53 65 c3 b1 6f 72


ñ → two bytes

Byte count = 6

5️⃣ Why []rune(s) is needed (WHY question)

Strings in Go are byte sequences

Unicode characters may be multi-byte

Indexing a string gives bytes

Converting to []rune:

Decodes UTF-8

Makes indexing safe for characters

6️⃣ Key syntax differences (table)
Expression	Meaning
string	UTF-8 encoded bytes
[]byte(s)	Raw bytes
[]rune(s)	Unicode characters
len(s)	Bytes
len([]rune(s))	Characters
%x	Hex byte
%c	Character
7️⃣ One-line mental model (REMEMBER THIS)

String → bytes
Rune → character

Or:

Index string → byte
Range / rune → character

8️⃣ Why this matters in real backend code

This affects:

User names

Password length validation

SMS / Email limits

Database field sizes

Internationalization (Indian languages, emojis)

❌ len(password) → WRONG
✅ len([]rune(password)) → CORRECT

9️⃣ Final flow summary (as you wanted)

We first print characters by converting the string to runes, which represent Unicode characters.
Then we print raw bytes to show how UTF-8 stores those characters internally.