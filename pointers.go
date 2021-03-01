package main

import (
	"fmt"
)

func main() {
	a := 10
	b := a
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	c := &a
	fmt.Println("c ", c)
	a = 23
	fmt.Println("c after a changes", c)
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)

}

var myStruct struct {
	name string
}
