package main

import (
	"fmt"
	"time"
)

func main() {
	channelDeclaration()
	useHello()
	useHello2()
	
}

// channels are declared with a type, and the onlt data allowed to pass through a channell must be of the type it was declared with
func channelDeclaration() {
	// create channell a of type int, it is emmpty, sero value of a chanell is nil
	var a chan int
	if a == nil{
		fmt.Println("channel a is nil, going to define it")
		// . we use the make() function to define a channel
		a = make(chan int)
		fmt.Printf("type of a is %T \n", a)
	}
}

// sending and recieving from a channel
//data := <- a  // read/recieving from channel a and storing it in variable data
//a <- data // write/send to channel a
// the direction of the arrow in ref to the channel name shows the channel path


// This declares a function called hello that takes in one parameter, a channel of type bool
// In this case, the channel is used to signal that the hello function has finished executing.
func hello(done chan bool){
	fmt.Println("Hello world goroutine")
	// This line sends the value true to the done channel. This is used to signal that the hello function has finished executing.
	done <- true
}

func useHello() {
	// This line creates a new channel of type bool and assigns it to the variable done
	done := make(chan bool)
	// This line starts a new goroutine that runs the hello function with the done channel as its parameter
	go hello(done)
	// This line waits for a value to be received from the done channel. This is used to synchronize the useHello function with the hello function, ensuring that the useHello function does not continue until the hello function has finished executing. In this case it must finish printing the "Hello world goroutine".
	<- done
	fmt.Println("use hello function")
}


// use Sleep to better understand the blocking concept
func hello2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func useHello2() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello2(done)
	<-done
	fmt.Println("Main received data")
}