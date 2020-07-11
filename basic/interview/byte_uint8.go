package main

import "fmt"

func main() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	test(a)
	test(b)
	test(c)
}

func test(x byte) {
	fmt.Println(x)
}
