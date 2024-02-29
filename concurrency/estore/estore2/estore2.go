package main

import (
	"fmt"
	"time"
)

// This code demonstrates a simple example of concurrent programming in Go using goroutines and channels. The main function acts as the coordinator, with one goroutine generating orders and another goroutine processing them concurrently. This pattern allows for efficient handling of tasks in parallel, improving the overall performance of the program.

type Order struct{
	ID int
	Amount int
}
func main() {
	orders := make(chan Order)

	// order/goods producer goroutine
	go func() {
		// ensure that the orders channel is closed after all orders are sent
		defer close(orders)
		// use a for loop to create 3 orders and send each of them into the orders channel with '<-'
		for i := 1; i <= 3; i++ {
			orders <- Order{ID: i, Amount: int(i * 10)}
	}
	}()

	// use a for-range loop to loop through each order as they are recieved from the orders channel, we print a message for each successful order
	for order := range orders{
		fmt.Printf("Processing order ID %d, Amount: $%.d \n", order.ID, order.Amount)
		time.Sleep(3 * time.Second)
		fmt.Printf("Order %d sent successfully \n", order.ID)
	}
}