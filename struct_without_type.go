package main

import "fmt"

func main() {
	aDoctor := struct {
		name string
	}{name: "ddd"}

	aAnotherDoctor := aDoctor
	//aAnotherDoctor.name = "kkk"
	fmt.Println(aDoctor)
	fmt.Println(aAnotherDoctor)
}
