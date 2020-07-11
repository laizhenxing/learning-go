package main

import "fmt"

func main() {
	var nilChan chan int
	//nilChan := make(chan int)
	nilChan <- 1
	fmt.Println("nilChan=", nilChan)
}
