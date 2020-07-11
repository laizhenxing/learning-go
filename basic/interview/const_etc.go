package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

var (
	//a = nil	// use of untyped nil
	b interface{} = nil
	//c string = nil	// cannot use nil as type string in assignment
	d error = nil
)

func main() {
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
	fmt.Println("k: ", k)
	fmt.Println("p: ", p)
}
