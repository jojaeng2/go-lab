package main

import "fmt"

func main() {
	arr := [2][2]int{
		{1, 2},
		{3, 4}, // last comma should be there
	}

	fmt.Println(arr) // [[1 2] [3 4]]

	// Iterating over each element
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}

	// output:
	// arr[0][0] = 1
	// arr[0][1] = 2
	// arr[1][0] = 3
	// arr[1][1] = 4
}
