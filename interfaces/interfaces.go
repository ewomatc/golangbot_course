// basically an interface is a collection of methods that a type must have, for exammple a collection of methods that a Television struct must have. example turnOn() turnOff() increaseVolume(). etc.
package main

import (
	"fmt"
)

func main() {
	createBook()

}

// learning interfaces using an inventory example
// we can add different items to the inventory
// this nis the general interface for the items
type Item interface {
	getName() string
	getPrice() int
	displayDetails()
}

// book represents a book in the inventory
type Book struct {
	Title string
	Author string
	Price int
}

// getName method returns the name of the book
func (b Book) getName()string{
	return b.Title
}

func (b Book) getPrice()int {
	return b.Price
}

func (b Book) displayDetails() {
	fmt.Printf("Title: %s \n Author: %s \n $%d \n", b.Title, b.Author, b.Price)
}

// Here, Book is a type that has three methods (GetName(), GetPrice(), and DisplayDetails()) matching the interface Item. Thus, Book implicitly satisfies the Item interface

// using this interface

// create a function to print details of any item in the inventory
func printItemDetails(item Item) {
	item.displayDetails()
	fmt.Printf("Total price: $%d \n", item.getPrice())
}

func createBook() {
	book := Book{
		Title: "The great gatsby",
		Author: "F. Scott Fitzgerald",
		Price: 20,
	}

	// use the printItemdetails(0 func to print this book's details
printItemDetails(book)
}

// practical use case 2
// We will write a simple program that calculates the total expense for a company based on the individual salaries of the employees

// define salarycalculator interface
type SalaryCalculator interface {
	CalculateSalary()
}

// define permanent emplyment and contract struct types
type Permanent struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

// create the calculatesalary method for contracts and permanent types.
// Salary of permanent employee is the basic pay plus pf
func (p Permanent) CalculateSalary() int{
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary()int{
	return c.basicpay
}

// we will use range to loop through the salarycalculator slice, and add the salaries of individual employees together, this way, we'll get the total expenses.
func totalExpenses(s []SalaryCalculator) {
	// initialize expense to a value of 0
	expense := 0
	// loop through s, which is salarycalculator slice, we use _ instead of index since we wont need to use the index, and v is the value.
	for _, value := range s {
		expense = expense + value.CalculateSalary()
	}
	fmt.Printf("Total expense per month: $%d \n", expense)	
}

// use the total expense function
func useTotalExpense() {
	// pemmp - permanent employee, cemp, contract employee
	pemp1 := Permanent{
		empId: 1,
		basicpay: 5000,
		pf: 20,
	}
	pemp2 := Permanent{
		empId: 2,
		basicpay: 6000,
		pf: 30,
	}
	cemp1 := Contract{
		empId: 3,
		basicpay: 3000,
	}
	// create a slice that contains these emplpoyees, and assign them to a variable employees.
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpenses(employees)
}

// The biggest advantage of this is that totalExpense can be extended to any new employee type without any code changes. Let’s say the company adds a new type of employee Freelancer with a different salary structure. This Freelancer can just be passed in the slice argument to totalExpense without even a single line of code change to the totalExpense function. This method will do what it’s supposed to do as Freelancer will also implement the SalaryCalculator interface :).
