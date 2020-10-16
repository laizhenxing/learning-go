package lru

import (
	"container/list"
)

// LRU Cache, 不是并发安全
type Cache struct {
	maxBytes  int64                         // 允许使用最大内存
	nbytes    int64                         // 当前使用的内存
	ll        *list.List                    // 双向链表
	cache     map[string]*list.Element      // 存储节点的map
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数，可以为nil
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// Cache的构造函数
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 这里规定： front-队尾，back-队首
// 查找功能，1.从字典中找到对应的双向链表的结点，2.将该节点移动到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除功能，缓存淘汰，移除最近最少访问节点（队首）
func (c *Cache) RemoveOldest()  {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil{
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增/更新
func (c *Cache) Add(key string, value Value)  {
	if c.maxBytes == 0 {
		return
	}

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	// 超过最大容量，需要移除队首的节点
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// 当前链表的长度
func (c *Cache) Len() int {
	return c.ll.Len()
}