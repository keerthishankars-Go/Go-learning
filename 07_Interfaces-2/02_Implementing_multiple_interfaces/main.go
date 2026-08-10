package main

//A type can implement more than one interface.

import "fmt"

type salaryCalculator interface {
	DisplaySalary()
}

type leavesCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken

}

func main() {
	e := Employee{
		firstName:   "Keerthi",
		lastName:    "shankar",
		basicPay:    3000000,
		pf:          100000,
		totalLeaves: 35,
		leavesTaken: 12,
	}

	var s salaryCalculator = e
	s.DisplaySalary()

	var l leavesCalculator = e

	fmt.Println("The leaves left =", l.CalculateLeavesLeft())

}

// The program above has two interfaces SalaryCalculator and LeaveCalculator declared in lines 7 and 11 respectively.

// The Employee struct defined in line no. 15 provides implementations for the DisplaySalary method of SalaryCalculator interface in line no. 24 and the CalculateLeavesLeft method of LeaveCalculator interface interface in line no. 28. Now Employee implements both SalaryCalculator and LeaveCalculator interfaces.

// In line no. 41 we assign e to a variable of type SalaryCalculator interface and in line no. 43 we assign the same variable e to a variable of type LeaveCalculator. This is possible since e which of type Employee implements both SalaryCalculator and LeaveCalculator interfaces.

// This program prints,

// Naveen Ramanathan has salary $10200
// Leaves left = 25
