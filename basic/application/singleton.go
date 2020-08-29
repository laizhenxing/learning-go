// 使用 sync.Once 实现单例模式
// 使用原子操作配合互斥锁可以实现高效的单例模式
package application

import "sync"

type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
