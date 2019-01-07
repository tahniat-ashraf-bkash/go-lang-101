package main

import "fmt"

func main() {

	//for loop big
	for i := 0; i < 10; i++ {
		fmt.Println("i", i)
	}

	i := 10
	//for loop - working as while

	for i > 0 {
		fmt.Println("i", i)
		i--
	}

}
