package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golangbot123.com")

	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			if dnsErr.Timeout() {
				fmt.Println("Operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}

// Let’s write a program that converts the error to *DNSError type and calls the above mentioned methods to determine whether the error is temporary or due to timeout.

// In the program above, in line no. 9, we are trying to get the IP address of an invalid domain name golangbot123.com. In line no. 13 we get the underlying value of the error by using the As function and converting it to *net.DNSError. Then we check whether the error is due to timeout or is temporary in line nos. 14 and 18 respectively.

// In our case, the error is neither temporary nor due to timeout and hence the program will print,

// Generic DNS error lookup golangbot123.com: no such host
// If the error was temporary or due to a timeout, then the corresponding if statement would have executed and we can handle it appropriately.
