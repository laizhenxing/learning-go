package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Println("3<<1", 3<<1)
	fmt.Println("1<<1", 1<<0)
	fmt.Println("1<<1", 1<<1)
	fmt.Println("1<<2", 1<<2)
	fmt.Println("1<<3", 1<<3)
	fmt.Println("1<<4", 1<<4)
	fmt.Println("1<<5", 1<<5)
	fmt.Printf("%d %08b\n", x, x)
	fmt.Printf("%d %08b\n", y, y)

	fmt.Printf("%08b\t%d\t%d\n", x&y, x&y, 2)
	fmt.Printf("%08b\t%d\t%d\n", x|y, x|y, 38)
	fmt.Printf("%08b\t%d\t%d\n", x^y, x^y, 36)
	fmt.Printf("%08b\t%d\t%d\n", x&^y, x&^y, 32)
}
