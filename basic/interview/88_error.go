package main

import (
	"fmt"
	"os"
)

func Foo2() error {
	var err *os.PathError = nil

	return err
}

func main() {
	//var e error = nil
	err := Foo2()	// 值为nil,动态类型为 *os.PathError
	fmt.Println(err)
	//fmt.Println(e == nil)
	// 只有在值和动态类型都为nil的情况下，接口值才为nil
	fmt.Println(err == nil)
	//fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", err)
}
