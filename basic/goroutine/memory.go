package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter = 0

func add(a, b int, c *sync.Mutex)  {
	sum := a + b
	c.Lock()
	counter++
	c.Unlock()
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

func main() {
	start := time.Now()
	s := &sync.Mutex{}
	for i := 0; i < 100; i++ {
		go add(i, i+1, s)
	}

	for {
		s.Lock()
		c := counter
		s.Unlock()
		runtime.Gosched()
		if c >= 100 {
			break
		}
	}

	end := time.Now()
	t := end.Sub(start).Seconds()
	fmt.Println("执行耗时(s): ", t)
}
