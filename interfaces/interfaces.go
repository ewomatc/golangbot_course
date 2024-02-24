// basically an interface is a collection of methods that a type must have, for exammple a collection of methods that a Television struct must have. example turnOn() turnOff() increaseVolume(). etc.
package main

import "fmt"

func main() {
	createBook()

}

// learning interfaces using an inventory example
// we can add different items to the inventory
// this nis the general interface for the items
type Item interface {
	getName() string
	getPrice() int
	displayDetails()
}

// book represents a book in the inventory
type Book struct {
	Title string
	Author string
	Price int
}

// getName method returns the name of the book
func (b Book) getName()string{
	return b.Title
}

func (b Book) getPrice()int {
	return b.Price
}

func (b Book) displayDetails() {
	fmt.Printf("Title: %s \n Author: %s \n $%d \n", b.Title, b.Author, b.Price)
}

// Here, Book is a type that has three methods (GetName(), GetPrice(), and DisplayDetails()) matching the interface Item. Thus, Book implicitly satisfies the Item interface

// using this interface

// create a function to print details of any item in the inventory
func printItemDetails(item Item) {
	item.displayDetails()
	fmt.Printf("Total price: $%d \n", item.getPrice())
}

func createBook() {
	book := Book{
		Title: "The great gatsby",
		Author: "F. Scott Fitzgerald",
		Price: 20,
	}

	// use the printItemdetails(0 func to print this book's details
printItemDetails(book)
}