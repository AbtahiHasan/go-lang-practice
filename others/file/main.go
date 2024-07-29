package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create(".env")
	// if err != nil {
	// 	fmt.Println("error while create file", err)
	// 	return
	// }

	// defer file.Close()

	// bytes, errors := io.WriteString(file,"TOKEN=34432134a"+"\n")

	// if errors != nil {
	// 	fmt.Println("Error while writing file", errors)
	// 	return
	// }

	// fmt.Println("bytes", bytes)

	file, err := os.Open(".env")

	if err != nil {
			fmt.Println("error while create file", err)
			return
		}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error while reading file", err)
			return
		}

		fmt.Println(string(buffer[:n]))
		fmt.Println("n", n)
		fmt.Println("buffer", buffer)
		
	}

}