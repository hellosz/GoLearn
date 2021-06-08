package main

import (
	"bytes"
	"fmt"
	"testing"
)

// TestEcho 白盒测试
func TestEcho(t *testing.T) {
	// 表格驱动测试
	var seeds = []struct {
		Newline   bool
		Seperator string
		Args      []string
		Want      string
	}{
		{true, ",", []string{"hello", "world"}, "hello,world\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
	}

	for _, seed := range seeds {
		// 初始化输出流，覆盖命令行文件中的参数
		out = new(bytes.Buffer)
		desc := fmt.Sprintf("echo(%t, %s, %s)", seed.Newline, seed.Seperator, seed.Args)

		// 调研echo
		if err := echo(seed.Newline, seed.Seperator, seed.Args); err != nil {
			t.Errorf("execute %s failed, %s\n", desc, err)
		}

		// 验证输出结果
		get := out.(*bytes.Buffer).String()
		if get != seed.Want {
			t.Errorf("echo(%t, %s, %s) get %s, want %s", seed.Newline, seed.Seperator, seed.Args, get, seed.Want)
		}
	}
}
