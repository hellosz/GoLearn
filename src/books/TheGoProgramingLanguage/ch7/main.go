package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	// "hellosz.top/src/books/TheGoProgramingLanguage/ch7/Expr"
)

func main() {
	// interfaceValue()
	// arrayValue()
	// compareSlice()
	// compareErrors()
	// Expr.TestMap()
	value := errors.New("error")
	typeSwitch(value)
}

// typeSwitch 进行类型断言
func typeSwitch(value interface{}) {
	switch x := value.(type) {
	case nil:
		fmt.Println("nil vlaue")
	case int, uint:
		fmt.Println("int value")
	case rune:
		fmt.Println("rune value")
	case byte:
		fmt.Println("byte value")
	case string:
		fmt.Println("string value")
	case bool:
		fmt.Println("bool value")
		// fallthrough // 不允许使用fallthrough
	default:
		fmt.Printf("other type, x = %T, x = %+v\n", x, x)
	}

	switch value {
	case 1:
		fmt.Println("1 value")
		fallthrough
	default:
		fmt.Println("hello world")
	}

}

// interfaceValue 接口值的动态类型和类型值
func interfaceValue() {
	var w io.Writer
	fmt.Printf("w = %+v, w = %T, pointer = %p\n", w, w, w)

	w = os.Stdout
	fmt.Printf("w = %+v, w = %T, pointer = %p\n", w, w, w)

	w = new(bytes.Buffer)
	fmt.Printf("w = %+v, w = %T, pointer = %p\n", w, w, w)

	w = nil
	fmt.Printf("w = %+v, w = %T, pointer = %p\n", w, w, w)
}

// arrayValue 查看切片和数组的实际值
func arrayValue() {
	var arr = []int{1, 2, 3}
	fmt.Printf("arr = %+v, type is %T, literal type is %s\n", arr, arr, reflect.TypeOf(arr).Kind())

	var arr2 = [3]int{1, 2, 3}
	fmt.Printf("arr2 = %+v, type is %T, literal type is %s\n", arr2, arr2, reflect.TypeOf(arr2).Kind())
}

// compareSlice 切片比较大小
func compareSlice() {
	// 报错，因为slice是不可以进行比较的
	var x interface{} = []int{1, 2, 3}
	if x == x {
		fmt.Println("x equals x")
	}
}

func compareErrors() {
	err1, err2 := errors.New("error"), errors.New("error")
	if err1 == err2 {
		fmt.Println("equals")
	} else {
		fmt.Println("not equals")
	}

	// 类型和值都一样，但是因为是指针类型，所以不相等
	fmt.Printf("err1, err2 = %p, %p\n", err1, err2)

	err3, err4 := New("error"), New("error2")
	if err3 == err4 {
		fmt.Println("equals")
	} else {
		fmt.Println("not equals")
	}

	// 类型和值都一样，所以相等，但是如果改变err4成error2，此时不想等
	fmt.Printf("err3, err4 = %p, %p\n", err3, err4)
}

func New(text string) error {
	return myErrorString{text}
}

type myErrorString struct {
	text string
}

func (e myErrorString) Error() string {
	return e.text
}
