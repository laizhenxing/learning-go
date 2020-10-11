package geecache

import (
	"fmt"
	"log"
	"sync"
)

// 核心数据结构，负责与用户的交互，并控制缓存值和获取的流程
// 可以认为是缓存的命名空间
type Group struct {
	name      string // 名称
	getter    Getter // 未命中时获取数据源的回调
	mainCache cache  // 并发缓存
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		// 缓存命中
		log.Println("[GeeCache] hit")
		return v, nil
	}

	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

// 获取数据源数据，并将书加入缓存
func (g *Group) getLocally(key string) (ByteView, error)  {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)

	return value, nil
}

// 将数据加入缓存
func (g *Group) populateCache(key string, value ByteView)  {
	g.mainCache.add(key, value)
}

// 回调Getter
// 缓存不存在，应从数据元获取数据并添加到缓存
type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}