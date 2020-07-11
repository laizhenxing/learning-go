package main

import (
	"context"
	"fmt"
	"time"

	v3 "go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		config v3.Config
		cli *v3.Client
		err error
	)

	// 客户端配置
	config = v3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立etcd客户端
	cli, err = v3.New(config)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	fmt.Println("connect to etcd server success!!")

	// put 实例
	kv := v3.NewKV(cli)
	// 如果不指定v3.WithPrevKV()，res.PrevKV是不会有值
	res, err := kv.Put(context.TODO(), "cron/jobs/job2", "hello2", v3.WithPrevKV())
	if err != nil {
		fmt.Println("etcd put error: ", err)
	}
	fmt.Println("Revision: ", res.Header.Revision)
	// 覆盖过一个旧值
	if res.PrevKv != nil {
		fmt.Println("PrevValue: ", string(res.PrevKv.Value))
	}
}
