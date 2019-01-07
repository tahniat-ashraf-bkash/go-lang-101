package main

import (
	"fmt"
)

func main() {

	x, y := "Hello", 65

	if y == 70 && x == "Hello" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	switch y {
	case 60:
		fmt.Println("y", 60)

	case 65:

		fmt.Println("y", 65)

	default:
		fmt.Println("Falling into default ...")

	}

}
