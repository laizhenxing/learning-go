package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err:= os.Open("array_for.go")
	if err != nil {
		fmt.Println("os.Open: ", err)
	}
	// defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("ioutil.ReadAll: ", err)
	}
	fmt.Println(string(b))
}
