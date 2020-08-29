// 基本数据类型定义
package module

import "net/http"

// 数据的接口类型
type Data interface {
	// 用于判断数据是否有效
	Valid() bool
}

/*************************HTTP请求****************************/
// HTTP请求数据类型
type Request struct {
	httpReq *http.Request
	depth   uint32
}

// 判断请求是否有效
func (r *Request) Valid() bool {
	return r.httpReq != nil && r.httpReq.URL != nil
}

// 获取请求实例
func NewRequest(req *http.Request, depth uint32) *Request {
	return &Request{
		req,
		depth,
	}
}

// 获取HTTP请求
func (r *Request) HttpReq() *http.Request {
	return r.httpReq
}

// 获取请求深度
func (r *Request) Depth() uint32 {
	return r.depth
}

/*************************HTTP响应****************************/
// HTTP响应数据类型
type Response struct {
	// HTTP响应
	httpResp *http.Response
	// 响应深度
	depth uint32
}

// 判断响应是否有效
func (r *Response) Valid() bool {
	return r.httpResp != nil && r.httpResp.Body != nil
}

// 创建一个响应实例
func NewResponse(resp *http.Response, depth uint32) *Response {
	return &Response{
		resp,
		depth,
	}
}

// 获取HTTP响应
func (r *Response) HttpResp() *http.Response {
	return r.httpResp
}

// 获取响应深度
func (r *Response) Depth() uint32 {
	return r.depth
}

/*************************条目****************************/
// 条目数据类型
type Item map[string]interface{}

// 判断条目是否有效
func (i Item) Valid() bool {
	return i != nil
}
