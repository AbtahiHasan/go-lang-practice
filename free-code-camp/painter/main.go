package main

import "fmt"

func main() {
	myString := "hello"
	myStringPtr := &myString

	fmt.Println(*myStringPtr)
	*myStringPtr = "world"
	fmt.Println(*myStringPtr)
	fmt.Println(myString)
}