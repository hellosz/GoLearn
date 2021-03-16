package main

import (
	"fmt"
	"retriver/mock"
	"retriver/real"
	"time"
)

func main() {
	var r mock.Retriver
	r = mock.Container{"fabulous"}
	// fmt.Println(retriver.Download(r))
	inspec(r)

	// 指针类型的方法
	r = &real.HttpClient{"Mozilla/5.0", time.Minute}
	inspec(r)
	// fmt.Println(r.Get("http://www.imooc.com"))

	// 获取接口类型的三种方式
	// 1. 使用Printf打印
	// 2. 使用switch类型判断
	// 3. 使用Type Assertion

	fmt.Println("use Type Assertion 判断类型")
	// 转换成mock类型
	// rr := r.(mock.Container)
	// fmt.Printf("%T %s", rr, rr.Content)
	rr := r.(*real.HttpClient)
	fmt.Printf("%T %s", rr, rr.UserAgent)
}

// 使用switch查看接口类型
func inspec(r mock.Retriver) {
	fmt.Println(">> inspect")
	fmt.Println(r)
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case mock.Container:
		fmt.Printf("%v %s\n", v, v.Content)
	case *real.HttpClient:
		fmt.Printf("%v %s\n", v, v.UserAgent)
	}
}

// 使用type Assertion判断接口类型
