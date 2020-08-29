package progressbar

import "fmt"

type Bar struct {
	percent int64  // 百分比
	cur     int64  // 当前进度条位置
	total   int64  // 总进度
	rate    string // 进度条
	graph   string // 显示符号
}

func (b *Bar) NewOption(start, total int64) {
	b.cur = start
	b.total = total
	if b.graph == "" {
		b.graph = "█"
	}
	b.percent = b.getPercent()
	for i := 0; i < int(b.percent); i += 2 {
		b.rate += b.graph // 初始化进度条
	}
}

func (b *Bar) getPercent() int64 {
	return int64(float32(b.cur) / float32(b.total) * 100)
}

func (b *Bar) NewOptionWithGraph(start, total int64, graph string) {
	b.graph = graph
	b.NewOption(start, total)
}

func (b *Bar) Play(cur int64) {
	b.cur = cur
	last := b.percent
	b.percent = b.getPercent()
	if b.percent != last && b.percent%2 == 0 {
		b.rate += b.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", b.rate, b.percent, b.cur, b.total)
}

func (b *Bar) Finish() {
	fmt.Println()
}
