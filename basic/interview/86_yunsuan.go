package main

import "fmt"

func main() {
	n := 43210
	fmt.Println(n%60)
	fmt.Println(n/(60*60), " hours and ", n%(60*60), " seconds")
}
