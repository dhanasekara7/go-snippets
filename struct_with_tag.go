package main

import (
	"fmt"
	"reflect"
)

type Animal2 struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal2{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
