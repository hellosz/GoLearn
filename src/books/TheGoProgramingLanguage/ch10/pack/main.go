package main

// 引入方式1
// import "fmt"
// import "os"

// 引入方式2-分组
import (
	another_rand "crypto/rand" // 重命名导入，解决冲突
	"fmt"
	"math/rand"
	"os"

	// "strings"
	_ "time" // 使用“_”进行重命名导入，对包级别的变量执行初始化求值，并执行它的init函数
)

func main() {
	fmt.Println("hello import")
	args := os.Args[1:]
	fmt.Println(args)

	// 导入同名rand
	fmt.Println(rand.Int())
	fmt.Println(another_rand.Read([]byte{1, 2, 3}))
}
