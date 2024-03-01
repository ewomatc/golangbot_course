package main

import (
	"fmt"
	"time"
)

func main() {
	useServers()

}

// The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random. 

// example

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func useServers() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	// we'll use select to get the response fromm the quickest channel
	select {
		// read from output1 channel
	case s1 := <- output1:
		fmt.Println(s1)
		// read from output2 channel
	case s2 := <- output2:
		fmt.Println(s2)
	}

	// since server2 is the provider/goroutine for output2 and it only sleeps for 3 seconds, the select syntax will use output2
}