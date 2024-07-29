package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("err", err)
	}

	defer res.Body.Close()

	data, errors := io.ReadAll(res.Body)

	if errors != nil {
		fmt.Println()
	}

	println(string(data))


}