package main

import (
	"fmt"
	"time"
)

func main() {
	useWriteToBuffer()
	closingBufferedChannels()

}

// Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

// ch := make(chan type, capacity)

// capacity in the above syntax should be greater than 0 for a channel to have a buffer. The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating unbuffered channels

// an example to understand bocking in buffered channels


//  we create a write() goroutine that should write values 0 to 4 into a xhannel
func writeToBuffer(ch chan int) {
	for i := 0; i <5;i++{
		ch <- i
		fmt.Printf("successfuly wrote %d to ch \n", i)
	}
	close(ch)
}

func useWriteToBuffer() {
	// create a channel
	myChannel := make(chan int, 2)
	go writeToBuffer(myChannel)
	// the above will only write two values which are 0 and 1, because this myChannel has a capacity of 2
	time.Sleep(2 * time.Second)
	for value := range myChannel {
		fmt.Printf("read value %d from ch \n", value)
		time.Sleep(2 * time.Second)
	}

	// now the channel will automatically block once this buffers capacity of 2 is reached.
	// it will then have to wait for a value to be read from the buffer, then resume writing into the buffer since there will now be space init. 
}


// DEADLOCK
/*
	note: when creating/writing to buffered chanels, there must be a goroutine running concurrently and reading data from it. If there isn't, once the capacity is reached, there will be a deadlock and the program will panic at runtime. 
*/

// CLOSING BUFFERED CHANNELS
/*
When we close buffered channels, we can still read data from them 'for as long as' there is still data in them. Once there is no more data in the channel, it will  return the zero value of whatever data we were trying to read
*/
// example

func closingBufferedChannels() {
	ch := make(chan int, 5)
	ch <- 7
	ch <- 8
	close(ch)
	// close the channel after writing 2 datas to it, it has a capacity of 5
	// use a for-range to read data from it
	for data := range ch{
		fmt.Printf("recieved data %d \n", data)
	}
}

// LENGTH VS CAPACITY OF A BUFFER
/*
the length of a buffer is the number of elements currently written to it, 
while the capacity is the maximum amount of data it can hold at once
*/

