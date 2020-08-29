package stub

import (
	"fmt"
	"sync/atomic"

	"webcrawler/errors"
	"webcrawler/helper/log"
	module2 "webcrawler/module"
)

// 日志记录器
var logger = log.DLogger()

// 组件内部基础接口的实现类型
type module struct {
	// 组件ID
	mid module2.MID
	// 网络地址
	addr string
	// 组件评分
	score uint64
	// 评分计数器
	scoreCalculator module2.CalculateScore
	// 调用计数
	calledCount uint64
	// 接受计数
	acceptedCount uint64
	// 成功完成计数
	completedCount uint64
	// 实时处理数
	handlingNumber uint64
}

// 用于创建一个内部组件基础类型实例
func NewModuleInternal(
	mid module2.MID,
	scoreCalculator module2.CalculateScore) (ModuleInternal, error) {
	parts, err := module2.SplitMID(mid)
	if err != nil {
		return nil, errors.NewIllegalParameterError(
			fmt.Sprintf("illegal mid: %q: %s", mid, err))
	}
	return &module{
		mid:             mid,
		addr:            parts[2],
		scoreCalculator: scoreCalculator,
	}, nil
}

func (m *module) ID() module2.MID {
	return m.mid
}

func (m *module) Addr() string {
	return m.addr
}

func (m *module) Score() uint64 {
	return atomic.LoadUint64(&m.score)
}

func (m *module) SetScore(score uint64) {
	atomic.StoreUint64(&m.score, score)
}

func (m *module) ScoreCalculator() module2.CalculateScore {
	return m.scoreCalculator
}

func (m *module) CalledCount() uint64 {
	return atomic.LoadUint64(&m.calledCount)
}

func (m *module) AcceptedCount() uint64 {
	return atomic.LoadUint64(&m.acceptedCount)
}

func (m *module) CompletedCount() uint64 {
	return atomic.LoadUint64(&m.completedCount)
}

func (m *module) HandlingNumber() uint64 {
	return atomic.LoadUint64(&m.handlingNumber)
}

func (m *module) Counts() module2.Counts {
	return module2.Counts{
		CalledCount:    atomic.LoadUint64(&m.calledCount),
		AcceptedCount:  atomic.LoadUint64(&m.acceptedCount),
		CompletedCount: atomic.LoadUint64(&m.completedCount),
		HandlingNumber: atomic.LoadUint64(&m.handlingNumber),
	}
}

func (m *module) Summary() module2.SummaryStruct {
	counts := m.Counts()
	return module2.SummaryStruct{
		ID:        m.ID(),
		Called:    counts.CalledCount,
		Accepted:  counts.AcceptedCount,
		Completed: counts.CompletedCount,
		Handling:  counts.HandlingNumber,
		Extra:     nil,
	}
}

func (m *module) IncrCalledCount() {
	atomic.AddUint64(&m.calledCount, 1)
}

func (m *module) IncrAcceptedCount() {
	atomic.AddUint64(&m.acceptedCount, 1)
}

func (m *module) IncrCompletedCount() {
	atomic.AddUint64(&m.completedCount, 1)
}

func (m *module) IncrHandlingNumber() {
	atomic.AddUint64(&m.handlingNumber, 1)
}

func (m *module) DecrHandlingNumber() {
	atomic.AddUint64(&m.handlingNumber, ^uint64(0))
}

func (m *module) Clear() {
	atomic.StoreUint64(&m.calledCount, 0)
	atomic.StoreUint64(&m.acceptedCount, 0)
	atomic.StoreUint64(&m.completedCount, 0)
	atomic.StoreUint64(&m.handlingNumber, 0)
}
