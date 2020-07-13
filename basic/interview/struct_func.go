package main

import "fmt"

type X struct {

}

func (x *X) test()  {
	fmt.Println("x.test")
}

func main() {
	var x *X
	x.test()

	// X{} 是不可寻址的，不能直接调用方法。
	//X{}.test()	// cannot call pointer method on X literal / cannot take the address of X literal
	t := X{}
	t.test()
}
