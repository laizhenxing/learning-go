package main

import (
	"fmt"
	"sync"
)

func main() {
	sy := sync.Mutex{}
	sy.Lock()
	fmt.Println("hello world~")
	sy.Unlock()
}
