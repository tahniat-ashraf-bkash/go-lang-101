package main

import "fmt"

func main() {

	//declare only
	var mapX map[int]string

	mapX = make(map[int]string)

	mapX[1] = "T"
	mapX[2] = "A"
	mapX[3] = "P"

	fmt.Println("mapx", mapX)

	fmt.Println("mapX[3]", mapX[3])

	mapY := map[int]string{
		1: "S",
		2: "A",
		3: "D",
	}

	fmt.Println("mapY", mapY)
}
