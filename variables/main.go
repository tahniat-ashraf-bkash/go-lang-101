package main

import (
	"fmt"
)

var i int
var s = "Hello World"
var b bool

func main() {

	b1 := true
	f := 0.21
	fmt.Println("i", i)
	fmt.Println("s", s)
	fmt.Println("f", f)
	fmt.Println("b", b)
	fmt.Println("b1", b1)

	s, t := "Hello s", 32

	fmt.Println("s", "t", s, t)

	fmt.Printf("The value of s is %v", s)

}
