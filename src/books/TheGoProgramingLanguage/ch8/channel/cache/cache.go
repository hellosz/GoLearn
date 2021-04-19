package main

import (
	"fmt"
	"os"
)

func main() {
	input := make(chan []string)

	go func() {
		input <- os.Args[1:]
	}()

	for in := range input {
		for _, str := range in {
			fmt.Println(str)
		}

		// 此处必须使用goroutine进行解耦，否则会发生死锁
		go func() {
			input <- []string{"link1", "link2"}
		}()
	}
}
