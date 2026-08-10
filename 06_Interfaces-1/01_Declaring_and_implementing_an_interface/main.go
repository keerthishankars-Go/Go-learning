package main

// We will write a simple program that calculates the total expense for a company based on the salaries of the employees. For brevity, we have assumed that all expenses are in USD.

import "fmt"

type SalaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

//salary of permanent employee is the sum of basic pay and pf
func (p Permanent) calculateSalary() int {
	return p.basicPay + p.pf

}

//salary of contract employee is the basic pay alone
func (c Contract) calculateSalary() int {
	return c.basicPay
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("Total expense per month $%d:", expense)
}

func main() {
	pemp1 := Permanent{
		empId:    1,
		basicPay: 5000,
		pf:       20,
	}

	pemp2 := Permanent{
		empId:    2,
		basicPay: 6000,
		pf:       30,
	}

	cemp1 := Contract{
		empId:    3,
		basicPay: 1000,
	}

	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}

// Line no. 7 of the above program declares the SalaryCalculator interface with a single method CalculateSalary() int.

// In line no. 23 we add the method CalculateSalary() int to the receiver type Permanent. Now Permanent is said to implement the interface SalaryCalculator. This is quite different from other languages like Java where a class has to explicitly state that it implements an interface using the implements keyword. This is not needed in Go and Go interfaces are implemented implicitly if a type contains all the methods declared in the interface.

// We have two kinds of employees in the company, Permanent and Contract defined by structs in line no. 11 and 17. The salary of permanent employees is the sum of the basicpay and pf whereas for contract employees it’s just the basic pay basicpay. This is expressed in the corresponding CalculateSalary methods in line. no 23 and 28 respectively. By declaring this method, both Permanent and Contract structs now implement the SalaryCalculator interface.

// The totalExpense function declared in line no. 36 expresses the beauty of interfaces. This method takes a slice of SalaryCalculator interface []SalaryCalculator as a parameter. In line no. 60 we pass a slice that contains both Permanent and Contract types to the totalExpense function. The totalExpense function calculates the expense by calling the CalculateSalary method of the corresponding type. This is done in line. no 39.

// The biggest advantage of this is that totalExpense can be extended to any new employee type without any code changes. Let’s say the company adds a new type of employee Freelancer with a different salary structure. This Freelancer can just be passed in the slice argument to totalExpense without even a single line of code change to the totalExpense function. This method will do what it’s supposed to do as Freelancer will also implement the SalaryCalculator interface :).