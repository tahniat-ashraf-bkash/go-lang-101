package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Hello World")

	s := rand.Int()

	fmt.Println("s", s)

	random := rand.Intn(4)

	fmt.Println("random", random)

	rand.Seed(time.Now().UnixNano())

	fmt.Println("randomNumber", rand.Intn(3)+1)

}
