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

	"analysis02/helpers"
)

const (
	HandleLog = "/index.php?"
	HandleMovie= "/movie/"
	HandleList= "/list/"
	HandleHtml= ".html"
)

type cmdParams struct {
	filePath   string
	routineNum int
}

type logData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data  logData
	uid   string
	uNode urlNode
}

type urlNode struct {
	unType string
	unRid  int
	unUrl  string
	unTime string
}

type storageBlock struct {
	counterType string
	storageMode string // redis的存储模式
	uNode       urlNode
}

var logSavePath = "F:\\Code\\Go\\learning-go\\analysis02\\runtime\\log.txt"

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	// 获取命令行输入的参数
	filePath := flag.String("filePath", "F:\\phpstudy_pro\\WWW\\nginxLog\\dig.log", "the log filepath")
	routineNum := flag.Int("num", 5, "goroutine nums")
	// 解析参数，否则不生效
	flag.Parse()
	// 构造命令行参数
	params := cmdParams{
		filePath:   *filePath,
		routineNum: *routineNum,
	}

	// 打开日志记录文件
	fd, err := os.OpenFile(logSavePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	// 设置日志的输出位置
	log.Out = fd
	// 开始记录日志
	log.Infoln("analysis02 log start...")

	// 建立redis连接池
	redisPool, err := pool.New("tcp", "127.0.0.1:6379", 2 * params.routineNum)
	if err != nil {
		log.Errorln("[main] create redis pool error: ", err)
		panic(err)
	} else {
		// 保持连接池的持久访问
		go func() {
			redisPool.Cmd("PING")
			time.Sleep(3 * time.Second)
		}()
	}

	// 创建 channels
	var (
		logChan     = make(chan string, 3*params.routineNum)
		pvChan      = make(chan urlData, 3*params.routineNum)
		uvChan      = make(chan urlData, 3*params.routineNum)
		storageChan = make(chan storageBlock, 3*params.routineNum)
	)

	// 日志消费
	go readFileLineByLine(params, logChan)

	// 构造一组日志解析器
	for i := 0; i < params.routineNum; i++ {
		go parseUrlString(logChan, pvChan, uvChan)
	}

	// pv,uv 计数器
	go pvCounter(pvChan, storageChan)
	go uvCounter(uvChan, storageChan, redisPool)

	// 存储器
	go storageData(storageChan, redisPool)

	time.Sleep(time.Second)
}

func storageData(storageChan chan storageBlock, redisPool *pool.Pool) {
	// 逐层添加，加洋葱
	// 维度：天-时-分-秒
	// 层级：顶级-大分类-小分类-详情页
	// 存储模型： Redis SortedSet
	for data := range storageChan {
		prefix := data.counterType + "_"
		t := data.uNode.unTime
		keys := []string{
			prefix + "day_" + helpers.GetTime(t, helpers.Day),
			prefix + "hour_" + helpers.GetTime(t, helpers.Hour),
			prefix + "min_" + helpers.GetTime(t, helpers.Min),
			prefix + data.uNode.unType + "_day_" + helpers.GetTime(t, helpers.Day),
			prefix + data.uNode.unType + "_hour_" + helpers.GetTime(t, helpers.Hour),
			prefix + data.uNode.unType + "_min_" + helpers.GetTime(t, helpers.Min),
		}

		rowId := data.uNode.unRid
		for _, key := range keys {
			ret, err := redisPool.Cmd(data.storageMode, key, 1, rowId).Int()
			if ret <= 0 || err != nil {
				log.Errorln("[storageData] store log data to redis failed! errors: ", err)
			}
		}
	}
}

func uvCounter(uvChan chan urlData, storageChan chan storageBlock, redisPool *pool.Pool) {
	log.Infoln("[uvCounter] uv data parse start...")
	// 去重 PFADD
	for uv := range uvChan {
		hypeLogKey := "uv_hpll_" + helpers.GetTime(uv.data.time, helpers.Second)
		// PFADD redis超级日志 [写入redis]
		ret, err := redisPool.Cmd("PFADD", hypeLogKey, uv.uid, "EX", 86400).Int()

		if err != nil {
			log.Errorln("[uvCounter] check redis hypeLog error")
		}
		if ret != -1 {
			// 如果已存在记录则执行下一个循环，去重
			continue
		}

		stBl := storageBlock{
			counterType: "pv",
			storageMode: "ZINCRBY",
			uNode:       uv.uNode,
		}

		storageChan <- stBl
	}
}

