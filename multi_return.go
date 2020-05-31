package main

import "fmt"

func main() {
	x, e := divide(2, 3)
	fmt.Println(x, e)
}
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}
