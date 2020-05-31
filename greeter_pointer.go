package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	// copy of the struct got here
	fmt.Println(g.greeting, g.name)
	g.name = "kiki"
}

func main() {
	g := greeter{
		greeting: "happy bday",
		name:     "ddd",
	}
	g.greet()
	g.greet()

}
