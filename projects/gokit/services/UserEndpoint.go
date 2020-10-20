package services

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-kit/kit/endpoint"
	gkLog "github.com/go-kit/kit/log"
	"golang.org/x/time/rate"

	"gokit/util"
)

type UserRequest struct {
	Uid    int    `json:"uid"`
	Method string `json:"method"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Result string `json:"result"`
}

// 加入限流的中间件
func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.AllowN(time.Now(), 2) {
				return nil, util.NewError(429, "too many requests")
			}
			return next(ctx, request)
		}
	}
}

// 日志中间件
func UserLogMiddleware(logger gkLog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest)
			logger = gkLog.NewLogfmtLogger(os.Stdout)
			logger = gkLog.WithPrefix(logger, "gokit", "1.0")
			logger = gkLog.With(logger, "time", gkLog.DefaultTimestampUTC)
			logger = gkLog.With(logger, "caller", gkLog.DefaultCaller)

			logger.Log("method", r.Method, "event", "get user info", "userId", r.Uid)
			return next(ctx, request)
		}
	}
}

// token 验证中间件
func CheckTokenMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest)
			tokenClaim, err := util.ParseTokenWithClaims(r.Token)
			if err != nil {
				return nil, util.NewError(403, err.Error())
			}
			// 传递用户的信息
			nextCtx := context.WithValue(ctx, "user", tokenClaim.(*util.UserClaim).Name)
			return next(nextCtx, request)
		}
	}
}

func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		fmt.Println("当前登录的用户是：", ctx.Value("user"))
		result := ""
		if r.Method == "GET" {
			result = userService.GetName(r.Uid) + strconv.Itoa(util.ServicePort)
		} else if r.Method == "DELETE" {
			err = userService.DelUser(r.Uid)
			if err != nil {
				result = err.Error()
			} else {
				result = fmt.Sprintf("用户ID为%d,删除成功", r.Uid)
			}
		} else {
			result = "method 不是可行的方法"
		}
		return UserResponse{Result: result}, nil
	}
}

func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain:charset=utf-8", []byte(err.Error())
	w.Header().Set("content-type", contentType)
	if e, ok := err.(*util.MyError); ok {
		w.WriteHeader(e.Code)
		w.Write(body)
	} else {
		w.WriteHeader(500)
		w.Write(body)
	}
}
