package services

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func DecodeAccessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	// 读取请求体
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// 使用 gjson 进行解析
	params := gjson.Parse(string(body))
	if params.IsObject() {
		username := params.Get("username")
		password := params.Get("password")
		return AccessRequest{
			Username: username.String(),
			Password: password.String(),
			Method:   r.Method,
		}, nil
	}
	return nil, errors.New("参数错误")
}

func EncodeAccessResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}