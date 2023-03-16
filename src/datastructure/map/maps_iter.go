package main

import (
	"fmt"
)

func main() {
	var names = make(map[int]string)

	names[0] = "Freddy" // indexed insertion
	names[1] = "Shawn"
	names[2] = "Batman"
	names[3] = "Spiderman"
	names[4] = "Joker"

	for key, value := range names {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}

}
