package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment1()
	}
	wg.Wait()
}

func increment1() {
	counter++
	wg.Done()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}
