package main

import (
	"fmt"
)

func calculateBill(price, units int) int {
	var totalPrice = price * units
	return totalPrice
}

func main() {
	price, units := 90, 6
	totalPrice := calculateBill(price, units)
	fmt.Println("Total price is", totalPrice)
}