func pvCounter(pvChan chan urlData, storageChan chan storageBlock) {
	for pv := range pvChan {
		stBl := storageBlock{
			counterType: "pv",
			storageMode: "ZINCRBY",
			uNode:       pv.uNode,
		}

		storageChan <- stBl
	}
}

func parseUrlString(logChan chan string, pvChan, uvChan chan urlData) error {
	log.Infoln("[parseUrlString] parser url start...")
	for urlStr := range logChan {
		lData := cutFetchUrlString(urlStr)

		//构造uid md5(refer+ua)
		hasher := md5.New()
		hasher.Write([]byte(lData.refer + lData.ua))
		uid := hex.EncodeToString(hasher.Sum(nil))
		uData := urlData{
			data:  lData,
			uid:   uid,
			uNode: formatUrl(lData.url, lData.time),
		}

		pvChan <- uData
		uvChan <- uData
	}

	return nil
}

func formatUrl(url string, t string) urlNode {
	pos1 := str.IndexOf(url, HandleMovie, 0)
	if pos1 != -1 {
		// 详情页
		pos1 += len(HandleMovie)
		pos2 := str.IndexOf(url, HandleHtml, pos1)
		idStr := str.Substr(url, pos1, pos2-pos1)
		id, _ := strconv.Atoi(idStr)
		return urlNode{
			unType: "movie",
			unRid:  id,
			unUrl:  url,
			unTime: t,
		}
	} else {
		pos1 = str.IndexOf(url, HandleList, 0)
		if pos1 != -1 {
			// 列表页
			pos1 += len(HandleList)
			pos2 := str.IndexOf(url, HandleHtml, pos1)
			idStr := str.Substr(url, pos1, pos2-pos1)
			id, _ := strconv.Atoi(idStr)
			return urlNode{
				unType: "list",
				unRid:  id,
				unUrl:  url,
				unTime: t,
			}
		} else {
			// 首页
			return urlNode{
				unType: "home",
				unRid:  1,
				unUrl:  url,
				unTime: t,
			}
		}
	}
}

func cutFetchUrlString(urlStr string) logData {
	// 除去字符串的空格
	urlStr = strings.TrimSpace(urlStr)

	pos1 := str.IndexOf(urlStr, HandleLog, 0)
	if pos1 == -1 {
		// 没有找到开头的字符串直接返回
		return logData{}
	}
	pos1 += len(HandleLog)
	pos2 := str.IndexOf(urlStr, "HTTP/", pos1)
	// 截取字符串
	fnStr := str.Substr(urlStr, pos1, pos2-pos1)

	// 解析该url字符串
	urlInfo, err := url.Parse("http://gxcms/?" + fnStr)

	if err != nil {
		log.Infoln("[cutFetchUrlString] parse url info error: ", err)
		return logData{}
	}
	data := urlInfo.Query()
	return logData{
		time:  data.Get("time"),
		url:   data.Get("url"),
		refer: data.Get("refer"),
		ua:    data.Get("ua"),
	}
}

func readFileLineByLine(params cmdParams, logChan chan string) error {
	// 打开文件
	fd, err := os.Open(params.filePath)
	if err != nil {
		// 打开日志文件失败
		log.Infof("[readFileLineByLine] open logFile error: %s. %#v", params.filePath, err)
		return err
	}
	defer fd.Close()

	log.Infoln("[readFileLineByLine] reading line start...")
	buffer := bufio.NewReader(fd)
	counter := 0
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			if err == io.EOF { // 文件已读完
				// 等待3s后继续读取文件
				time.Sleep(3 * time.Second)
				log.Infoln("[readFileLineByLine] waiting for reading log file...")
			} else { // 发生错误
				log.Errorln("[readFileLineByLine] reading log file error: ", err)
			}
		}
		logChan <- line
		counter++

		// 每读取1000*routineNum行记录一条日志
		if counter%(1000*params.routineNum) == 0 {
			log.Infof("[readFileLineByLine] read line %d.\n", counter)
		}
	}

	return nil
}
