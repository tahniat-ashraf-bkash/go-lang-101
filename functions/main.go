package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(greet("Priyam"))

	fmt.Println(swap(1, 2))
	fmt.Println(shuffle(1, 2, 3))

}

func greet(name string) string {
	return ("Hello Sir " + name)
}

func swap(x int, y int) (int, int) {
	return y, x
}

func shuffle(x int, y int, z int) (int, int, int) {
	return z, x, y
}
