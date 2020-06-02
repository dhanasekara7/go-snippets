package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "good"
	time.Sleep(2 * time.Second)
}
