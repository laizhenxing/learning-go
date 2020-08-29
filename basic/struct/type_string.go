package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) String() string {
	// fmt.Sprintf 将结构体转字符串时，会调用结构体的 String() 方法，这样造成无限递归，导致堆栈溢出
	return fmt.Sprintf("printf: %+v", p)
}

func main()  {
	p := &Person{}
	p.String()
}
