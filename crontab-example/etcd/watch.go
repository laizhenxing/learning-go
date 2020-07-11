package main

import (
	"context"
	"fmt"
	"time"

	v3 "go.etcd.io/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func main() {
	var (
		err error
		cli *v3.Client
	)

	cfg := v3.Config{Endpoints: []string{"127.0.0.1:2379"}, DialTimeout: 5*time.Second}
	cli, err = v3.New(cfg)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	kv := v3.NewKV(cli)
	// 先 GET 到当前的值，并监听后续变化
	grsp, err := kv.Get(context.TODO(), "cron/jobs/job1")
	if err != nil {
		panic(err)
	}

	// 如果值存在
	if len(grsp.Kvs) != 0 {
		fmt.Println("当前值：", string(grsp.Kvs[0].Value))
	}

	//当前etcd集群事务id,单调递增
	wsId := grsp.Header.Revision + 1

	// 创建一个watcher
	watcher := v3.NewWatcher(cli)

	// 启动监听
	fmt.Println("从该版本向后监听：", wsId)
	wtChan := watcher.Watch(context.TODO(), "cron/jobs/job1", v3.WithRev(wsId))
	for ch := range wtChan {
		for _, ev := range ch.Events {
			switch ev.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(ev.Kv.Value), ev.Kv.CreateRevision, ev.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了; Revision: ", ev.Kv.ModRevision)
			}
		}
	}
}
