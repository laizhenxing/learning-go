package main

//import (
//	"context"
//	"flag"
//	"fmt"
//	"log"
//	"net"
//
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//
//	"github.com/laizhenxing/laracom/demo-service/proto/demo"
//)
//
//const (
//	address  = "localhost:9091"
//	grpcPort = ":9999"
//	httpPort = "8002"
//	appName  = "Demo Service"
//)
//
//type DemoService struct {
//}
//
//func (ds *DemoService) SayHello(ctx context.Context, request *demo.DemoRequest) (*demo.DemoResponse, error) {
//	return &demo.DemoResponse{Text: "你好，" + request.Name}, nil
//}
//
//func main() {
//	// 根据 -mode 参数判断启动哪个模式的代码
//	mode := flag.String("mode", "grpc", "mode:grpc/http/client")
//	flag.Parse()
//	fmt.Println("run mode: ", *mode)
//
//	switch *mode {
//	case "http":
//		httpSrv()
//	case "client":
//		client()
//	case "grpc":
//		fallthrough
//	default:
//		listener, err := net.Listen("tcp", grpcPort)
//		if err != nil {
//			log.Fatalf("监听指定端口失败：%v", err)
//		}
//
//		server := grpc.NewServer()
//		demo.RegisterDemoServiceServer(server, &DemoService{})
//
//		reflection.Register(server)
//
//		if err := server.Serve(listener); err != nil {
//			log.Fatalf("服务启动失败：%v", err)
//		}
//
//	}
//}
