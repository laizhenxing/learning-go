package http

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var Ch chan string
var Labels map[string]float64

func NewMultipleHostReverseProxy(target url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		go func() {
			Ch <- req.RequestURI
		}()
	}
	return &httputil.ReverseProxy{
		Director:       director,
	}
}

func Proxy()  {
	InitData()
	target := url.URL{
		Scheme:     "http",
		Host:		":9091",
	}
	proxy := NewMultipleHostReverseProxy(target)
	log.Fatal(http.ListenAndServe(":9090", proxy))
}

func InitData()  {
	Ch = make(chan string)	// api 通道
	Labels = make(map[string]float64)	// api访问次数存储
}