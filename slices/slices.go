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
	sliceLenAndCap()
	reslice()
	appendSlice()
	appendNil()
	appendTwoSlices()
	testSubtractTwo()
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

	branch2[0] = 100
	fmt.Println("genesis array after modifying branch 2: ", genesisArray)
}

// the length of a slice is the number of elements in the slice. The capacity of a slice is the number of elements in the base array, starting from the index the slice starts from, find the capacity of a slice a, you'll use the 'cap(a)' function
func sliceLenAndCap() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
}

// A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error. so apparently you can create a slice from a slice.
func reslice() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

// also note, when creating a slice and you dont specify the index1 and indexend values, the indexstart is automatically set to 0, and the index end is the length of the base array.

// appending slices
func appendSlice() {
	cars := []string{"ferrari", "Honda", "ford"} //create a slice of cars
	fmt.Println("cars: ", cars, "previous length: ", len(cars), "previous capacity: ", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars: ", cars, "new length: ", len(cars), "new capacity: ", cap(cars))
}

//in the above program, the capacity of cars is 3 initially. We append a new element to cars and assign the slice returned by append(cars, "Toyota") to cars again. Now the capacity of cars is doubled and becomes 6. The output of the above program is

// The zero value of a slice type is nil. A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.

func appendNil() {
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
}

// It is also possible to append one slice to another using the ... operator. You can learn more about this operator in the variadic functions tutorial.

func appendTwoSlices() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
	fmt.Println(cap(food))
}

// passing slices to functions
func subtractTwo(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func testSubtractTwo() {
	nums := []int{8, 7, 6}
	fmt.Println("nums slice befor calling subtractwo on it: ", nums)
	subtractTwo(nums)
	fmt.Println("nums slice after passing it into the subtracttwo function: ", nums)
}
