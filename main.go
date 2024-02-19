package main

import (
	"fmt"
)

func main() {
	var myArray [3]int
	myArray[0] = 12
	myArray[1] = 54
	myArray[2] = 89
	fmt.Println(myArray)
	shorthandArray()
	arrayValueType()
	arrayLength()
	iterateArrayWithRange()
	array2d()

}

// shorthand array declaration
func shorthandArray() {
	myArray2 := [3]int{12, 54, 89}
	fmt.Println(myArray2)
}

/* It is not necessary that all elements in an array have to be assigned a value during short hand declaration.

func main() {
	a := [3]int{12}
	fmt.Println(a)
}

will ouput [12 0 0]
*/

/*
You can ignore the length of the array in the declaration and replace it with ... and let the compiler find the length for you

func main() {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
}
*/

/*
The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized. Don’t worry about this restriction since slices exist to overcome this.

func main() {
	a := [3]int{5, 78, 8}
	var b [5]int
	b = a //not possible since [3]int and [5]int are distinct types
}
*/

//Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array.

func arrayValueType() {
	a := [...]string{"USA", "China", "India", "Japan", "France"} // declare an array a
	b := a                                                       //assign array 'a' to array 'b'
	b[0] = "Turkey"
	fmt.Println("array a is: ", a)
	fmt.Println("array b is: ", b)
}

// finding the length of an array with the len() function
// simply pass the array name to the len() func
func arrayLength() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of array a is: ", len(a))
}

func iterateArrayWithRange() {
	myArray := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)

	// i is index, v is value of the array
	for i, v := range myArray {
		fmt.Printf("%dth index of myArray has a value of %.2f \n", i, v)
		sum = sum + v
	}
	fmt.Println("Sum of all elements in the array is: ", sum)
}

// 2d arrays
func array2d() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Println(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	fmt.Println(b)
}
