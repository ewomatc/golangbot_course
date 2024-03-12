package main

// Mutex (short for mutual exclusion) is used to synchronize access to shared resources to prevent race conditions.

import (
	"fmt"
	"sync"
	"time"
)

// counter is a shared variable/resource
// sync.Mutex is used to protect access to counter
var (
	counter = 0
	mutex sync.Mutex
)

func increment() {
	mutex.Lock()		// lock tghe mutex to prevent concurrent access
	defer mutex.Unlock()	// unlock the mutex to allow oher goroutines to access it

	counter++
	fmt.Println("Incremented counter to:", counter)
}

func main() {
	// spawn multiple goroutines, in this case 10 goroutines will be spawned and each will increment the counter
	for i := 0; i < 10; i++ {
		go increment()
	}

	time.Sleep(2 * time.Second) // Ensure all goroutines finish before main exits
	fmt.Println("Final counter value:", counter)
}