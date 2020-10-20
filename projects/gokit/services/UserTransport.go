package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	if uid, ok := params["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid: uid,
			Method: r.Method,
			Token: r.URL.Query().Get("token"),	// 获取请求url中的token
		}, nil
	}

	return nil, errors.New("参数错误")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
