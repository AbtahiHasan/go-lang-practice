package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// for i,num := range numbers {
	// 	fmt.Println(i,num)
	// }

	// m := map[string]int{"age": 20, "id": 154}

	// for k, v := range m {
	// 	fmt.Println(k,v)
	// }

	for i, c := range "golang" {
		fmt.Println(i, string(c))

	}

}