1. Declaring an uninitialized slice
   var s []string
   fmt.Println("uninit:", s, s == nil, len(s) == 0)

Meaning:

var s []string → creates a nil slice

A nil slice has:

len = 0

cap = 0

value of s is nil

Output:
uninit: [] true true

✅ 2. Creating a slice using make
s = make([]string, 3)
fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

Meaning:

make([]string, 3) → creates a slice of length 3

Values default to empty strings: ["", "", ""]

When you specify only length, capacity = length.

Output:
emp: [ ] len: 3 cap: 3

✅ 3. Setting and accessing elements
s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("set:", s)
fmt.Println("get:", s[2])

Meaning:

Normal indexing like arrays.

Output:
set: [a b c]
get: c

✅ 4. Append elements
s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("apd:", s)

Meaning:

append() creates a new slice if current capacity is full.

The slice grows automatically.

Example growth:
[a b c] → append d → [a b c d]
→ append e,f → [a b c d e f]

Output:
apd: [a b c d e f]

✅ 5. Copying slices
c := make([]string, len(s))
copy(c, s)
fmt.Println("cpy:", c)

Meaning:

copy(dst, src) copies elements.

c becomes a separate slice (not linked to s).

Output:
cpy: [a b c d e f]

✅ 6. Slice ranges / slicing
l := s[2:5]
fmt.Println("sl1:", l)

s[2:5] → elements from index 2 to 4

(5 not included)

sl1: [c d e]

l = s[:5]
fmt.Println("sl2:", l)

s[:5] → start from 0 to index 4
sl2: [a b c d e]

l = s[2:]
fmt.Println("sl3:", l)

s[2:] → from index 2 to the end
sl3: [c d e f]

✅ 7. Slice literal (declared directly)
t := []string{"g", "h", "i"}
fmt.Println("dcl:", t)

Output:

dcl: [g h i]

✅ 8. Comparing slices using slices.Equal
t2 := []string{"g", "h", "i"}
if slices.Equal(t, t2) {
fmt.Println("t == t2")
}

Meaning:

Go does NOT allow comparing slices using ==.
You must use the slices package.

Output:

t == t2

✅ 9. Two-dimensional slices
twoD := make([][]int, 3)

This creates a slice with 3 rows, but each row is nil initially.

Build the inner slices:
for i := range 3 {
innerLen := i + 1
twoD[i] = make([]int, innerLen)
for j := range innerLen {
twoD[i][j] = i + j
}
}

What happens:

i innerLen row created values assigned
0 1 [0] [0]
1 2 [0 1] [1 2]
2 3 [0 1 2] [2 3 4]
Final output:
2d: [[0] [1 2] [2 3 4]]

🎉 FINAL SUMMARY (simple + powerful)
✔ Slice vs Array
Feature Slice Array
Size dynamic fixed
Can grow? YES (append) NO
Nil allowed? YES NO
Passed by reference (backed by same array) value
Comparison slices.Equal == only works for arrays
