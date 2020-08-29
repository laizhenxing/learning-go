package stub

import module2 "webcrawler/module"

// 组件的内部基础接口类型
type ModuleInternal interface {
	module2.Module
	// 把调用计数增加1
	IncrCalledCount()
	// 把接受计数增加1
	IncrAcceptedCount()
	// 把成功完成计数增加1
	IncrCompletedCount()
	// 把实时处理数加1
	IncrHandlingNumber()
	// 把实时处理数减1
	DecrHandlingNumber()
	// 清空所有计数
	Clear()
}