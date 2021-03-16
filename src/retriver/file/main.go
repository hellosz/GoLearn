package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 打印文件内容
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContent(file)
}

// 接受Reader接口，读取读取数据
func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	str := `
	hello
	once again
	`
	printFileContent(strings.NewReader(str))

	filename := "./test.txt"
	printFile(filename)
}
