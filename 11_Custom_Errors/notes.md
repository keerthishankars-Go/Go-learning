Creating custom errors using the New function
The simplest way to create a custom error is to use the New function of the errors package.

Before we use the New function to create a custom error, let’s understand how it is implemented. The implementation of the New function in the errors package is provided below.

package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
        return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
        s string
}

func (e *errorString) Error() string {
        return e.s
}

The implementation is pretty simple. errorString is a struct type with a single string field s. The Error() string method of the error interface is implemented using a errorString pointer receiver in line no. 14.

The New function in line no. 5 takes a string parameter, creates a value of type errorString using that parameter and returns the address of it. Thus a new error is created and returned.

Now that we know how the New function works, lets use it in a program of our own to create a custom error.

==============================================================================

This is a very important Go pattern: **creating and returning your own errors**.

You are now moving from:


os.Open()


where Go gives you an error,

to:

> "I am designing a function. If something is wrong, I need to create and return my own error."

Let's understand the thinking.

---

# 1. What is the problem?

You have a function:


func circleArea(radius float64) (float64, error)


Meaning:

This function has **two possible outcomes**.

### Success:

Return area.

Example:

text
radius = 10

area = 314.15


Return:


(area, nil)


---

### Failure:

Radius cannot be negative.

Example:

text
radius = -20


Return:


(0, error)


---

# 2. Why return two values?

This is a Go design pattern.

Many functions return:


value, error


Example:


file, err := os.Open()


Same pattern.

Meaning:

text
Did it work?
       |
       |
       +---- yes → use value
       |
       +---- no → handle error


---

# 3. Function declaration


func circleArea(radius float64) (float64, error)


Let's break syntax.

Normal function:


func add(a int, b int) int


means:

Input:

text
two integers


Output:

text
one integer


---

Your function:


func circleArea(radius float64) (float64, error)


Input:

text
radius → float64


Output:

text
1st value → float64
2nd value → error


So Go expects:


return something, something


---

# 4. Inside function


if radius < 0 {


Business rule:

text
A circle radius cannot be negative.


So we stop.

---

# 5. Creating an error


errors.New("Area calculation failed, radius is less than zero")


This creates an error object.

Think:

Before:

text
error = nothing


After:

text
error
 |
 v
"Area calculation failed..."


---

So this:


return 0, errors.New("Area calculation failed, radius is less than zero")


means:

Return:


area = 0

error = something went wrong


Why return 0?

Because area calculation failed.

---

# 6. Success case

If radius is positive:


return math.Pi * radius * radius, nil


Example:

radius:

text
10


Calculation:

text
3.14159 * 10 * 10

=314.159


Return:


314.159, nil


Meaning:


area = 314.159

error = nothing


---

# 7. Main function


radius := -20.0


You intentionally gave invalid input.

---

Call:


area, err := circleArea(radius)


The function returns:

text
area = 0

err = "Area calculation failed..."


Memory:


area
 |
 0


err
 |
 Error object


---

# 8. Error checking


if err != nil {


This means:

> Did the function fail?

Remember:

Success:


err == nil


Failure:


err != nil


---

Since radius is negative:


err != nil


is true.

---

Print:


fmt.Println(err)


Output:


Area calculation failed, radius is less than zero


Then:


return


Stop execution.

---

# Complete flow


main()

 |
 |
 v

circleArea(-20)

 |
 |
 v

radius < 0 ?

YES

 |
 |
 v

return 0, error


 |
 |
 v

main receives:

area = 0
err = error


 |
 |
 v

if err != nil

 |
 |
 v

print error



---

# Now change radius


radius := 20.0


Flow:


circleArea(20)

radius < 0 ?

NO

calculate area

return:

1256.63, nil


main:

err == nil

print:

Area of circle 1256.64


---

# Important Go error pattern

You will see this everywhere:


result, err := someFunction()

if err != nil {
    return err
}

// continue using result


Example:

Database:


user, err := repo.GetUser(id)

if err != nil {
    return err
}

fmt.Println(user.Name)


HTTP:


resp, err := client.Do(req)

if err != nil {
    return err
}


File:


file, err := os.Open("data.txt")

if err != nil {
    return err
}


---

# Why not use exceptions like Java/Python?

Go intentionally avoids:

java
try {
   something()
}
catch(Exception e){

}


Instead:


value, err := something()

if err != nil {

}


Why?

Because errors are explicit.

When you read Go code:


data, err := fetchData()


you immediately know:

> "This operation can fail."

---

# Interview explanation

If asked:

**"How does Go handle errors?"**

Say:

> "Go treats errors as values. Functions commonly return the result and an error value. The caller checks whether the error is nil. If it is nil, the operation succeeded; otherwise, the caller handles the failure. We create custom errors using errors.New or fmt.Errorf."

---

Your learning progression is now exactly the production path:


defer
  ↓
error values
  ↓
custom errors
  ↓
errors.Is
  ↓
errors.As
  ↓
error wrapping
  ↓
service-level error handling


This is the foundation behind the error handling you see in real Go backend code.
