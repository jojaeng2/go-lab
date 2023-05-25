package main

import "fmt"

func main() {

	v := func() {
		fmt.Println("Inside anonymous function")
	}

	v()

	// func() {
	// 	fmt.Println("Inside anonymous function")
	// }()
}
