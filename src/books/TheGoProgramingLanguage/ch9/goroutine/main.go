package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Read([]byte{1, 2, 3})
	for {
		go fmt.Println("0")
		fmt.Println("1")
	}
}
