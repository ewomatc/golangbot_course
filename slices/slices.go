package main

import "fmt"

// general syntax for a slice is []T where T is the type

func main() {
	a := [5]int{76, 77, 78, 79, 80} //create an array 'a'
	var b []int = a[1:4]            // creates a slice of array 'a' from index 1 to index 3
	fmt.Println(a)
	fmt.Println(b) //will print [77, 78, 79]
	sliceSyntax2()
	modifyingSlice()
	modifyingSlice2()
}

// second way to create a slice
func sliceSyntax2() {
	mySlice := []int{77, 78, 79}
	fmt.Println(mySlice)
}

// a slice does not own any data of its own, it is just a representation of the underlying array.
// any change made to the slice will reflect in/affect the original array

func modifyingSlice() {
	darr := [...]int{57, 68, 79, 81, 92, 103, 114, 125}
	dslice := darr[2:5]
	fmt.Println("array before slicing and modifying: ", darr)

	// modify the slice
	for i := range dslice {
		dslice[i]++ // increment the slice, and it will reflect in the original array
	}
	fmt.Println("array after slicing and modifying from index 2 to index 4; ", darr)
}

// if i have 3 slices for example, that bot have the same underlying array, the change that each of them make will be reflected in the base array

func modifyingSlice2() {
	genesisArray := [3]int{78, 79, 80}

	branch1 := genesisArray[:] // this '[:]' indicates that the slice should contain all ements of the specified array
	branch2 := genesisArray[:]

	fmt.Println("genesis array before any modification: ", genesisArray)
	branch1[0] = 90
	fmt.Println("genesis array after modifying branch 1: ", genesisArray)

	branch2[1] = 100
	fmt.Println("genesis array after modifying branch 2: ", genesisArray)
}
