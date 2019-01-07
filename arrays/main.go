package main

import "fmt"

func main() {

	var arrayDetailed [3]int

	arrayDetailed[0] = 69

	fmt.Println("arrayDetailed",arrayDetailed)
	

	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("my_array", myArray)

	fmt.Println("myArray[3]", myArray[3])

}
