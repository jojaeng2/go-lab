package main

import "fmt"

func defFunc() {
	fmt.Println("This is a deffered function")

	if r := recover(); r != nil {
		fmt.Println("Recovered from panic that is: ", r)
	}
}

func main() {
	defer defFunc()
	panic("Not working!!")
}
