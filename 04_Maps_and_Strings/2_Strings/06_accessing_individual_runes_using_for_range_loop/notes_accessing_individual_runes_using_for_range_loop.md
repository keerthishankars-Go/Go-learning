1️⃣ Overall program structure (quick)
package main
import "fmt"


Executable program

fmt needed for printing

2️⃣ Function: charsAndBytePosition
func charsAndBytePosition(s string) {


Defines a function

Takes a string (UTF-8 encoded bytes)

3️⃣ The MOST IMPORTANT LINE 🔥
for index, rune := range s {

Syntax breakdown
Part	Meaning
for	Loop keyword
range s	Iterate over string by runes
index	Byte index where rune starts
rune	Unicode code point (int32)

⚠️ Important:

index is byte position, NOT character position

rune is one full Unicode character

4️⃣ What Go does internally (very important)

When you do:

for index, rune := range s


Go automatically:

Reads UTF-8 bytes

Decodes them into a rune

Moves index by byte length of that rune

You don’t do this manually — Go does it.

5️⃣ Print statement
fmt.Printf("%c starts at byte %d\n", rune, index)

Part	Meaning
%c	Print character
%d	Print decimal number
rune	Unicode character
index	Starting byte position
6️⃣ main() execution flow
name := "Señor"
charsAndBytePosition(name)


String assigned

Function called

Loop starts iterating

7️⃣ Actual runtime flow (step-by-step)

String: "Señor"

UTF-8 byte layout:
S     e     ñ         o     r
53    65    c3 b1     6f    72
0     1     2  3      4     5   ← byte positions

Iteration steps
1️⃣ First iteration

index = 0

rune = 'S'

S starts at byte 0

2️⃣ Second iteration

index = 1

rune = 'e'

e starts at byte 1

3️⃣ Third iteration

index = 2

rune = 'ñ' (2 bytes!)

ñ starts at byte 2

4️⃣ Fourth iteration

Go skips 2 bytes automatically

index = 4

rune = 'o'

o starts at byte 4

5️⃣ Fifth iteration

index = 5

rune = 'r'

r starts at byte 5

8️⃣ Why o starts at byte 4 (THE KEY INSIGHT)

ñ uses 2 bytes (c3 b1)

So:

byte 2 → start of ñ

byte 3 → continuation

byte 4 → next character

Go’s range knows UTF-8 rules and jumps correctly.

9️⃣ Why for range is BETTER than indexing
Method	What you get	Safe for Unicode?
s[i]	Byte	❌
[]rune(s)[i]	Character	✅
for range s	Character + byte index	✅✅

✔️ for range is preferred in production Go

🔟 One-line mental model (INTERVIEW READY)

for range string decodes UTF-8 and gives runes, while also telling you the byte position where each rune starts.