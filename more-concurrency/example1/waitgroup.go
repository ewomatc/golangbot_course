package main

import (
	"fmt"
	"sync"
)

// WaitGroup is used to wait for a collection of goroutines to finish. It allows goroutines to signal when they are done so that the main goroutine can proceed.


func main() {
	var wg sync.WaitGroup // create new waitgroup variable

	wg.Add(2)		// add 2 goroutines to the waitgroup, tyhis indicates that 2 goroutines will be spawned

	// we will create 2 goroutines now
	go func() {
		// wg.done() is used to signal the waitgroup that this goroutine has finished executing. wg.done() decrements the counter of the waitgroup (which we indicated was 2) each time a goroutine finishes.
		defer wg.Done()		
		fmt.Println("Goroutine 1 finished executing, over")
	}() //notice that we call the function after writing it

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 finished executing, over")

	}()

	wg.Wait()		// this line blocks the main function from executing until the waitgroup counter becomes 0, meaning that all goroutines have finished executing.

		fmt.Println("All goroutines finished executing, main function, over")

}