package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func processIt (fn func(a int) int) int {
// 	return fn(5)
// }
func processIt () func(a int) int {
	return func (a int) int {
		return a
	}
}

func main() {
	// fmt.Println(add(1, 2))
	fn := processIt()
	fmt.Println(fn(2))
}