package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for _, c := range s {
		fmt.Print(string(c), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// run two different goroutine

	go fun("Hello")
	go fun("World")

	time.Sleep(1 * time.Second)
}
