package util

import (
	"log"
	"strconv"

	"github.com/google/uuid"
	csapi "github.com/hashicorp/consul/api"

)

const IP = "192.168.37.132"
var (
	ConsulClient *csapi.Client
	ServiceID    string
	ServiceName  string
	ServicePort  int
)

func init() {
	config := csapi.DefaultConfig()
	config.Address = IP + ":8500"

	// 创建客户端
	client, err := csapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
	// 服务配置
	ServiceID = "userService" + uuid.New().String()
}

func SetServiceNameAndPort(name string, port int)  {
	ServiceName = name
	ServicePort = port
}

func RegisterService() {
	// 服务注册信息
	reg := csapi.AgentServiceRegistration{
		ID:      ServiceID,
		Name:    ServiceName,
		Port:    ServicePort,
		Address: IP,
		Tags:    []string{"primary"},
	}
	// 健康检查
	check := csapi.AgentServiceCheck{
		Interval: "5s",
		HTTP:     "http://" + IP + ":" + strconv.Itoa(ServicePort) + "/health",
	}
	reg.Check = &check

	// 注册
	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func UnRegisterService() {
	ConsulClient.Agent().ServiceDeregister(ServiceID)
}
