package main

import "fmt"

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int)  {
	o.Quantity += n
}

func (o *Orange) Decrease(n int)  {
	o.Quantity -= 5
}

func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity)
}

func main() {
	//var o Orange
	var o *Orange = &Orange{}
	o.Increase(10)
	o.Decrease(5)
	// println 输出时不会调到string方法，它是指针方法
	fmt.Println(o)
}
