package module

// Counts 代表用于汇集组件内部计数的类型
type Counts struct {
	// 调用计数
	CalledCount uint64
	// 接受计数
	AcceptedCount uint64
	// 成功完成计数
	CompletedCount uint64
	// 实时处理数
	HandlingNumber uint64
}

// 代表组件摘要结构的类型
type SummaryStruct struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepted  uint64      `json:"accepted"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra"`
}

// Module代表组件的基础接口类型
// 该接口的实现类型必须是并发安全的
type Module interface {
	// 用于获取当前的组件的ID
	ID() MID
	// 用于获取当前组件的网络地址的字符串形式
	Addr() string
	// 用于获取当前组件的评分
	Score() uint64
	// 用于设置当前组件的评分
	SetScore(score uint64)
	// 用于获取评分计算器
	ScoreCalculator() CalculateScore
	// 用于获取当前组件被调用的计数
	CalledCount() uint64
	// 用于获取组件接受的调用计数
	// 组件一般会由于超负荷或者参数有误而拒绝调用
	AcceptedCount() uint64
	// 用于获取当前组件已经成功完成调用的计数
	CompletedCount() uint64
	// 用于互殴去当前组件正砸处理的调用的数量
	HandlingNumber() uint64
	// 用于一次性获取所有计数
	Counts() Counts
	// 用于获取组件摘要
	Summary() SummaryStruct
}

// 下载器接口类型
type Downloader interface {
	Module
	// 根据请求获取内容并响应
	Download(req *Request) (*Response, error)
}
