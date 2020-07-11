package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type resource struct {
	url    string
	target string
	start  int
	end    int
}

var UserAgentList =  []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
	"Opera/8.0 (Windows NT 5.1; U; en)",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.50",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0",
	"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 SE 2.X MetaSr 1.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/4.4.3.4000 Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 UBrowser/4.0.3214.0 Safari/537.36",
}

func main() {
	total := flag.Int("total", 100, "how many log you will make.")
	filePath := flag.String("filePath", "F:\\phpstudy_pro\\WWW\\nginxLog\\dig.log", "the path will the file be stored.")
	flag.Parse()
	fmt.Println(*total, *filePath)

	// 构造出真实的网站url集合
	res := ruleResource()
	list := buildUrl(res)
	//fmt.Println(list)

	logStr := ""
	// 按照要求，生成 total 行日志内容
	for i := 0; i < *total; i++ {
		currentUrl := list[randInt(0, len(list)-1)]
		referUrl := list[randInt(0, len(list)-1)]
		ua := UserAgentList[randInt(0, len(UserAgentList)-1)]
		logStr = logStr + makeLog(currentUrl, referUrl, ua) + "\n"

	}
	// 打开文件
	fd, err := os.OpenFile(*filePath, os.O_RDWR|os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	fd.Write([]byte(logStr))
}

func ruleResource() []resource {
	var res []resource
	// 首页
	r1 := resource{
		url:    "http://gxcms/",
		target: "",
		start:  0,
		end:    0,
	}
	// 列表页
	r2 := resource{
		url:    "http://gxcms/list/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    21,
	}
	r3 := resource{
		url:    "http://gxcms/movie/{$id}.html",
		target: "{$id}",
		start:  1,
		end:    12924,
	}
	res = append(res, r1, r2, r3)
	return res
}

func buildUrl(res []resource) (list []string) {
	for _, r := range res {
		if len(r.target) == 0 {
			list = append(list, r.url)
		} else {
			for i := r.start; i <= r.end; i++ {
				strUrl := strings.Replace(r.url, r.target, strconv.Itoa(i), -1)
				list = append(list, strUrl)
			}
		}
	}

	return
}

func makeLog(current, refer, ua string) string {
	u:=url.Values{}
	u.Set("time",time.Now().Format("2006-01-02 15:04:05"))
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramsStr := u.Encode()

	logTemplate := "127.0.0.1 - - [12/Jun/2020:17:58:20 +0800] \"GET /index.php?{$paramsStr} HTTP/1.1\" 200 15 \"{$refer}\" \"{$ua}\""
	log := strings.Replace(logTemplate, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "{$refer}", refer, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)

	return log
}

// 随机获取一个整数
func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}

	return r.Intn(max-min) + min
}