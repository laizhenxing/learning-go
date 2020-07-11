package main

import (
	"fmt"
	"sync"
)

var  sharedRsc = make(map[string]interface{})

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	m := sync.Mutex{}
	c:=sync.NewCond(&m)

	go func() {
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc2"])
		c.L.Unlock()
		wg.Done()
	}()

	// 写入数据
	c.L.Lock()
	sharedRsc["rsc1"] = "rsc1111"
	sharedRsc["rsc2"] = "rsc2222"
	// 写完广播
	c.Broadcast()
	c.L.Unlock()
	// 等待携程的执行
	wg.Wait()
}
