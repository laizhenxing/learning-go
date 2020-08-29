package module

// 计算器组件评分的函数类型
type CalculateScore func(counts Counts) uint64

// 建议的组件评分计算函数
func CalculateScoreSimple(counts Counts) uint64 {
	return counts.CalledCount +
		counts.AcceptedCount << 1 +
		counts.CompletedCount << 2 +
		counts.HandlingNumber << 4
}

// 设置给定组件评分
// 结果代表是否成功更新
func SetScore(m Module) bool {
	calculator := m.ScoreCalculator()
	if calculator == nil {
		calculator = CalculateScoreSimple
	}
	newScore := calculator(m.Counts())
	if newScore == m.Score() {
		 return false
	}
	m.SetScore(newScore)
	return true
}