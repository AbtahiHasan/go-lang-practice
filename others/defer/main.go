package main

import "fmt"

func main() {
	fmt.Println("first")
	defer fmt.Println("secound")
	fmt.Println("third")
}