The second way to get more information from the error is to find out the underlying type and get more information by calling methods on the struct type.

Let’s understand this better by means of an example.

The DNSError struct type in the standard library is defined as follows,

  
type DNSError struct {
    ...
}
  
func (e *DNSError) Error() string {
    ...
}
func (e *DNSError) Timeout() bool { 
    ... 
}
func (e *DNSError) Temporary() bool { 
    ... 
}

The DNSError struct has two methods Timeout() bool and Temporary() bool which return a boolean value that indicates whether the error is because of a timeout or is it a temporary one.