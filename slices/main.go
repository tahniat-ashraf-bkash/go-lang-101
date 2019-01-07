package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	myArraySlice := []string{"Tahniat", "Ashraf", "Priyam", "Sohan"}

	fmt.Println("myArraySlice", myArraySlice)

	print(myArraySlice)

	p := myArraySlice[:0]

	print(p)

	p = myArraySlice[:4]

	print(p)

	p = p[2:4]

	print(p)

	p = p[:2]

	print(p)

	var a []string

	print(a)

	array := [5]string{"A", "B", "C", "D", "E"}

	p = array[0:3]

	print(p)

}

func print(x []string) {
	fmt.Printf("x [] string %v len(x) %d cap(x) %d\n", x, len(x), cap(x))
}
