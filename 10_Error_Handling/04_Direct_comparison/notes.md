The third way to get more details about an error is the direct comparison with a variable of type error. Let’s understand this by means of an example.

The Glob function of the filepath package is used to return the names of all files that matches a pattern. This function returns an error ErrBadPattern when the pattern is malformed.

**ErrBadPattern is defined in the filepath package as a global variable.**

**var ErrBadPattern = errors.New("syntax error in pattern")**

errors.New() is used to create a new error. We will discuss this in detail in the next tutorial.

ErrBadPattern is returned by the Glob function when the pattern is malformed.
