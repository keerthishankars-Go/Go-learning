1. Converting the error to the underlying type and retrieving more information from the struct fields
If you read the documentation of the Open function carefully, you can see that it returns an error of type *PathError. PathError is a struct type and its implementation in the standard library is as follows,

type PathError struct {
    Op   string
    Path string
    Err  error
}
  
func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
In case you are interested to know where the above source code exists, it can be found here https://cs.opensource.google/go/go/+/refs/tags/go1.19:src/io/fs/fs.go;l=250

From the above code, you can understand that *PathError implements the error interface by declaring the Error() string method. This method concatenates the operation, path, and the actual error and returns it. Thus we got the error message,

open /test.txt: No such file or directory
The Path field of PathError struct contains the path of the file which caused the error.

We can use the As function from errors package to convert the error to its underlying type. The As function’s description talks about error chain.  A simple description of As is that it tries to convert the error to a error type and returns either true or false indicating whether the conversion is successful or not.

==============================================================================

This topic (`errors.As`) feels confusing because it introduces a new idea:

> **An error is not always just a message. It can be a specific type containing extra information.**

Let's build the mental model slowly.

---

## 1. Normal error handling you already know

Example:


f, err := os.Open("test.txt")

if err != nil {
    fmt.Println(err)
    return
}


If file does not exist:

Output:


open test.txt: no such file or directory


You only see a message.

But internally, the error is not just a string.

---

# 2. What actually comes from os.Open?

When this happens:


f, err := os.Open("test.txt")


and file is missing, Go creates a special error:

text
os.PathError


Something like:


PathError
|
+-- Op: "open"
|
+-- Path: "test.txt"
|
+-- Err: "file not found"


The error contains structured data.

---

So internally:


err


is not:


"file not found"


It is:


*os.PathError{
    Op: "open",
    Path:"test.txt",
    Err: ...
}


---

# 3. Why do we need errors.As?

Suppose:


fmt.Println(err)


You get:


open test.txt: no such file or directory


But you want:


Which file failed?


You need access to:


Path


like:


err.Path ❌


You cannot do that because:


err


has type:


error


not:


os.PathError


---

# 4. Remember error is an interface

In Go:


type error interface {
    Error() string
}


Meaning:

Any type that has:


Error() string


can become an error.

Example:


type MyError struct {

}

func (m MyError) Error() string {
    return "something wrong"
}


Now:


var err error = MyError{}


works.

---

So this:


err


is a box.

Inside that box may be:


error interface

+----------------+
| Error() method |
|                |
| actual value   |
| *PathError     |
+----------------+


---

# 5. Now understand errors.As

Code:


var pErr *os.PathError

errors.As(err, &pErr)


Meaning:

> "Hey Go, check whether this error contains a *os.PathError. If yes, give me access to it."

---

Let's visualize.

Before:


err
 |
 |
 v

*os.PathError{
    Path:"test.txt"
}


You create:


var pErr *os.PathError


Currently:


pErr = nil


Then:


errors.As(err, &pErr)


Go does:


Can I convert err into *os.PathError?

YES

Put it here:

pErr
 |
 v

*os.PathError{
 Path:"test.txt"
}


Now:


pErr.Path


works.

Output:


test.txt


---

# 6. Why do we pass `&pErr`?

This is the most confusing part.

You wrote:


errors.As(err, &pErr)


Why not:


errors.As(err, pErr)


?

Because `As` needs to modify `pErr`.

Remember:

Before:


pErr = nil


After:


pErr = actual PathError


To change a variable inside another function, Go needs its address.

Example:


func change(x *int){
    *x = 10
}

func main(){

    a:=5

    change(&a)

    fmt.Println(a)
}


Output:


10


Same idea.

---

# 7. Why is pErr a pointer?

The tutorial mentions:


func (e *PathError) Error() string


Notice:


*PathError


not:


PathError


Meaning:

Only the pointer implements error.

So:

This works:


var err error = &os.PathError{}


This does not:


var err error = os.PathError{}


because the method belongs to:


*PathError


Therefore:


var pErr *os.PathError


---

# 8. Full execution flow

Your program:


f, err := os.Open("test.txt")


File missing:


err = *os.PathError


---

Then:


if err != nil


true.

---

Create variable:


var pErr *os.PathError


Memory:


pErr = nil


---

Try conversion:


errors.As(err,&pErr)


Go checks:


Is err a *os.PathError?

YES


Now:


pErr
 |
 v

PathError{
 Path:"test.txt"
}


---

Then:


fmt.Println(pErr.Path)


Output:


Failed to open file at path test.txt


---

# 9. Difference between normal type assertion and errors.As

You may know:


value, ok := err.(*os.PathError)


This is a type assertion.

Example:


pErr, ok := err.(*os.PathError)

if ok {
    fmt.Println(pErr.Path)
}


It works.

But `errors.As` is preferred in production because it also works with **wrapped errors**.

Example later:


database error
      |
      |
payment error
      |
      |
PathError


`errors.As` can search inside the chain.

---

# 10. When will you use errors.As in real backend?

Example:

Database:


if err != nil {

    var pqErr *pq.Error

    if errors.As(err,&pqErr){

        if pqErr.Code=="23505" {
            // duplicate key
        }
    }
}


HTTP:


var netErr net.Error

if errors.As(err,&netErr){

    if netErr.Timeout(){
        retry()
    }
}


---

# The mental model to remember

Think:


error interface
       |
       |
       v
actual error type
       |
       |
       +---- PathError
       |
       +---- TimeoutError
       |
       +---- DatabaseError


`errors.As` asks:

> "Can I extract this specific type from the error?"

---

## Interview explanation

> "`errors.As` is used to check whether an error is of a specific type and extract that underlying error value. It is useful when we need additional information from typed errors. The second argument is a pointer because As needs to populate that variable with the matched error value."

---

Your learning sequence is actually moving into **real production Go now**:


panic/recover
      ↓
error interface
      ↓
custom errors
      ↓
errors.Is
      ↓
errors.As
      ↓
error wrapping


These are exactly the concepts you will see in mature Go backend code.
