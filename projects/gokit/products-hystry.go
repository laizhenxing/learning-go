package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

type Product struct {
	ID    int
	Title string
	Price int
}

func main_1() {
	// 设置随机因子
	rand.Seed(time.Now().UnixNano())
	for {
		p, _ := getProduct()
		fmt.Println(p)
		time.Sleep(time.Second * 1)
	}
}

// 使用熔断器做超时处理
func main_2() {
	rand.Seed(time.Now().UnixNano())

	// 设置配置，自定义配置超时时间4s
	configA := hystrix.CommandConfig{
		Timeout: 4000,
	}
	// 配置与命令绑定
	hystrix.ConfigureCommand("get_prod", configA)
	// 配置多个命令
	hystrix.ConfigureCommand("get_prod2", hystrix.CommandConfig{
		Timeout: 2000, // 超时2s
	})
	for {
		err := hystrix.Do("get_prod2", func() error {
			p, _ := getProduct()
			fmt.Println(p)
			return nil
		}, nil)
		if err != nil {
			fmt.Println("Do hystrix err: ", err)
		}
		time.Sleep(time.Second * 1)
	}
}

// 服务降级
func main_3() {
	rand.Seed(time.Now().UnixNano())

	// 设置配置，自定义配置超时时间4s
	configA := hystrix.CommandConfig{
		Timeout: 4000,
	}
	// 配置与命令绑定
	hystrix.ConfigureCommand("get_prod", configA)
	// 配置多个命令
	hystrix.ConfigureCommand("get_prod2", hystrix.CommandConfig{
		Timeout: 2000, // 超时2s
	})
	for {
		err := hystrix.Do("get_prod2", func() error {
			p, _ := getProduct()
			fmt.Println(p)
			return nil
		}, func(err error) error {
			fmt.Println(defaultProduct())
			return errors.New("服务降级指定错误.")
		})
		if err != nil {
			fmt.Println("Do hystrix err: ", err)
		}
		time.Sleep(time.Second * 1)
	}
}

// 异步执行
func main_4() {
	rand.Seed(time.Now().UnixNano())

	// 设置配置，自定义配置超时时间4s
	configA := hystrix.CommandConfig{
		Timeout: 2000,
	}
	// 配置与命令绑定
	hystrix.ConfigureCommand("get_prod", configA)

	// 定义存储结果的 channel
	ch := make(chan Product, 1)
	for {
		err := hystrix.Go("get_prod", func() error {
			p, _ := getProduct()
			//fmt.Println(p)
			ch <- p
			return nil
		}, func(err error) error {
			//fmt.Println(defaultProduct())
			//ch <- defaultProduct()
			return errors.New("服务降级指定错误.")
		})

		// 等待输出结果
		select {
		case prod := <-ch:
			fmt.Println(prod)
		case e := <-err:
			fmt.Println(e)
		}

		time.Sleep(time.Second * 1)
	}
}

// 指定最大并发数
func main_5() {
	rand.Seed(time.Now().UnixNano())

	// 设置配置，自定义配置超时时间4s
	configA := hystrix.CommandConfig{
		Timeout:               2000, // 指定超时时间2s
		MaxConcurrentRequests: 5,    // 指定最大并发数
	}
	// 配置与命令绑定
	hystrix.ConfigureCommand("get_prod", configA)

	// 定义存储结果的 channel
	ch := make(chan Product, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			err := hystrix.Go("get_prod", func() error {
				p, _ := getProduct()
				fmt.Println(p)
				ch <- p
				return nil
			}, func(err error) error {
				fmt.Println(err)
				P, _ := defaultProduct()
				ch <- P
				return errors.New("服务降级指定错误.")
			})

			// 等待输出结果
			select {
			case prod := <-ch:
				fmt.Println(prod)
			case e := <-err:
				fmt.Println(e)
			}

		}()
	}

	wg.Wait()
}

// 熔断器参数设置
func main_6() {
	rand.Seed(time.Now().UnixNano())

	// 设置配置，自定义配置超时时间4s
	configA := hystrix.CommandConfig{
		Timeout:                2000, // 指定超时时间2s
		MaxConcurrentRequests:  5,    // 指定最大并发数
		RequestVolumeThreshold: 3,    // 熔断器请求阀值,默认20
		ErrorPercentThreshold:  20,   // 错误百分比，默认50
	}
	// 配置与命令绑定
	hystrix.ConfigureCommand("get_prod", configA)
	// 获取当前服务熔断器的状态
	c, _, _ := hystrix.GetCircuit("get_prod")
	for i := 0; i < 100; i++ {
		err := hystrix.Do("get_prod", func() error {
			p, _ := getProduct()
			fmt.Println(p)
			return nil
		}, func(err error) error {
			//fmt.Println(err)
			p, e := defaultProduct()
			fmt.Println(p)
			//return errors.New("服务降级指定错误." + err.Error())
			return e
		})

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("熔断器状态： ", c.IsOpen())
		time.Sleep(time.Second * 1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 设置配置，自定义配置超时时间4s
	configA := hystrix.CommandConfig{
		Timeout:                2000, // 指定超时时间2s
		MaxConcurrentRequests:  5,    // 指定最大并发数
		RequestVolumeThreshold: 3,    // 熔断器请求阀值,默认20
		ErrorPercentThreshold:  20,   // 错误百分比，默认50
		SleepWindow:            int(time.Second * 100),	// 在熔断器开启之后，尝试请求服务是否可用间隔时间，默认5s
	}
	// 配置与命令绑定
	hystrix.ConfigureCommand("get_prod", configA)
	// 获取当前服务熔断器的状态
	c, _, _ := hystrix.GetCircuit("get_prod")
	for i := 0; i < 100; i++ {
		err := hystrix.Do("get_prod", func() error {
			p, _ := getProduct()
			fmt.Println(p)
			return nil
		}, func(err error) error {
			//fmt.Println(err)
			p, e := defaultProduct()
			fmt.Println(p)
			//return errors.New("服务降级指定错误." + err.Error())
			return e
		})

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("熔断器状态： ", c.IsOpen())
		time.Sleep(time.Second * 1)
	}
}

// 支持最大并发
func getProduct() (Product, error) {
	r := rand.Intn(10)
	// 演示随机超时
	if r < 6 {
		fmt.Println("延迟3s")
		time.Sleep(3 * time.Second)
	}

	return Product{
		ID:    1,
		Title: "product1",
		Price: r,
	}, nil
}

func defaultProduct() (Product, error) {
	return Product{
		ID:    999,
		Title: "推荐商品",
		Price: 999,
	}, nil
}
