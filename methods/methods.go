package main

import "fmt"

func main() {
	useDisplaySalary()
	useNewNameAndAge()
	useAdd()
}

// A method is just a function with a special receiver type between the func keyword and the method name.

// we are going to create a method to display employee salary, the method uses a type struct called Employee
type Employee struct {
	name string
	salary int
	currency string
	age int
}

// create the display salary method, it has the Employee struct as the reciever type
// e is the reciever name and Employee is the reciever type (a struct)
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

// now we create a function that creates a new emmployee 'emp1' using the Employee struct, and then we use the diaplaySalary() method to display emp1's salary.
func useDisplaySalary() {
	emp1 := Employee{
		name: "Adolf",
		salary: 5000,
		currency: "USD",
	}

	// how to call the displaySalary() method on emp1
	emp1.displaySalary()
}

// remember structs are value types, and we have created methods with value recievers, thjat is structs as the reciever types. 
// however it is very possible to create methods with pointer recievers, that is passing in a pointer as the reciever type to the method.




// method with pointer reciever
func (e *Employee) changeAgeAndName(newAge int, newName string) {
	e.age, e.name = newAge, newName

}

func useNewNameAndAge() {
	// create a new employer with the employer struct
	emp2 := Employee{
		name: "Andrew",
		age: 25,
	}

	fmt.Printf("\nEmployee name and age before change: %s, %d \n", emp2.name, emp2.age)
	emp2.changeAgeAndName(26, "Mark")
	fmt.Printf("Employee name and age after change: %s, %d \n", emp2.name, emp2.age)
}


// when should we use pointer receivers or value recievers
// pointer recievers should bve used when changes made inside the mmethod need to be visible whrever thye method is called. Value recievers won't work inn this case becaiuse of their scope. 
// we use pointer reciever above because the method is modifying the age and name passed to it wherever it is called, and we want that modification to be visible.
// Pointers receivers can also be used in places where it’s expensive to copy a data structure. Consider a struct that has many fields. Using this struct as a value receiver in a method will need the entire struct to be copied which will be expensive. In this case, if a pointer receiver is used, the struct will not be copied and only a pointer to it will be used in the method.
// In all other situations, value receivers can be used.


// Methods with non-struct receivers

//So far we have defined methods only on struct types. It is also possible to define methods on non-struct types, but there is a catch. To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package. So far, all the structs and the methods on structs we defined were all located in the same main package and hence they worked.

// the code below wont work
// func (a int) add(b int) {
// }

// func useAdd() {

// }

// This is not allowed since the definition of the method add and the definition of type int is not in the same package. This program will throw compilation error cannot define new methods on non-local type int

// The way to get this working is to create a type alias for the built-in type int and then create a method with this type alias as the receiver.

type myInt int

func (a myInt) add(b myInt) myInt{
	return a + b
}

func useAdd() {
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)

	fmt.Println("sum is", sum)
}