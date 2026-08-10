“Slices are passed by value, but changes inside the function are visible outside. HOW??”

Let’s clear this completely.

1️⃣ First: What a slice REALLY is (mentally)

A slice is NOT the data itself.

Internally, a slice looks like this (conceptually):

type slice struct {
    ptr *T   // points to underlying array
    len int
    cap int
}


So a slice is just:

a pointer to an array

plus length

plus capacity

👉 The actual elements live in an array somewhere else

2️⃣ Your code (simplified)
func subtractOne(numbers []int) {
    for i := range numbers {
        numbers[i] -= 2
    }
}

func main() {
    nos := []int{8, 7, 6}
    subtractOne(nos)
}

3️⃣ Step-by-step FLOW (this is the key)
🟢 Step 1: Slice creation in main
nos := []int{8, 7, 6}


What happens internally:

Array in memory:     [8, 7, 6]
Slice nos:
    ptr → address of array[0]
    len = 3
    cap = 3


So:

nos ──► [8, 7, 6]

🟢 Step 2: Function call
subtractOne(nos)


⚠️ Important:

The slice header is copied (pass by value)

BUT the pointer inside it still points to the SAME array

So now:

main:
  nos ──► [8, 7, 6]

subtractOne:
  numbers ──► [8, 7, 6]


Different slice variables
✅ Same underlying array

🟢 Step 3: Loop inside function
numbers[i] -= 2


Iteration by iteration:

i	Before	After
0	8	6
1	7	5
2	6	4

What is actually modified?

👉 The array, not the slice header.

So the array becomes:

[6, 5, 4]

🟢 Step 4: Function returns

The local variable numbers is destroyed
BUT the array remains modified.

Back in main:

nos ──► [6, 5, 4]


So when you print:

fmt.Println(nos)


You see:

[6 5 4]

4️⃣ Why arrays behave differently (VERY IMPORTANT)
Array example:
func change(arr [3]int) {
    arr[0] = 100
}


Here:

Entire array is copied

Function works on its own copy

Memory picture:

main array    [8 7 6]
function arr  [8 7 6]   ← separate copy


So original stays unchanged.

5️⃣ Why slices behave like this (DESIGN REASON)

Go designers wanted:

Efficient passing of large data

No copying of big arrays

Easy modification of shared data

So slices are lightweight descriptors, not containers.

6️⃣ One-line rule (memorize this)

Arrays copy data.
Slices copy pointers to data.

7️⃣ Real-world analogy 🧠

Think of:

Array = actual house

Slice = Google Maps location pin

Passing a slice to a function:

You give someone the same address

If they change furniture, everyone sees it

Passing an array:

You give them a photocopy of the house

Changes don’t affect the original