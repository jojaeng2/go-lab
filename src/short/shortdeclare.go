package main

import (
	"fmt"
)

func main() {
	// declare the function
	f := func() {
		fmt.Println("Inside a function")
	}

	// call using the name
	f() // Inside a function
}
