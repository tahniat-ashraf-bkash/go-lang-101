package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Hello World")

	rand.Seed(time.Now().UnixNano())

	c := rand.Intn(5) + 1

	fmt.Println("c", c)

}
