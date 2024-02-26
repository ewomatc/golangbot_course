package main

import (
	"fmt"
	"time"
)

// to create a goroutine, simply prefix the method or functioon call with the 'go' keyword
func hello() {
	fmt.Println("Hello world goroutine")
}

func main()  {
	go hello()
	// we use the time.sleep to sleep (await) on the main goroutine for 1 second so that the hello() go routine can fiunish executing.
	time.Sleep(1 * time.Second)
	go alphabets()
	go numbers()
	// wait 3000 milliseconds for these above 2 functions to finish execyting before exiting the program
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main function")
}

// note that in a normal way, channels are used instead of the time.sleep() hack.

// we can use mmultiple goroutines like so.
func numbers() {
	for i := 1; i <= 5; i++ {
		// this function prints numbers from 1 to 5 with a 250 millisecond interval
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		// this prings letters a to  e with a 400 miliseconds interval
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

