package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mgutz/str"
	"github.com/sirupsen/logrus"

	"analysis/helpers"
)

const (
	HandleDig   = "/index.php?"
	HandleMovie = "/movie/"
	HandleList  = "/list/"
	HandleHtml  = ".html"
)

type cmdParams struct {
	logFilePath string
	routineNum  int
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data  digData
	uid   string
	unode urlNode
}

type urlNode struct {
	unType string // 详情页或者列表页或者首页
	unRid  int    // ResourceID
	unUrl  string // 当前页面的url
	unTime string // 当前页面的访问时间
}

type storageBlock struct {
	counterType string
	storageMode string
	uNode       urlNode
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	// 获取参数
	logFilePath := flag.String("logFilePath", "F:\\phpstudy_pro\\WWW\\nginxLog\\dig.log", "log filepath.")
	// 并发数量控制
	routineNum := flag.Int("routineNum", 5, "consume number by goroutine")
	// 日志存储路径
	logSavePath := flag.String("logSavePath", "F:\\Code\\Go\\learning-go\\analysis\\runtime\\log.txt", "log save path.")
	flag.Parse()

	params := cmdParams{
		logFilePath: *logFilePath,
		routineNum:  *routineNum,
	}

	// 打日志
	logFd, err := os.OpenFile(*logSavePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer logFd.Close()
	log.Out = logFd
	log.Infoln("Exec start.")
	log.Infof("Params: logFilePath=%s, routineNum=%d", params.logFilePath, params.routineNum)

	// 初始化一些channel,用于数据传递
	var (
		logChannel     = make(chan string, 3*params.routineNum)
		pvChannel      = make(chan urlData, params.routineNum)
		uvChannel      = make(chan urlData, params.routineNum)
		storageChannel = make(chan storageBlock, params.routineNum)
	)

	// Redis Pool
	redisPool, err := pool.New("tcp", "127.0.0.1:6379", 2*params.routineNum)
	if err != nil {
		log.Infof("RedisPool: redis connect error.")
		panic(err)
	} else {
		go func() {
			redisPool.Cmd("PING")
			time.Sleep(3 * time.Second)
		}()
	}

	// 日志消费者
	go readFileLineByLine(params, logChannel)

	// 创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	// 创建 PV UV 统计器
	// pvChannel：输入， storageChannel: 输出
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel, redisPool)

	// 创建数据存储器
	go dataStorage(storageChannel, redisPool)

	time.Sleep(1000 * time.Second)
}

// 日志消费者
func readFileLineByLine(params cmdParams, logChan chan string) error {
	// 打开文件
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Infoln("readFileLineByLine: Open file failed.")
		return err
	}
	defer fd.Close()

	// 逐行读取文件
	count := 0
	bufferRead := bufio.NewReader(fd)
	for {
		line, err := bufferRead.ReadString('\n')
		log.Infof("ReadFileLineByLine line string: %s", line)
		if err != nil {
			if err == io.EOF { // 文件已经读取完毕
				time.Sleep(3 * time.Second)
				log.Infof("ReadFileLineByLine: readline wait...")
			} else {
				log.Warningf("ReadFileLineByLine: read log error")
			}
		}

		logChan <- line
		count++
		if count%(1000*params.routineNum) == 0 {
			log.Infof("ReadFileLineByLine line: %d", count)
		}
	}

	return nil
}

// 日志解析器
func logConsumer(logChan chan string, pvChan, uvChan chan urlData) error {
	for logStr := range logChan {
		// 切割字符串，抠出打点的数据
		data := cutLogFetchData(logStr)

		// 构造uid md5(refer + ud)
		hasher := md5.New()
		hasher.Write([]byte(data.refer + data.ua))
		uid := hex.EncodeToString(hasher.Sum(nil))

		uData := urlData{
			data:  data,
			uid:   uid,
			unode: formatUrl(data.url, data.time),
		}

		pvChan <- uData
		uvChan <- uData
	}

	return nil
}

