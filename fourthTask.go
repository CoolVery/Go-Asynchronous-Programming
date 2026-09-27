package main

import "sync"

func Merge(chanels []<-chan int) <-chan int {

	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range chanels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}