package main

import "fmt"

func main() {
	str := "this is fun"
	fmt.Println(str)
	changeme(&str)
	fmt.Println(str)

	str1 := "this"
	str2 := "is"
	fmt.Println("before swap=", str1, str2)
	swapme(&str1, &str2)
	fmt.Println("after swap=", str1, str2)

}

func changeme(abcd *string) {
	*abcd = "THIS IS FUN"
}

func swapme(x1 *string, x2 *string) {
	*x1, *x2 = *x2, *x1
}
