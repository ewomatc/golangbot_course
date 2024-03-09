package main

import (
	"fmt"
	"time"
)

// select is used to wait on multiple channel operations. It blocks until one of its cases can proceed, then executes that case.



func main() {

	//we create 2 channels which we'll write data into, then use 'select' to read data from.
	ch1 := make(chan string)
	ch2 := make(chan string)

	// we then create 2 channels, each write data into their respective channels after somme time
	// select is used yo read data from the channel that delivers first
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "Hello from channel 2"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		ch1 <- "Hello from channel 1"
	}()

	select{
		// ch1 send the data to the msg1 variable
	case msg1 := <- ch1:
		fmt.Println("We recieved this from channel 1:", msg1)
	case msg2 := <- ch2:
		fmt.Println("We recieved this from chanel 2:", msg2)
	}
}

// this program will proint the message from channel 2 because we got that one first/faster, this is determined by the select syntax.