func cutLogFetchData(logStr string) digData {
	logStr = strings.TrimSpace(logStr)

	pos1 := str.IndexOf(logStr, HandleDig, 0)
	if pos1 == -1 { // 没找到
		return digData{}
	}
	pos1 += len(HandleDig)
	pos2 := str.IndexOf(logStr, "HTTP/", pos1)
	d := str.Substr(logStr, pos1, pos2-pos1)

	urlInfo, err := url.Parse("http://gxcms/?" + d)
	if err != nil {
		log.Errorln("urlInfo read error: ", err)
		return digData{}
	}
	data := urlInfo.Query()
	return digData{
		data.Get("time"),
		data.Get("url"),
		data.Get("refer"),
		data.Get("ua"),
	}
}

// pv 统计器
func pvCounter(pvChan chan urlData, storageChan chan storageBlock) {
	for data := range pvChan {
		sItem := storageBlock{
			counterType: "pv",
			storageMode: "ZINCRBY",
			uNode:       data.unode,
		}
		storageChan <- sItem
	}
}

// uv 统计器 [需要去重]
func uvCounter(uvChan chan urlData, storageChan chan storageBlock, redisPool *pool.Pool) {
	for data := range uvChan {
		// HeyperLoglog redis - 去重操作
		hyperLogLogKey := "uv_hpll_" + helpers.GetTime(data.data.time, helpers.Day)
		// PFADD redis超级日志
		ret, err := redisPool.Cmd("PFADD", hyperLogLogKey, data.uid, "EX", 86400).Int()
		if err != nil {
			log.Warningln("uvCounter check redis hyperLogLog error.")
		}
		if ret != 1 {
			continue
		}

		sItem := storageBlock{
			counterType: "uv",
			storageMode: "ZINCRBY",
			uNode:       data.unode,
		}
		storageChan <- sItem
	}
}

// 数据存储器
func dataStorage(storageChan chan storageBlock, redisPool *pool.Pool) {
	for block := range storageChan {
		prefix := block.counterType + "_"

		// 逐层添加，加洋葱的过程
		// 维度：天-小时-分钟
		// 层级：顶级-大分类-小分类-终极页面
		// 存储模型： Redis SortedSet
		setKeys := []string{
			prefix + "day_" + helpers.GetTime(block.uNode.unTime, helpers.Day),
			prefix + "hour_" + helpers.GetTime(block.uNode.unTime, helpers.Hour),
			prefix + "min_" + helpers.GetTime(block.uNode.unTime, helpers.Day),
			prefix + block.uNode.unType + "_day_" + helpers.GetTime(block.uNode.unTime, helpers.Day),
			prefix + block.uNode.unType + "_hour_" + helpers.GetTime(block.uNode.unTime, helpers.Hour),
			prefix + block.uNode.unType + "_min_" + helpers.GetTime(block.uNode.unTime, helpers.Min),
		}

		rowId := block.uNode.unRid
		for _, key := range setKeys {
			//fmt.Println("block.storageMode", block.storageMode, key)
			ret, err := redisPool.Cmd(block.storageMode, key, 1, rowId).Int()
			//fmt.Println("ret: ", ret, err)
			if ret <= 0 || err != nil {
				log.Errorln("DataStorage redis storage error. ", block.storageMode, key, rowId)
			}
		}
	}
}

// 解析url，构造urlNode
func formatUrl(url, t string) urlNode {
	// 从大量的着手，详情页>列表页>=首页
	pos1 := str.IndexOf(url, HandleMovie, 0)
	if pos1 != -1 {
		// 详情页
		pos1 += len(HandleMovie)
		pos2 := str.IndexOf(url, HandleHtml, 0)
		idStr := str.Substr(url, pos1, pos2-pos1)
		id, _ := strconv.Atoi(idStr)
		return urlNode{"movie", id, url, t}
	} else {
		// 列表页
		pos1 = str.IndexOf(url, HandleList, 0)
		if pos1 != -1 {
			pos1 += len(HandleList)
			pos2 := str.IndexOf(url, HandleHtml, 0)
			idStr := str.Substr(url, pos1, pos2-pos1)
			id, _ := strconv.Atoi(idStr)
			return urlNode{"list", id, url, t}
		} else {
			// 首页
			return urlNode{"home", 1, url, t}
		}
	}
}
