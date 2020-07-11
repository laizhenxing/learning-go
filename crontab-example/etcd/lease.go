package main

import (
	"context"
	"fmt"
	"time"

	v3 "go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		err error
		cli *v3.Client
	)
	cfg := v3.Config{Endpoints: []string{"127.0.0.1:2379"},DialTimeout: 5 * time.Second}
	cli, err = v3.New(cfg)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// 使用10s租期
	resp, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println("grant err")
		panic(err)
	}
	kv := v3.NewKV(cli)
	_, err = kv.Put(context.TODO(), "/cron/jobs/job1", "test1", v3.WithLease(resp.ID))
	if err != nil {
		fmt.Println("put error")
		panic(err)
	}

	// 续约永久租期
	//ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	//if err != nil {
	//	fmt.Println("keepalive error")
	//	panic(err)
	//}
	//ka := <-ch
	//fmt.Println("ttl: ", ka.TTL)

	// 续约一次租期
	ka, err := cli.KeepAliveOnce(context.TODO(), resp.ID)
	if err != nil {
		fmt.Println("keepalive error")
		panic(err)

	}
	fmt.Println("ttl: ", ka.TTL)

	for {
		rsp, err := kv.Get(context.TODO(), "/cron/jobs/job1")
		if err != nil {
			panic(err)
		}
		if rsp.Count == 0 {
			fmt.Println("过期了。。")
			break
		}
		fmt.Println("获取数据", rsp.Kvs)
		time.Sleep(2*time.Second)
	}
}