package main

import "fmt"

const i = 100
var j = 123

func main() {
	fmt.Println(&j, j)
	fmt.Println(&i, i)	//  cannot take the address of i
}
