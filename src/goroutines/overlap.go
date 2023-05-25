// package main

// import (
// 	"fmt"
// 	"time"
// )

// func g(v int) {
// 	fmt.Print(v*2, v*v, " ")
// }

// func SpawnGoroutines(n int) {
// 	for i := 0; i < n; i++ {
// 		go g(i)
// 	}
// }

// func main() {
// 	go SpawnGoroutines(10)

// 	time.Sleep(1 * time.Second)
// }
