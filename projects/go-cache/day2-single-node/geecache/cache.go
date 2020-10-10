package geecache

import (
	"sync"

	"day2-single-node/geecache/lru"

)

// 实现缓存的并发控制，添加互斥锁
type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView)  {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 延迟初始化，主要用于提高性能，减少程序内存要求
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), true
	}

	return
}
