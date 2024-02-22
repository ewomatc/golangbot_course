package main

import (
	"fmt"
)

// A pointer is a variable that stores the memory address of another variable.

func main() {
	pointer()
	zeroValue()
	dereferencePointer()
	pointerFunction()
	callHello()
	useModify()
}

func pointer() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T \n", a)
	fmt.Println("address of b is", a)
}

// the & synbol iis used to get the memory address of the variable b, since b has a type int, a has *int. a stores the memory address of b, so a is a pointer to b

// zero value of a pointer is nil
func zeroValue() {
	a := 25
	var b *int

	if b == nil {
		fmt.Println("b before pointing to a is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

// dereferencing a pointer
// this mmeans accessing the value of the 'pointed' variable :)
// general syntax is *pointer '*' before the pointer

func dereferencePointer() {
	b := 255
	a := &b
	fmt.Println("Memory address of b is", a)
	fmt.Println("Value of b is", *a)
	// we can also modify the value of the pointed
	*a++
	fmt.Println("new value of b is", b)
}

// passing pointer to a function
func change(value *int) {
	*value = 55 // this function takes in a pointer, and then uses dereferencing to change the value of tge pointed to 55
}

func pointerFunction() {
	a := 58
	fmt.Println("a before change()", a)
	b := &a
	change(b)
	fmt.Println("a after change()", a)
}

// function retuirning a pointer to a local variable
func hello() *int {
	i := 5    // local variable
	return &i // return a pointer to local variable 'i'
}

func callHello() {
	d := hello()
	fmt.Println("value of d", d)
}

// we will create a function that takes in a slice and uses a pointer to modify the first item of the slice.
func modify(sls []int) {
	sls[0] = 90 // modify the 0th index of the passed slice to the number 90
}

func useModify() {
	a := [3]int{89, 90, 91} // create an array 'a' and pass in 3 numbers
	modify(a[:])            // pass a slice that refernces the entire array 'a' to the modify func
	// it will change the first item of the a slice to 90, which will then reflect in the underlying array
	fmt.Println(a)
}
