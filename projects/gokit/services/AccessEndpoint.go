package services

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type AccessRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Method   string `json:"method"`
}

type AccessResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

func AccessEndpoint(accessService IAccessService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(AccessRequest)
		result := AccessResponse{Status: "OK"}
		if r.Method == "POST" {
			token, err := accessService.GetToken(r.Username, r.Password)
			if err != nil {
				result.Status = "error: " + err.Error()
			} else {
				result.Token = token
			}
			return result, err
		} else {
			result.Status = "请求方法不正确"
		}
		return result, errors.New("请求方法不正确")
	}
}