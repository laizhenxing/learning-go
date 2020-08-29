package downloader

import (
	"bufio"
	"net/http"
	"testing"
	"webcrawler/module/stub"

	"webcrawler/module"
)

func TestMyDownloader_Download(t *testing.T) {
	mid := module.MID("D1|127.0.0.1:8080")
	cli := &http.Client{}
	dl, err := NewDownloader(mid, cli, nil)
	if err != nil {
		t.Fatalf("errors happend: %s (mid: %s, cli: %#v)", err, mid, cli)
	}
	if dl == nil {
		t.Fatalf("can not create a downlader")
	}
	if dl.ID() != mid {
		t.Fatalf("Inconsistent MID for downloader: expected: %s, actual: %s", mid, dl.ID())
	}
	mid = module.MID("D1127.0.0.1")
	dl, err = NewDownloader(mid, cli, nil)
	if err == nil {
		t.Fatalf("No error when create a downloader with illegal MID: %q!\nerror: %s", mid, err)
	}
	mid = module.MID("D1|127.0.0.1:8080")
	cli = nil
	dl, err = NewDownloader(mid, cli, nil)
	if err == nil {
		t.Fatalf("No error when create a downloader with nil http client")
	}
}

func TestDownload(t *testing.T)  {
	mid := module.MID("D1|127.0.0.1:8080")
	httpCli := &http.Client{}
	dl, _ := NewDownloader(mid, httpCli, nil)
	url := "http://www.baidu.com/robots.txt"
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("An error occurs when creating a HTTP request: %s (url: %s)",
			err, url)
	}
	depth := uint32(0)
	req := module.NewRequest(httpReq, depth)
	resp, err := dl.Download(req)
	if err != nil {
		t.Fatalf("An error occurs when downloading content: %s (req: %#v)",
			err, req)
	}
	if resp == nil {
		t.Fatalf("Couldn't create download for request %#v!", req)
	}
	if resp.Depth() != depth {
		t.Fatalf("Inconsistent depth: expected: %d, actual: %d", depth, resp.Depth())
	}

	httpResp := resp.HttpResp()
	if httpResp == nil {
		t.Fatalf("Invalid HTTP Response! {URL: %s}", url)
	}
	body := httpResp.Body
	if body == nil {
		t.Fatalf("Invalid HTTP Response body! {URL: %s}", url)
	}
	r := bufio.NewReader(body)
	line, _, err := r.ReadLine()
	if err != nil {
		t.Fatalf("reading body error: %s! {URL: %s}", err, url)
	}
	lineStr := string(line)
	expectedLine := "User-agent: Baiduspider"
	if lineStr != expectedLine {
		t.Fatalf("Read Line error [expected: %s, actual: %s]", expectedLine, lineStr)
	}

	// 测试有误情况
	_, err = dl.Download(nil)
	if err == nil {
		t.Fatalf("no error when donwload")
	}
	url = "http:///www.baidu.com/robots.txt"
	httpReq, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("create http request from baidu error: %s", err)
	}
	req = module.NewRequest(httpReq, 0)
	resp, err = dl.Download(req)
	if err == nil {
		t.Fatalf("no error with invalid url: %s", url)
	}
	req = module.NewRequest(nil, 0)
	resp, err = dl.Download(req)
	if err == nil {
		t.Fatalf("no error with nil http request")
	}
}

func TestCount(t *testing.T)  {
	mid := module.MID("D1|127.0.0.1:8080")
	httpCli := &http.Client{}
	// 测试初始化后的计数
	dl, _ := NewDownloader(mid, httpCli, nil)
	di := dl.(stub.ModuleInternal)
	if di.CalledCount() != 0 {
		t.Fatalf("called count: expected: %d, actual: %d", 0, di.CalledCount())
	}
	if di.AcceptedCount() != 0 {
		t.Fatalf("accepted count: expected: %d, actual: %d", 0, di.AcceptedCount())
	}
	if di.CompletedCount() != 0 {
		t.Fatalf("completed count: expected: %d, actual: %d", 0, di.CompletedCount())
	}
	if di.HandlingNumber() != 0 {
		t.Fatalf("handling number: expected: %d, actual: %d", 0, di.HandlingNumber())
	}

	// 测试处理失败时的计数
	dl, _ = NewDownloader(mid, httpCli, nil)
	di = dl.(stub.ModuleInternal)
	url := "http:///www.baidu.com/robots.txt"
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("new http request error: %s {URL: %s}", err, url)
	}
	req := module.NewRequest(httpReq, 0)
	_, err = dl.Download(req)
	//if err != nil {
	//	t.Fatalf("download error: %s {URL: %s}", err, url)
	//}
	if di.CalledCount() != 1 {
		t.Fatalf("called count unexpected! expected: %d, actual: %d", 1, di.CalledCount())
	}
	if di.AcceptedCount() != 1 {
		t.Fatalf("Inconsistent accepted count for internal module: expected: %d, actual: %d",
			1, di.AcceptedCount())
	}
	if di.CompletedCount() != 0 {
		t.Fatalf("completed count unexpected! expected: %d, actual: %d", 0, di.CompletedCount())
	}
	if di.HandlingNumber() != 0 {
		t.Fatalf("handling number unexpected! expected: %d, actual: %d", 0, di.HandlingNumber())
	}
}





























