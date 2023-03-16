package main

import (
	"fmt"
)

func main() {
	var x int32
	var y uint32
	var z uint8
	fmt.Println("Type Conversion")
	x = 26700
	y = uint32(x)     // data preserved because number is inside range
	z = uint8(x)      // data loss due to out of range conversion
	fmt.Println(y, z) // prints 26700 76
}
