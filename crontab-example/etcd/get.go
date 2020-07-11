package main

import (
	"context"
	"fmt"
	"time"

	v3 "go.etcd.io/etcd/clientv3"
)

func main()  {
	cfg := v3.Config{Endpoints: []string{"127.0.0.1:2379"},DialTimeout: 5 * time.Second}
	cli, err := v3.New(cfg)
	if err != nil {
		fmt.Println("etcd connect error")
		panic(err)
	}
	defer cli.Close()

	kv := v3.NewKV(cli)
	rsp, err := kv.Get(context.TODO(), "cron/jobs", v3.WithFromKey())
	if err != nil {
		fmt.Println("[etcd] get key error")
		panic(err)
	}
	if rsp.Kvs != nil {
		fmt.Println("count: ", rsp.Count)
		fmt.Println("values: ", rsp.Kvs)
	}
}