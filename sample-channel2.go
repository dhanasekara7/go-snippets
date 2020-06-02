package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- j
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(2 * time.Second)
}
