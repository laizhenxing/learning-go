package errors

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// 下载器错误
	ERROR_TYPE_DOWNLOAD ErrorType = "downloader error"
	// 分析器错误
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	// 条目处理器错误
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	// 调度器错误
	ERROR_TYPE_SCHEDULER ErrorType = "shceduler error"
)

// 错误类型
type ErrorType string

// 爬虫错误接口
type CrawlerError interface {
	// 用于获取错误类型
	Type() ErrorType
	// 用于获取错误提示信息
	Error() string
}

// crawlerError 爬虫错误的实现类型
type crawlerError struct {
	// 错误类型
	errType ErrorType
	// 错误信息
	errMsg string
	// 完整错误提示信息
	fullErrMsg string
}

// 创建一个新的爬虫错误值
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &crawlerError{
		errType:    errType,
		errMsg:     strings.TrimSpace(errMsg),
	}
}

// 根据给定的错误值创建一个新的爬虫错误值
func NewCrawlerErrorByError(errorType ErrorType, err error) CrawlerError {
	return NewCrawlerError(errorType, err.Error())
}

func (ce *crawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *crawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.getFullMsg()
	}
	return ce.fullErrMsg
}

// 用于生成错误提示信息，并非相应的字段赋值
func (ce *crawlerError) getFullMsg()  {
	var buff bytes.Buffer
	buff.WriteString("crawler error: ")
	if ce.errType != "" {
		buff.WriteString("[")
		buff.WriteString(string(ce.errType))
		buff.WriteString("] ")
	}
	buff.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buff.String())
}

// 非法参数的错误类型
type IllegalParameterError struct {
	msg string
}

// 创建一个IllegalParameterError类型的实例
func NewIllegalParameterError(errMsg string) IllegalParameterError {
	return IllegalParameterError{
		msg: fmt.Sprintf("illegal parameter[%s]", strings.TrimSpace(errMsg)),
	}
}

func (ipe IllegalParameterError) Error() string {
	return ipe.msg
}

