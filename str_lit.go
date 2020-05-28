package main

import "fmt"

func main() {
	s := "this is a string"
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", s[2], s[2])
	fmt.Printf("%v %T\n", string(s[2]), string(s[2]))

	fmt.Println(".......................")
	s2 := "this is also a string"
	s3 := s + " " + s2
	fmt.Printf("%v %T\n", s3, s3)
}
