This example is showing an important combination of:

1. **Methods on structs**
2. **Defer with method calls**
3. **When the deferred method actually executes**

Let's go step by step.

---

## 1. Understand the struct


type person struct {
	firstName string
	lastName string
}


You created a custom type:

text
person

+------------+
| firstName  |
| lastName   |
+------------+


---

## 2. Understand the method


func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}


This is a method.

Normal function:


func fullName(p person)


Method style:


func (p person) fullName()


The receiver:


p person


means:

> This function belongs to the `person` type.

So you can call:


p.fullName()


because `p` is a `person`.

---

# 3. main execution

Create object:


p := person{
	firstName: "John",
	lastName: "Smith",
}


Memory:


p

+----------------+
| John           |
| Smith          |
+----------------+


---

# 4. The important line


defer p.fullName()


Many people think:

> "The name will be printed later, so it will use the latest p."

But the rule:

**Arguments and receiver values are evaluated when defer is registered.**

So this:


defer p.fullName()


is treated like:


defer person.fullName(p)


Go stores a copy of `p`.

At this moment:


p.firstName = John
p.lastName  = Smith


So Go remembers:


Call:

fullName(
    person{
        John,
        Smith,
    }
)


later.

---

# 5. Next line executes


fmt.Printf("Welcome ")


Output:


Welcome 


---

# 6. main function ends

Before returning, Go executes deferred functions.

Now:


p.fullName()


runs.

Inside:


fmt.Printf("%s %s",
    p.firstName,
    p.lastName,
)


prints:


John Smith


---

# Final output


Welcome John Smith


---

# Execution timeline


main starts

        |
        v

create p

p = John Smith


        |
        v

defer p.fullName()

(save method call)


        |
        v

print "Welcome "


        |
        v

main returning


        |
        v

execute deferred method


        |
        v

print "John Smith"


---

# Important experiment

What if you change the value?


func main() {

	p := person{
		firstName: "John",
		lastName: "Smith",
	}

	defer p.fullName()

	p.firstName = "Alex"

	fmt.Printf("Welcome ")
}


What do you expect?

Many think:


Welcome Alex Smith


But output:


Welcome John Smith


Because the receiver was copied at defer time.

---

# But if you use pointer receiver?

Change:


func (p person) fullName()


to:


func (p *person) fullName()


and:


defer p.fullName()


Now `p` is a pointer.

The deferred function holds the address.

So if:


p.firstName = "Alex"


before the function ends, it sees:


Alex Smith


---

# Real backend connection

You will see this pattern often:

### Database transaction


tx, err := db.Begin()

defer tx.Rollback()


Meaning:

> "If this function exits unexpectedly, clean up."

---

### HTTP response


resp, err := client.Do(req)

defer resp.Body.Close()


Meaning:

> "No matter how this function exits, release resources."

---

### Mutex


mutex.Lock()

defer mutex.Unlock()


Meaning:

> "When this function finishes, release the lock."

---

## Interview answer

If asked:

**"What happens when we defer a method call?"**

Say:

> "When a method call is deferred, the receiver and arguments are evaluated immediately, but the execution of the method is postponed until the surrounding function returns. For value receivers, a copy of the receiver is stored. For pointer receivers, the stored pointer refers to the original object."

This concept is very important when you start writing production Go code.

=============================================================================

Good observation. You noticed something subtle.

You are asking:

> "Why did we use `fmt.Printf` instead of `fmt.Println`? Is it because it came in a single line?"

The answer: **Yes, partly — but the main reason is formatting control.**

Let's compare.

---

## 1. `fmt.Println`

Example:


fmt.Println(p.firstName, p.lastName)


Output:


Keerthi shankar


`Println` automatically:

* adds spaces between arguments
* adds a newline (`\n`) at the end

Internally it behaves like:

text
Print values
+
Add newline


So:


fmt.Println("Welcome")
fmt.Println("Keerthi")


Output:


Welcome
Keerthi


---

## 2. `fmt.Printf`

Your code:


fmt.Printf("%s, %s", p.firstName, p.lastName)


Here you are telling Go:

> "Print exactly this format."

The format string:


"%s, %s"


means:


%s → replace with a string


So Go substitutes:


"%s, %s"

      |
      v

"Keerthi, shankar"


Output:


Keerthi, shankar


Notice:

No newline.

---

## Why did your output come in one line?

Because:

### First:


fmt.Printf("Welcome ")


prints:


Welcome 


There is a space at the end.

It does NOT move to the next line.

---

Then deferred function runs:


fmt.Printf("%s, %s", p.firstName, p.lastName)


prints:


Keerthi, shankar


Together:


Welcome Keerthi, shankar


One line.

---

If you used:


fmt.Println("Welcome ")


then:


Welcome

Keerthi, shankar


because `Println` adds a newline.

---

## Difference visually

### Println


fmt.Println("Hello")
fmt.Println("World")


Output:


Hello
World


---

### Printf


fmt.Printf("Hello ")
fmt.Printf("World")


Output:


Hello World


---

## Why developers use Printf?

Mostly for formatting.

Example:

Without Printf:


age := 27

fmt.Println("Age is", age)


Output:


Age is 27


Works.

But:


fmt.Printf("Age is %d years old", age)


Output:


Age is 27 years old


You control the exact sentence.

---

## Common formatting verbs you should know

### String


%s


Example:


fmt.Printf("%s", "Go")


Output:


Go


---

### Integer


%d


Example:


fmt.Printf("%d", 100)


Output:


100


---

### Float


%f


Example:


fmt.Printf("%f", 3.14)


Output:


3.140000


---

### Any value


%v


Very common in debugging:


fmt.Printf("%v", person)


---

## Now connect with your defer example

Your code:


defer p.fullName()

fmt.Printf("Welcome ")


Execution order:

1. `defer` registers the method

2. `Printf` executes immediately:


Welcome 


3. main returns

4. deferred method executes:


Keerthi, shankar


Final:


Welcome Keerthi, shankar


---

### Interview-style explanation:

> "`fmt.Printf` is used when I need formatted output control using placeholders like `%s`, `%d`, etc. Unlike `Println`, it does not automatically add a newline, so multiple Printf calls can continue printing on the same line."

Your observation was correct — the single-line output happened because `Printf` does not add `\n` automatically.
