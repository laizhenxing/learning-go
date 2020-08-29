package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"go-kit-example/services"
)

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	Val string `json:"val"`
	Err string `json:"err,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	Val int `json:"val"`
}

func MakeUppercaseEndpoint(svc services.StringService) endpoint.Endpoint  {
	return func(ctx context.Context, req interface{}) (rsp interface{}, err error) {
		r := req.(UppercaseRequest)
		v, err := svc.Uppercase(r.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc services.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)

		return CountResponse{v}, nil
	}
}