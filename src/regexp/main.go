package main

import (
	"fmt"
	"regexp"
)

const text = `
allen's email is 237241007@qq.com
allen's email is 123@qq.com
allen's email is 123421@qq.com
`

func main() {
	reg := regexp.MustCompile(`([0-9a-zA-Z])+@[0-9a-zA-Z]+.[0-9a-zA-Z]+`)
	match := reg.FindAllString(text, -1)
	fmt.Println(match)
}
