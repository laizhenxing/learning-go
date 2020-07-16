package main

import "fmt"

func main() {


	for i := 0; i < 10; i++ {
	loop:
		fmt.Println(5)
	}
	// goto 不能跳转到其他函数或者内层代码
	goto loop

}
