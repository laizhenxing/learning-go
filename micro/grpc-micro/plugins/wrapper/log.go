package wrapper

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/client"
)


type logWrapper struct {
	client.Client
}


func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("LogWrapper开始调用...")

	return l.Client.Call(ctx, req, rsp, opts...)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

