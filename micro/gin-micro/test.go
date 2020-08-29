package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	ghttp "github.com/micro/go-plugins/client/http/v2"

	"gin-micro/model/protos/prods"
)

func main() {
	reg := etcd.NewRegistry(
		registry.Addrs("localhost:2380"),
	)

	sel := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.Random),
	)

	callApi2(sel)
}

func originCall(reg registry.Registry) {
	srv, err := reg.GetService("go.micro.gin-micro")
	if err != nil {
		log.Fatal(err)
	}

	next := selector.RoundRobin(srv)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}

	addr := node.Address
	rsp, err := callAp11(addr, "/v1/prods", "POST")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

// 原始调用
func callAp11(addr, path, method string) (string, error) {
	//var rsp http.Response
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	cli := http.DefaultClient
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	return string(buf), err
}

// 使用 go-plugins包
func callApi2(sel selector.Selector) {
	cli := ghttp.NewClient(
		client.Selector(sel),
		client.ContentType("application/json"),
	)
	req := cli.NewRequest(
		"go.micro.gin-micro",
		"/v1/prods",
		model.ProdRequest{Size: 5},
	)

	var rsp model.ProdResponse
	err := cli.Call(context.Background(), req, &rsp)
	if err != nil {
		fmt.Println(err)
	}

	if rsp.GetData() == nil {
		fmt.Println("nil")
	}
	fmt.Println(rsp.GetData())
}

func roundRobin(reg registry.Registry) {
	for {
		srv, err := reg.GetService("go.micro.gin-micro")
		if err != nil {
			log.Fatal(err)
		}

		// 随机
		//next := selector.Random(srv)

		// 轮询
		next := selector.RoundRobin(srv)

		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(node.Address)
		time.Sleep(1 * time.Second)
	}
}
