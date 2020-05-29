package main

import "fmt"

func main() {
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("too low")
	}
	if guess > number {
		fmt.Println("too high")
	}
	if guess == number {
		fmt.Println("you got it!!!")
	}
	fmt.Println(guess <= number, guess >= number, guess == number)

	if guess < 1 || guess > 100 {
		fmt.Println("guss is not within 1 and 100")
	}

	fmt.Println(returnTrue())
}

func returnTrue() bool {
	fmt.Println("thisis returnTrue function")
	return true
}
