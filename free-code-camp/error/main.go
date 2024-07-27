package main

import (
	"errors"
	"fmt"
)

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide DIVIDEND %v by zero", de.dividend)
}

func main() {
	

	var err error = errors.New("this is error")
	
		fmt.Println(err)
	// num := 2.1544
	// str := "str"
	// fmt.Printf("%.2f",num) // return 2.15
	// fmt.Printf("%v", num) // return 2.1544
	// fmt.Printf("%s", str) // return str
}