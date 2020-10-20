#### go-kit 三层架构

##### 1. Transport

     主要负责与 HTTP、gRPC、 thrift 等相关的逻辑

##### 2. Endpoint

    定义 Request 和 Response 格式，并可以使用装饰器包装函数，以此来实现各种中间件嵌套

##### 3. Service

    业务类、接口等

#### etcd

    etcd 是使用 Go 语言开发的一个开源的、高可用的分布式 key-value 存储系统，可以用于配置共享和服务的注册与发现。

#### consul

    docker 启动命令： docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0
    -server: 代表以服务器方式启动
    -bootstrap: 指定自己为 leader, 而不需要选举
    -ui: 启动一个内置管理的web 界面
    -client: 指定客户端可以访问的ip,设置0.0.0.0则任意访问

##### 手动注册一个服务

  1. 创建一个json文件
    `register-consul.json`:

         {
           "ID": "userservice",
           "Name": "userservice",
           "Tags": [
             "primary"
           ],
           "Address": "192.68.37.132",
           "Port": 8080,
           "Check": {
               "HTTP": "http://192.168.37.132:8080/health",
               "Interval": "5s"
           }
         }

  2. 使用 API 接口进行注册 

    curl --request PUT --data @register-consul.json http://localhost:8500/v1/agent/service/register


  3. 查看注册的服务
  `http://192.168.37.132:8500/ui/dc1/services`

  4. 反注册（注销）服务
  
    PUT /v1/agent/service/deregister/:server_id application/json


#### API 限流

##### 令牌桶算法
  
##### 内置包(rate)核心方法

1. Wait/WaitN

    `WaitN.go`

    ```
      r := rate.NewLimiter(1, 5)
      ctx := context.Background()
      for {
        // 每次消耗两个，否则等待
        err := r.WaitN(ctx, 2)
        if err != nil {
          log.Fatal(err)
        }
        fmt.Println(time.Now().Format("2006-01-02 15:04:05))
        time.Sleep(time.Second)
      }
    ```

2. Allow/AllowN
3. Reserve/ReserveN

#### 熔断器

##### 第三方库 `https://github.com/afex/hystrix-go`

> 本质是：隔离远程服务请求，防止级联故障

##### 熔断器的状态

    1. 关闭：默认状态。如果请求次数异常超过设定比例，则打开熔断器。
    2. 打开：当熔断器打开的时候，直接执行降级方法。
    3. 半开：定期地尝试发起请求来确认系统是否恢复。如果恢复了，熔断器将转为关闭状态或者保持打开。通过设置 SleepWindow 来设置尝试请求间隔时间。默认5s.
