package main

import (
	"fmt"
	"sync"
)

// Fan-Out is when multiple functions read from the same channel until that channel is closed. Fan-In is when multiple goroutines write to the same channel until that channel is closed.

func producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)

	go producer(ch)

	for i := 0; i < 3; i++ {
		go consumer(ch, &wg)
	}

	wg.Wait()
}
