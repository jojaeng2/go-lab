package main

import (
	"fmt"
)

type Student struct {
	name string
}

func main() {
	// declare
	var students []Student

	s := Student{"JO"}
	h := Student{"HO"}
	jo := append(students, s)
	ho := append(students, h)
	fmt.Println(jo)
	fmt.Println(ho)
}
