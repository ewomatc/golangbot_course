package main

import "fmt"

type author struct{
	firstName string
	lastName string
	bio string
}

// create a fullName method() that takes in the author struct as a reciever value
func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}


// blogpost struct
type blogPost struct {
	title string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title:", b.title)
	fmt.Println("Content:", b.content)
	fmt.Println("Author:", b.author.fullName())
	fmt.Println("Bio:", b.author.bio)
}

// imagine we have a website/webpage we want to display several blogs, details. We can create a slice of the blogPosts and embed it in a struct. we will then use a for-range loop to loop through each blogpost and display it
type webpage struct {
	blogPosts []blogPost
}
// a contents method on the webpage struct to display each blogpost on the webpage
func (w webpage) contents() {
	fmt.Println("Webpage content")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Great",
		"Diro",
		"Software Engineer, Web3 Developer",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}

	w := webpage{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
}