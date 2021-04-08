package main

import (
	"fmt"
	"time"
)

func main() {
	t := "30/Mar/2021:09:55:25 +0000"
	res, _ := time.Parse("02/Jan/2006:15:04:05 -0700", t)

	fmt.Println(res)
	fmt.Println(res.Format("2006-01-02 03:04:05"))
}
