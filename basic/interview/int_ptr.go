package main

import "fmt"

var p1 *int
//var err error

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar()  {
	// 使用了未定义的p1指针
	fmt.Println(*p1)	// panic: runtime error: invalid memory address or nil pointer dereference
}

func main() {
	// p1是局部变量，不是全局变量，这里覆盖了全局变量p1
	// 解决：可以先声明 err
	p1, err := foo()
	if err != nil {
		println(err)
		return
	}
	bar()
	fmt.Println(*p1)
}