package main

import "fmt"

const (
	isUser = 1 << iota
	isAnalyst
	isAnalystManager
	isAdmin
)

func main() {
	fmt.Println(1 << 0)
	fmt.Println("isUser=", isUser)
	fmt.Println("isAnalyst=", isAnalyst)
	fmt.Println("isAnalystManager=", isAnalystManager)
	fmt.Println("isAdmin=", isAdmin)
	fmt.Println(">>>>>>>>>>>>>>>")
	var roles byte = isAnalyst | isAdmin
	fmt.Printf("%v\n", roles&isAnalyst == isAnalyst)
	fmt.Printf("%v\n", roles&isAnalystManager == isAnalystManager)
}
