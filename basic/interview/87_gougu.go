package main

import "fmt"

func main() {
	// 3:011 2: 010 3^2=001=1
	// 4: 100 4^2=110=6
	// 5: 101 5^2=111=7
	// 6:110 2:010 6^2=100=4
	// 8:1000 8^2=1010=10
	// 10:1010 10^2=1000=8
	fmt.Println(3^2)
	fmt.Println(4^2)
	fmt.Println(5^2)
	fmt.Println(6^2)
	fmt.Println(8^2)
	fmt.Println(10^2)
	fmt.Println(3^2+4^2 == 5^2)
	fmt.Println(6^2+8^2 == 14)
}
