package main

import "fmt"

func main() {
	// can not use "divide" before var declaration
	var divide func(float64, float64) (float64, error)
	divide = func(num float64, den float64) (float64, error) {
		if den == 0.0 {
			return 0.0, fmt.Errorf("divide by zero")
		}
		return num / den, nil
	}

	d, err := divide(4, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
