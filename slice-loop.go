package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	for k, v := range s {
		fmt.Println("k=", k, " v=", v)
	}

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	// use key only
	for k := range statePopulations {
		fmt.Println(k)
	}
	// use value only
	for _, v := range statePopulations {
		fmt.Println(v)
	}

	str1 := "this is fun"
	for k, v := range str1 {
		fmt.Println("k=", k, " v=", string(v))
	}
}
