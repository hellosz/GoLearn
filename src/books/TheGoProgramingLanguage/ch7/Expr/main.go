package Expr

import "fmt"

// map为引用类型
func TestMap() {
	fmt.Println("hello world!")

	vars := map[string]bool{}
	fmt.Printf("vars = %+v\n", vars)

	testMap("test", vars)
	fmt.Printf("vars = %+v\n", vars)
}

func testMap(text string, vars map[string]bool) error {
	vars[text] = true
	return nil
}
