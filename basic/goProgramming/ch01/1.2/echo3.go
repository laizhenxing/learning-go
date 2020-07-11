// Echo 3 打印命令行参数
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {

	}
	fmt.Println(strings.Join(os.Args[1:], " "))		// 使用strings包Join方法拼接字符串
	end := time.Now().UnixNano()
	fmt.Println("运行时间(s): ", end - start)
}
