package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

// tcp扫描器，扫描端口
func main() {
	// 命令行输入参数 domain,start-port,end-port,timeout
	domain := flag.String("h", "", "hostname")
	startPort := flag.Int("sp", 80, "start port")
	endPort := flag.Int("ep", 100, "end port")
	timeout := flag.Duration("t", 1*time.Second, "timeout")
	flag.Parse()

	dPorts := make([]int, 0)
	var wg sync.WaitGroup
	// 用互斥锁来解决竞争条件
	mutex := sync.Mutex{}

	for p := *startPort; p <= *endPort; p++ {
		wg.Add(1)
		go func(port int) {
			ok := isOpen(*domain, port, *timeout)
			if ok {
				mutex.Lock()
				dPorts = append(dPorts, port)
				mutex.Unlock()
			}
			wg.Done()
		}(p)
	}

	wg.Wait()
	fmt.Println("open ports: ", dPorts)
}

func isOpen(domain string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", domain, port), timeout)
	if err != nil {
		//fmt.Println(err)
		return false
	}
	conn.Close()
	return true
}
