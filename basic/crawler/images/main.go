package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"crawler/helpers"
)

// 下载图片，出入的图片名称
func DownloadFile(url, filename string) bool {
	rsp, err := http.Get(url)
	helpers.HandlerError(err, "images.main http.Get image")
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	helpers.HandlerError(err, "images.main ioutil.ReadAll")
	filename = "./img/" + filename
	// 写入数据
	err = ioutil.WriteFile(filename, bytes, 0655)
	if err != nil {
		return false
	}

	return true
}

// 并发爬虫思路
// 1. 初始化数据管道
// 2. 爬虫写出：26个协程向管道中添加图片链接
// 3. 任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4. 下载协程：从管道里读取链接并下载

var (
	chanImageUrls chan string
	wg sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func main() {
	// 下载测试
	//DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201909/10242/e5794daf58_4.jpg", "1.jpg")

	// 1. 初始化channel
	chanImageUrls = make(chan string, 10000)
	chanTask = make(chan string,  26)
	// 2.爬虫协程
	for i := 1; i < 27; i++ {
		wg.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	// 3.任务统计协程，统计26个任务是否都完成，完成则关闭管道

	wg.Add(1)
	go CheckOk()
	// 4.协程下载: 从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}

// 下载图片
func DownloadImg()  {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", url)
		} else {
			fmt.Printf("%s 下载失败\n", url)
		}
	}
	wg.Done()
}

// 截取url中图片名字
func GetFilenameFromUrl(url string) string {
	// 返回最后一个/位置
	lastIndex := strings.LastIndex(url,"/")
	// 切割
	filename := url[lastIndex+1:]
	// 加时间戳重命名
	timePrefix := strconv.Itoa(int(time.Now().Unix()))
	filename = timePrefix + "_" + filename
	return filename
}

// 爬取图片链接到管道
// url s是整个页面的链接
func getImgUrls(url string)  {
	urls := getImgs(url)
	// 遍历切片，将数据存入管道
	for _, r := range urls {
		chanImageUrls <- r
	}
	// 标记当前协程
	// 每完成一个任务，写入一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

// 获取页面中图片的链接
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	res := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(res))
	for _, r := range res {
		url := r[0]
		urls = append(urls, url)
	}
	return
}

func GetPageStr(url string) (pageStr string) {
	rsp, err := http.Get(url)
	helpers.HandlerError(err, "images.main http.Get GetPageStr")
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	helpers.HandlerError(err, "images.main ioutil.ReadAll")
	pageStr = string(bytes)
	return
}

// 检测热舞是否完成
func CheckOk()  {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}