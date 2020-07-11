package main

import "sync"

var cache = struct {
	sync.Mutex	// 互斥锁嵌入到当前的匿名结构体中，所以可以使用 sync.Mutex 的所有方法
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LookUp(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}