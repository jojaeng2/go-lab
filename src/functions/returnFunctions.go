package main

import (
	"fmt"
)

func main() {
	f := getFunction("Ho")

	f()
}

func getFunction(name string) func() {
	return func() {
		fmt.Println("hello ", name)
	}
}
