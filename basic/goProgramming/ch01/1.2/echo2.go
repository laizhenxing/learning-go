// Echo 2 打印命令行参数
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {

	}
	s, sep := "", ""	// 短变量声明并初始化。只能用在函数内部，不能用于包变量
	for _, arg := range os.Args[1:] {	// _ 表示空白标识符（blank identifier）
		s += sep + arg
		sep = " "
	}

	//fmt.Println(os.Args[0])
	fmt.Println(s)
	end := time.Now().UnixNano()
	fmt.Println("运行时间(s): ", end - start)
}
