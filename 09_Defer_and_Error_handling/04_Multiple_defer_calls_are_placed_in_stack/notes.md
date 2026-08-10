This example is very good because it shows the **real behavior of multiple `defer` statements**.

The key concept:

> **`defer` follows LIFO (Last In First Out) order.**

Think of it like a stack of plates.

You put plates:


Plate 1
Plate 2
Plate 3


You remove:


Plate 3
Plate 2
Plate 1


The last one added comes out first.

---

## Code


func main() {
	str := "Gopher"

	fmt.Printf("Original String: %s\n", string(str))

	fmt.Printf("Reversed String: ")

	for _, v := range str {
		defer fmt.Printf("%c", v)
	}
}


---

# Step 1: Initial string


str := "Gopher"


Memory:


Index:

0   1   2   3   4   5

G   o   p   h   e   r


---

# Step 2: First print


fmt.Printf("Original String: %s\n", string(str))


Output:


Original String: Gopher


---

# Step 3: Second print


fmt.Printf("Reversed String: ")


Output:


Reversed String: 


No newline.

---

# Step 4: Loop starts

This:


for _, v := range str


iterates each character.

First iteration:


v = 'G'


Execute:


defer fmt.Printf("%c", v)


Important:

It does NOT print G now.

It stores the deferred call.

Stack:


TOP
 |
 | fmt.Printf("%c",'G')
 |
BOTTOM


---

Second iteration:


v = 'o'


Add defer:

Stack:


TOP
 |
 | fmt.Printf("%c",'o')
 | fmt.Printf("%c",'G')
 |
BOTTOM


---

Third:


v = 'p'


Stack:


TOP
 |
 | Print('p')
 | Print('o')
 | Print('G')
 |
BOTTOM


---

Fourth:


v = 'h'


Stack:


TOP
 |
 | Print('h')
 | Print('p')
 | Print('o')
 | Print('G')


---

Fifth:


v = 'e'


Stack:


TOP
 |
 | Print('e')
 | Print('h')
 | Print('p')
 | Print('o')
 | Print('G')


---

Sixth:


v = 'r'


Final stack:


TOP
 |
 | Print('r')
 | Print('e')
 | Print('h')
 | Print('p')
 | Print('o')
 | Print('G')
 |
BOTTOM


---

# Step 5: main function ends

Now Go says:

> "The function is returning. Execute all deferred functions."

Remember:

Stack is LIFO.

So:

First:


Print('r')


Output:


r


Remove it.

---

Next:


Print('e')


Output:


re


---

Next:


Print('h')


Output:


reh


---

Next:


Print('p')


Output:


rehp


---

Next:


Print('o')


Output:


rehpo


---

Last:


Print('G')


Output:


rehpoG


---

# Final output


Original String: Gopher
Reversed String: rehpoG


---

# The important syntax detail

This line:


defer fmt.Printf("%c", v)


has two things:

### 1. Function call


fmt.Printf("%c", v)


Normally executes immediately.

---

### 2. defer keyword


defer


changes the timing:

Instead of:


Execute now


it becomes:


Store this call.
Execute when main returns.


---

# Important: Why does it reverse?

Because there are two opposite orders:

The loop order:


G → o → p → h → e → r


Deferred storage:


G
Go
Gop
Goph
Gophe
Gopher


Execution order:


r → e → h → p → o → G


So:


rehpoG


---

# Real backend connection

You may see this pattern in cleanup:


func process() {

    file1.Open()
    defer file1.Close()

    file2.Open()
    defer file2.Close()

}


Stack:


TOP
 |
 Close file2
 Close file1
 |
BOTTOM


Execution:


Close file2 first
Close file1 second


Why?

Because the last resource acquired is usually the first resource released.

---

# Interview explanation

If asked:

> "What happens when multiple defer statements exist?"

Answer:

> "Deferred function calls are stored in a stack. When the surrounding function returns, deferred calls execute in reverse order of registration because Go follows LIFO order. The last deferred function is executed first."

---

One thing to remember:


defer = execute later

multiple defer = stack

stack = LIFO

LIFO = reverse order


This single concept explains the entire example.
