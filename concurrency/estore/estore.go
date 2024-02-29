package main

import (
	"fmt"
	"time"
)

// an ecommerce order processing example to understand chanels better

type Order struct{
	ID int
	Amount int
}

func PaymentVerification(order Order, paymentCh chan <- Order) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Payment verified for order %d\n", order.ID)
	paymentCh <- order
}

func InventoryUpdate(order Order, inventoryCh chan <- Order) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Inventory updated for order %d\n", order.ID)
	inventoryCh <- order
}

func WarehouseShipping(order Order) {
	// Simulate warehouse shipping process
	time.Sleep(3 * time.Second)
	fmt.Printf("Products shipped for order %d\n", order.ID)
}

func main() {
	// create an orders slice.
	orders := []Order{
		{ID: 1, Amount: 50},
		{ID: 2, Amount: 30},
		{ID: 3, Amount: 70},
	}


	// we create these channels that will be used to communicate between the paymentverification goroutine and inventoryupdate goroutine
	// these channels have a type struct, the struct being Order
	paymentCh := make(chan Order)
	InventoryCh := make(chan Order)

	// process orders automatically

	// this loop iterates over the orders slice.
	// it starts a new PaymentVerification() goroutine for each order, taking in the current order and paymentCh channel as argumments.
	// this goroutine will print payment verified for order 'the order'
	for _, order := range orders{
		go PaymentVerification(order, paymentCh)
	}


	// this other loop iterates through the orders slice, and for each order, it recieves the orders whose payments have been verified one by one from the paymentch channel.
	// for each verified order, it starts a new goroutine, inventoryupdate() which takes the inventorych channel as an argument.
	// the inventoryupdate goroutine will update each order in the innventory concurrently.
	for range orders{
		order := <-paymentCh
		go InventoryUpdate(order, InventoryCh)
	}


	// just like the previous loop, this loop recieves data(the updated inventory) from the inventory channel one by one(for each order), and stores it in an order variable. 
	// each updated order is then passed to the warehouseshippinng function, and voila, each order has been shipped :).
	for range orders{
		order := <- InventoryCh
		WarehouseShipping(order)
	}

	fmt.Println("All orders processed successfully")
}

