package main

import "fmt"

const (
	isUser = 1 << iota
	isAnalyst
	isAnalystManager
	isAdmin
)

func main() {
	var roles byte = isAnalyst | isAdmin
	fmt.Printf("%v\n", roles&isAnalyst == isAnalyst)
	fmt.Printf("%v\n", roles&isAnalystManager == isAnalystManager)
}
