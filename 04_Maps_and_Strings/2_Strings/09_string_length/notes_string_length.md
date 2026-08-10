1️⃣ Imports & setup (syntax level)
import (
	"fmt"
	"unicode/utf8"
)

What this means

fmt → printing

unicode/utf8 → utilities for UTF-8 encoded strings

Go strings are UTF-8 by default

This package understands Unicode rules

2️⃣ main() starts execution
func main() {


Every Go program starts here.

3️⃣ First string: "Señor"
word1 := "Señor"


Declares a string

Internally stored as UTF-8 bytes

Actual byte layout:
S   e   ñ       o   r
53  65  c3 b1   6f  72

4️⃣ Printing the string
fmt.Printf("String: %s\n", word1)


%s → print string

Output:

String: Señor

5️⃣ Correct string length (characters)
fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))

Syntax breakdown

utf8.RuneCountInString(word1)

Reads UTF-8 bytes

Decodes Unicode characters

Counts runes (characters)

Result
Length: 5


✔️ S, e, ñ, o, r → 5 characters

6️⃣ Byte length (raw storage)
fmt.Printf("Number of bytes: %d\n", len(word1))

What len does

Counts bytes

Does NOT decode Unicode

Result
Number of bytes: 6


Because:

ñ occupies 2 bytes

7️⃣ Separator
fmt.Printf("\n")


Prints a blank line for readability

8️⃣ Second string: "Pets"
word2 := "Pets"


All ASCII characters

Each character = 1 byte

9️⃣ Length & bytes for "Pets"
utf8.RuneCountInString(word2) → 4
len(word2)                   → 4


✔️ Same result because ASCII = single-byte UTF-8

🔟 Execution flow (step-by-step)
main()
 ├── word1 := "Señor"
 ├── print word1
 ├── count runes in word1 → 5
 ├── count bytes in word1 → 6
 ├── blank line
 ├── word2 := "Pets"
 ├── count runes in word2 → 4
 └── count bytes in word2 → 4

1️⃣1️⃣ Why len(s) is NOT string length (key insight)
Expression	What it returns
len(s)	Number of bytes
utf8.RuneCountInString(s)	Number of characters

Because:

Go strings = UTF-8 bytes

Some characters need multiple bytes

1️⃣2️⃣ When you MUST NOT use len(s)

❌ Password length validation
❌ Username character limits
❌ SMS character count
❌ UI limits
❌ Any user-facing text logic

1️⃣3️⃣ Correct patterns (production-grade)
Character count
count := utf8.RuneCountInString(s)

Byte size (DB / network)
size := len(s)

1️⃣4️⃣ One-line mental model (remember forever)

len(string) counts bytes
RuneCountInString(string) counts characters

1️⃣5️⃣ Final summary (as you requested)

The program shows that Go strings are byte-based.
len() counts raw bytes, while utf8.RuneCountInString() correctly counts Unicode characters.
For ASCII strings both values match, but for Unicode strings they differ.