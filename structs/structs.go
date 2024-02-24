package main

import (
	"fmt"
	"golangbot-course/structs/computer"
)

func main() {
	createEmployee()
	anonnymousStruct()
	accessIndividualalues()
	pointerToStruct()
	useNestedStruct()
	useSpec()
}

// general struct syntax
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// the above struct is a named struct because it creates a new data type called employee
// we can make it more compact by declaring same types in a single line
type Employee2 struct {
	firstName, lastName string
	age, salary         int
}

// it is not recomended though since it doesnt mmake each field declaration explicits

// create some emplyees with the Employee struct type
func createEmployee() {
	// creating by specifying field names
	emp1 := Employee{
		firstName: "Sam",
		lastName:  "Brown",
		age:       25,
		salary:    20000,
	}

	// creating without specifying field names
	emp2 := Employee{
		"John", "Doe", 22, 23000,
	}
	// again, this syntax isnt recommended, refrain from using it

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Employee 2:", emp2)
}

// anonymous structs
// we can create a struct withous specifying its name,
func anonnymousStruct() {
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Bustin",
		lastName:  "Jeiber",
		age:       24,
		salary:    21000,
	}

	fmt.Println("Employee 3:", emp3)
}

// accessing individual fields of a struct
func accessIndividualalues() {
	emp6 := Employee{
		firstName: "Dave",
		lastName:  "McKinnon",
		age:       25,
		salary:    22000,
	}

	fmt.Println("Emp6 first name:", emp6.firstName)
	fmt.Println("Emp6 last name:", emp6.lastName)
	fmt.Println("Emp6 age:", emp6.age)
	fmt.Println("Emp6 previous salary:", emp6.salary)
	emp6.salary = 25000
	fmt.Println("Emp6 new salary:", emp6.salary)
}

// creating pointers to structs
func pointerToStruct() {
	// emp8 points to the memory location of the Employee struct
	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       26,
		salary:    26000,
	}

	// Go allows us to access the field values withour explicit dereference 
	// this code below
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)
	// with derefernce would be 
	// (*emp8).firstName
}

// nested structs
type Address struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	address Address // addres is a nested struct inside person struct
}

// how to use nested structs
func useNestedStruct() {
	p := Person{
		name: "Naveen",
		age: 20,
		address: Address{
			city: "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city) // access address and city nested
	fmt.Println("State:", p.address.state)

}

// use the imported the Spec struct from computer dir
func useSpec() {
	spec := computer.Spec{
		Maker: "apple", // access the exported Maker field
		Price: 50000,
	}
	fmt.Println("PC maker:", spec.Maker)
	fmt.Println("PC price:", spec.Price)

}