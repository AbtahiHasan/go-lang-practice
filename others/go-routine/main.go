package main

import (
	"fmt"
	"sync"
)

func worker (i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("work %d is finished\n", i)
}


func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("all work finished")
}