package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("%v \n", KB)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
