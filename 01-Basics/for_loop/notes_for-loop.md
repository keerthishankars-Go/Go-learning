Here's a clear and simple explanation of your Go code:

for n := range 6 {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}

❗ First: This is a new Go 1.22 feature

In Go 1.22, you can use:

for i := range <integer>


This runs a loop from 0 to integer-1.

So:

for n := range 6


Means:

n = 0
n = 1
n = 2
n = 3
n = 4
n = 5

🔍 Step-by-step execution
Loop:
for n := range 6  → n goes from 0 to 5

Inside the loop:
if n%2 == 0 {
    continue
}


n % 2 == 0 → true for even numbers

continue → skip the rest of this iteration

So "even numbers" are skipped.

Then print odd numbers:
fmt.Println(n)

✔ Final Output
1
3
5


Because only 1,3,5 are odd.

💡 Summary (easy to remember)
range 6

→ gives: 0,1,2,3,4,5

n%2 == 0

→ skip even numbers

fmt.Println(n)

→ print odd numbers