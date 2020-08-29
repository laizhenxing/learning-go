package module

// Type 代表组件的类型
type Type string

// 当前认可的组件类型的常量
const (
	// 下载器
	TYPE_DOWNLOADER Type = "downloader"
	// 分析器
	TYPE_ANALYZER Type = "analyzer"
	// 条目处理管道
	TYPE_PIPELINE Type = "pipeline"
)

// 合法的组件类型-字母映射
var legalTypeLetterMap = map[Type]string{
	TYPE_ANALYZER:   "A",
	TYPE_DOWNLOADER: "D",
	TYPE_PIPELINE:   "P",
}

// 合法的字母-组件类型映射
var legalLetterTypeMap = map[string]Type{
	"A": TYPE_ANALYZER,
	"D": TYPE_DOWNLOADER,
	"P": TYPE_PIPELINE,
}

func LegalType(moduleType Type) bool {
	if _, ok := legalTypeLetterMap[moduleType]; ok {
		return true
	}
	return false
}
