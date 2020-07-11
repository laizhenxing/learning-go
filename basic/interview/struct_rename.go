package main

import "fmt"

type User struct {}
type User1 User		// 创建了新的类型，不同类型
type User2 = User	// 别名

func (i User1) m1()  {
	fmt.Println("m1")
}

func (i User) m2()  {
	fmt.Println("m2")
}

func main() {
	var i1 User1
	var i2 User2
	i1.m2()	// i1.m2 undefined (type User1 has no field or method m2)
	i2.m2()
}
