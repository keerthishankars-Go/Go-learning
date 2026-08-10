**Providing more information about the error using struct type and fields**

It is also possible to use struct types which implement the error interface as errors. This gives us more flexibility with error handling. In our previous example, if we want to access the radius which caused the error, the only way now is to parse the error description Area calculation failed, radius -20.00 is less than zero. This is not a proper way to do this since if the description changes, our code will break.

We will use the strategy followed by the standard library explained in the previous tutorial under the section “Converting the error to the underlying type and retrieving more information from the struct fields” and use struct fields to provide access to the radius which caused the error. We will create a struct type that implements the error interface and use its fields to provide more information about the error.

The first step would be create a struct type to represent the error. The naming convention for error types is that the name should end with the text Error. So let’s name our struct type as areaError

type areaError struct {
	err    string
	radius float64
}
The above struct type has a field radius which stores the value of the radius responsible for the error and err field stores the actual error message.

The next step is to implement the error interface.

func (e *areaError) Error() string {  
    return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}
In the above snippet, we implement the Error() string method of the error interface using a pointer receiver *areaError. This method prints the radius and the error description.