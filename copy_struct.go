package main

import "fmt"

type doctor struct {
	name string
}

func main() {
	aDoctor := doctor{name: "ddd"}
	aAnotherDoctor := aDoctor
	aDoctor.name = "kiki"
	fmt.Printf("aDoctor=%v\n", aDoctor)
	fmt.Printf("aAnotherDoctor=%v\n", aAnotherDoctor)

	someDoctor := &aDoctor
	aDoctor.name = "kiki2"
	fmt.Printf("aDoctor=%v\n", aDoctor)
	fmt.Printf("someDoctor=%v\n", someDoctor)

}
