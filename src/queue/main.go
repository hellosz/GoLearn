package main

import (
	"fmt"
	queue "queue/entry"
)

func main() {
	queue := queue.Queue{1}
	queue.Push("hello")
	queue.Push(3)
	fmt.Print(queue)
	fmt.Print(queue.Pop())
	fmt.Println(queue.IsEmpty())

	fmt.Print(queue)
	fmt.Print(queue.Pop())
	fmt.Println(queue.IsEmpty())

	fmt.Print(queue)
	fmt.Print(queue.Pop())
	fmt.Println(queue.IsEmpty())
}
