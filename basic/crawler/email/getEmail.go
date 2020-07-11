// 1.明确目标（确定在哪个网站搜索）
// 2.爬（爬取内容）
// 3.取（筛选想要的）
// 4.处理数据
package email

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var reQQEmail = `(\d+)@qq.com`

func GetEmail()  {
	// 爬取网页数据
	rsp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	handlerError(err, "http.Get url")
	defer rsp.Body.Close()
	// 读取页面内容
	pageBytes, err := ioutil.ReadAll(rsp.Body)
	handlerError(err, "ioutil.ReadAll")
	// 字节转成字符串
	pageStr := string(pageBytes)
	// 过滤数据，获取QQ邮箱
	regx := regexp.MustCompile(reQQEmail)
	// -1表示获取全部
	res := regx.FindAllStringSubmatch(pageStr, -1)

	// 遍历结果
	for _, r := range res {
		fmt.Printf("email: %v\t\tQQ: %v\n", r[0], r[1])
	}

}

func handlerError(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}
