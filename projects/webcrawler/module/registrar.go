package module

import "sync"

type Registrar interface {
	// 注册实例组件
	Register(module Module) (bool, error)
	// 注销组件实例
	Unregister(mid MID) (bool, error)
	// 获取一个指定类型的组件实例
	// 使用负载均衡策略
	Get(mType Type) (Module, error)
	// 获取指定类型的所有组件实例
	GetAllByType(mType Type) (map[MID]Module, error)
	// 获取所有组件实例
	GetAll() map[MID]Module
	// 清除所有注册组件记录
	Clear()
}

func NewRegistrar() Registrar {
	return &registrar{
		moduleTypeMap: map[Type]map[MID]Module{},
	}
}

type registrar struct {
	// 组件类型与对应组件实例的映射
	moduleTypeMap map[Type]map[MID]Module
	// 组件注册专用读写锁
	rwLock sync.RWMutex
}

func (r *registrar) Register(module Module) (bool, error) {
	panic("implement me")
}

func (r *registrar) Unregister(mid MID) (bool, error) {
	panic("implement me")
}

func (r *registrar) Get(mType Type) (Module, error) {
	panic("implement me")
}

func (r *registrar) GetAllByType(mType Type) (map[MID]Module, error) {
	panic("implement me")
}

func (r *registrar) GetAll() map[MID]Module {
	panic("implement me")
}

func (r *registrar) Clear() {
	panic("implement me")
}

