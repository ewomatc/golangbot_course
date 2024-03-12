package main

import (
	"fmt"
)

// first class functions allow functions to be assigned to variables, passed as arguments to other functions, and returned from other functions

func main() {
	a := func() {
		fmt.Println("hello world first class functions")
	}

	a()
	fmt.Printf("Type of a is %T \n", a)

	useAddNums()
	useAddNums2()
}

// PASSING FUNCTIONS AS ARGUMENTS TO OTHER FUNCTIONS

// we create a function addNums() that takes in a function 'a' that takes in 2 ints as arguments and returns an int. 
func addNums(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func useAddNums() {
	// create an anonymous function 'f' whose signature matches the function 'a' that addNums() is recieving
	f := func(a, b int) int {
		return a + b
	}
	// pass the function 'f' to addNums()
	addNums(f)
}

// RETURNING FUNCTIONS FROM OTHER FUNCTIONS

// lets rewrite the program above to return a function from the addNums() function instead

// addNums2() returns a function that takes in 2 parameters a and b
func addNumms2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}


func useAddNums2() {
	s := addNumms2()
	// s contains the function returned by addNums2(). so when we print s below we pass  in 2 arguments and the addNNums2 function does its work.
	fmt.Println(s(60, 7))
}
