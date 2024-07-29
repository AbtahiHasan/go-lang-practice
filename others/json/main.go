package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Go lang", Age: 9}
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error", err)
	}

	// fmt.Println("data", string(data))
	
	var personData Person
	
	errors := json.Unmarshal(data, &personData)
	
	if errors != nil {
		return
	}
	fmt.Println("data", personData)



}