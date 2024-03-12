package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f,f.Name(), "opened successfully")
}

// if a function or method in Go returns an error, then the error has to be the last value returned.
// The idiomatic way of handling errors in Go is to compare the returned error to nil.
// A nil value indicates that no error has occurred and a non-nil value indicates the presence of an error