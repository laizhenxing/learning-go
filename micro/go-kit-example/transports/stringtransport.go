package transports

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"

	"go-kit-example/endpoints"
)

func DecodeUppercaseRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//if r.URL.Query().Get("s") != "" {
	//	req := endpoints.UppercaseRequest{S: r.URL.Query().Get("s")}
	//	return req, nil
	//}
	vars := mux.Vars(r)
	if s, ok := vars["s"]; ok {
		req := endpoints.UppercaseRequest{
			S: s,
		}
		return req, nil
	}
	return nil, errors.New("参数错误")
	//var req endpoints.UppercaseRequest
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	return nil, err
	//}
	//return req, nil
}

func DecodeCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if s, ok := vars["s"]; ok {
		req := endpoints.CountRequest{
			S: s,
		}
		return req, nil
	}
	return nil, errors.New("参数错误")
	//var req endpoints.CountRequest
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	return nil, err
	//}
	//return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, rsp interface{}) error {
	return json.NewEncoder(w).Encode(rsp)
}