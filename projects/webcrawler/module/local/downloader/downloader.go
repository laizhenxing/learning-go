package downloader

import (
	"net/http"

	"webcrawler/helper/log"
	"webcrawler/module"
	"webcrawler/module/stub"
)

// 日志记录器
var logger = log.DLogger()

// 创建一个下载器实例
func NewDownloader(
	mid module.MID,
	client *http.Client,
	scoreCalculator module.CalculateScore) (module.Downloader, error) {
	moduleBase, err := stub.NewModuleInternal(mid, scoreCalculator)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, genParameterError("nil http client")
	}
	return &myDownloader{
		ModuleInternal: moduleBase,
		httpClient:     *client,
	}, nil
}

type myDownloader struct {
	// 组件基础实例
	stub.ModuleInternal
	// 下载用的HTTP客户端
	httpClient http.Client
}

func (m *myDownloader) Download(req *module.Request) (*module.Response, error) {
	m.ModuleInternal.IncrHandlingNumber()
	defer m.ModuleInternal.DecrHandlingNumber()
	m.ModuleInternal.IncrCalledCount()

	if req == nil {
		return nil, genParameterError("nil request")
	}
	httpReq := req.HttpReq()
	if httpReq == nil {
		 return nil, genParameterError("nil HTTP Request")
	}

	m.ModuleInternal.AcceptedCount()
	logger.Infof("Do the request {URL: %s, depth: %d}...", httpReq.URL, req.Depth())

	httpResp, err := m.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	m.ModuleInternal.IncrCompletedCount()

	return module.NewResponse(httpResp, req.Depth()), nil
}
