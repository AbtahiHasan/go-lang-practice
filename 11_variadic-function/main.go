package main

import "fmt"

func sum(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 54, 6, 7, 85}
	result := sum(numbers...)
	fmt.Println(result)
}