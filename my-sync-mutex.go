package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex

var wg = sync.WaitGroup{}

func main() {
	var i int = 0

	wg.Add(3)
	go inc(&i)
	go inc(&i)
	go inc(&i)

	//fmt.Println(i)
	wg.Wait()
}

func inc(i *int) {
	mut.Lock()
	*i = *i + 1
	fmt.Println(*i)
	mut.Unlock()
	wg.Done()
}
