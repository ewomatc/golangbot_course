package main

// defer is used to execute a function call just before the surrounding function returns. note that a function will typically finish executing then return.
// we'll understand this with a simple example

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

// the function takes in a slice of numbers
func largest(nums []int) {
	defer finished()

	maxNum := nums[0]  // initialize maxNums with a value of the zeroth index of the slice
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	fmt.Println("Largest number in", nums, "is", maxNum)
}

// we can also defer a method call, so not only functions
type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	numsSlice1 := []int{78, 109, 2, 563, 300}
	largest(numsSlice1)

	// use the defer on a method
	person1 := person{
		firstName: "John",
		lastName: "Smith",
	}
	defer person1.fullName()
	fmt.Printf("Welcome ")

}