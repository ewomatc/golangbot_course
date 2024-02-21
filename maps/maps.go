package main

import "fmt"

func main() {
	employeeSalaryMap()
	employeeSalaryMapInitialized()
	retrieveMapValue()
	verifyMapKey()
	iterateMap()
	deleteMapItem()
	mapOfStruct()
	//mapEquality()

}

// a map is a data structure that simply stores key and value pairs
// the Go syntax for maps is:
// make(map[type of key] type of value)
// example, a map called employeeSalary that stores the employee's first name as the key and their salary as the value
func employeeSalaryMap() {
	employeeSalary := make(map[string]int)

	// The syntax for adding new items to a map is the same as that of arrays. The program below adds some new employees to the employeeSalary map.

	employeeSalary["Steve"] = 12000
	employeeSalary["Jammie"] = 15000
	employeeSalary["Mike"] = 10000
	fmt.Println(employeeSalary)
}

// we can also initializr the map when declaring it
func employeeSalaryMapInitialized() {
	employeeSalary := map[string]int{
		"steve":  12000,
		"jammie": 15000,
	}
	employeeSalary["mike"] = 10000
	fmt.Println("Employee salary initialized during declaration", employeeSalary)
}

// retrieving elements from a map
// the general syntax is: map[key]
func retrieveMapValue() {
	employeeSalary := map[string]int{
		"steve":  12000,
		"jammie": 15000,
		"mike":   10000,
	}

	employee := "jammie"               // assign jammie to an employee var
	salary := employeeSalary[employee] // pass the employee to a map retrieve syntax
	fmt.Printf("salary of %s is %d \n", employee, salary)
}

// note: if we try to retrieve the value for a key thats not in the map, fir example, we try to retrieve the salary of "Joe", it wont throw an error but return the zero value which is '0'

// how to check if a key exists
// the general syntax is value, ok := map[key]
func verifyMapKey() {
	employeeSalary := map[string]int{
		"steve":  12000,
		"jammie": 15000,
		"mike":   10000,
	}
	newEmployee := "Joe"
	value, ok := employeeSalary[newEmployee]
	if ok {
		fmt.Printf("Salary of %s is %d", newEmployee, value)
		return // we return to stop the execution
	}
	fmt.Printf("Employee %s was not found", newEmployee)

}

// iterating through a map
// the range from for loop is used to iterate through a map

func iterateMap() {
	employeeSalary := map[string]int{
		"steve":  12000,
		"jammie": 15000,
		"mike":   10000,
	}
	fmt.Println("Employees salary")
	for key, value := range employeeSalary {
		fmt.Printf("Employee: %s ******* Salary: %d \n", key, value)
	}
}

// deleting items from an array
// general syntax: delete(map, key)

func deleteMapItem() {
	employeeSalary := map[string]int{
		"steve":  12000,
		"jammie": 15000,
		"mike":   10000,
	}
	fmt.Println("map before deleting; ", employeeSalary)
	delete(employeeSalary, "mike")
	fmt.Println("map after deleting: ", employeeSalary)
}

// map of structs
// what if we want to also store the employees country in the maps too, we'll need an extra field. we can then make the employee a type struct and give it the fields of salary and country. Now the map stores the key value as emplyee namme to the employee struct.
type employee struct {
	salary  int
	country string
}

func mapOfStruct() {
	emp1 := employee{
		salary:  12000,
		country: "USA",
	}
	emp2 := employee{
		salary:  15000,
		country: "Austrailia",
	}
	emp3 := employee{
		salary:  10000,
		country: "Canada",
	}

	// create a map with key of string and value of the employee struct
	employeeInfo := map[string]employee{
		"steve": emp1,
		"jamie": emp2,
		"mike":  emp3,
	}

	// loop the emplyeeInfo map with the name and info
	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s ******* salary: %d ****** country: %s \n", name, info.salary, info.country)
	}

}

//  length of a map
// same as arrays and slices, the length of a map can be determined using he len() function.

// reference types
// just like slices, maps are reference types to the same undelying data structure, so if we assign a map (mapA)to a new variable say mapB, any change made to mapB will reflect in mapA

// maps equality
// you cant check the equality of a map using ==, you can only check with == if a map is nill
// func mapEquality() {
// 	map1 := map[string]int{
// 		"one": 1,
// 		"two": 2,
// 	}

// 	map2 := map1

// 	// the below line will not compile, it'll result in an error
// 	if map1 == map2 {

// 	}
// }

// thats all on maps :)
