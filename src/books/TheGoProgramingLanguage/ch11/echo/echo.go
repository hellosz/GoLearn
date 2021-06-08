//
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "seperator")
)

var out io.Writer = os.Stdout

func main() {
	// 解析输入参数
	flag.Parse()

	// 输出结果
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "Error is: %v\n", err)
		os.Exit(1)
	}
}

// echo 输出结果到指定的输出对象
func echo(newline bool, seperator string, args []string) error {
	fmt.Fprint(out, strings.Join(args, seperator))
	if newline {
		fmt.Fprintln(out)
	}

	return nil
}
