package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	ohiopop, ok := statePopulations["Ohiuo"]
	fmt.Println(ohiopop, ok)
	fmt.Println(statePopulations["Ohio"])
	statePopulations["Georgia"] = 12345
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")

	fmt.Println(">>>>>>>>>>>>")
	sp := statePopulations
	fmt.Println(sp)
	delete(sp, "Florida")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	fmt.Println(">>>>>>>>>>>>")

	m := map[[3]int]string{}
	fmt.Println(m)

	m2 := make(map[string]int, 1)
	m2 = map[string]int{
		"this": 12345,
		"is":   678910,
	}
	fmt.Println(m2)
	fmt.Println(len(m2))

}
