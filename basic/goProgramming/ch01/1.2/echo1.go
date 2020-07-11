// Echo 1 prints its command-line arguments
// echo 1 打印其命令行参数
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
	var s, sep string	// 没有显示初始化，赋予零值[""]
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]	// 字符串拼接
		sep = " "
	}

	fmt.Println(s)
	end := time.Now().UnixNano()
	fmt.Println("运行时间(s): ", end - start)
}
