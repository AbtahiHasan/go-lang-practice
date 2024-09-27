package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views +=  1
}

func  main()  {
	post := Post{
		views: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go post.inc(&wg)
	}

	wg.Wait()

	fmt.Println(post.views)
}