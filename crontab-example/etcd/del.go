package main

import (
	"context"
	"fmt"
	"time"

	v3 "go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		cfg v3.Config
		cli *v3.Client
		err error
		rsp *v3.DeleteResponse
	)

	cfg = v3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: 5 * time.Second}
	if cli, err = v3.New(cfg); err != nil {
		panic(err)
	}
	defer cli.Close()

	kv := v3.NewKV(cli)

	// 删除一个key
	//if 	rsp, err = kv.Delete(context.TODO(), "k", v3.WithPrevKV()); err != nil {
	//	fmt.Println("single delete error")
	//	panic(err)
	//}
	//if rsp.PrevKvs != nil {
	//	for _, v := range rsp.PrevKvs {
	//		fmt.Println("k: ", string(v.Key), ", v: ", string(v.Value), ", modRevision: ", v.ModRevision)
	//	}
	//}

	// 删除多个key
	// limit指定大于0的数会报错
	if rsp, err = kv.Delete(context.TODO(), "cron/jobs/", v3.WithPrevKV(), v3.WithFromKey(), v3.WithLimit(0)); err != nil {
		fmt.Println("multi delete error")
		panic(err)
	}
	if len(rsp.PrevKvs) != 0 {
		for _, v := range rsp.PrevKvs {
			fmt.Println("k: ", string(v.Key), ", v: ", string(v.Value), ", modRevision: ", v.ModRevision)
		}
	}
}