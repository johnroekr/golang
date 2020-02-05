package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
}

func main() {
	x := 123
	y := 456
	z := "hi"
	p := Person{}
	p_copy := new(Person)

	printIt(x, y, z, p, p_copy)
}

func printIt(v ...interface{}) {
	buff := v
	for _, temp := range buff {
		fmt.Println(reflect.TypeOf(temp))
	}
}
