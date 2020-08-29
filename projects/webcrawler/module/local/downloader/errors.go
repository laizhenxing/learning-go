package downloader

import "webcrawler/errors"

// 用于生成爬虫参数值
func genError(errMsg string) error {
	return errors.NewCrawlerError(errors.ERROR_TYPE_DOWNLOAD, errMsg)
}

// 用于生成爬虫参数错误值
func genParameterError(errMsg string) error {
	return errors.NewCrawlerErrorByError(errors.ERROR_TYPE_DOWNLOAD,
		errors.NewIllegalParameterError(errMsg))
}