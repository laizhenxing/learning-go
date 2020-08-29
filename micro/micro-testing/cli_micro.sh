#!/usr/bin/bash
baseCmd=$(micro --registry etcd --registry_address 127.0.0.1:2380)
# 列出服务
micro --registry etcd --registry_address 127.0.0.1:2380 list services

# 获取某个服务的信息
micro --registry etcd --registry_address 127.0.0.1:2380 get service hellomicro

# 调用服务 Endpoint(hellomicro HelloMicroService.HelloMicro) 参数（｛"id":2｝）
micro --registry etcd --registry_address 127.0.0.1:2380 call hellomicro HelloMicroService.HelloMicro "{\"id\": 2}"