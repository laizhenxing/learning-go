// 爬取 studygolang.com 的用户数据
package users

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"crawler/helpers"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt string
	Rank      string // 第n位会员
}

var (
	website       = "https://studygolang.com/"
	usersPage     = website + "users"
	favoritesPage = website + "favorites/"
	regUser       = "/user/[A-Za-z0-9_-\u4e00-\u9fa5]+"
	regEmail      = `\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`
	regCreatedAt  = `\d{4}(\-|\/|.)\d{1,2}(\-|\/|.)\d{1,2}\W\d{2}:\d{2}:\d{2}`
	regRank       = `第.*[\d]+.*位会员`
)

func GetUsers() []User {
	start := time.Now()
	users := make([]User, 0)
	userChan := make(chan User, 10)
	// 获取网页信息
	rsp, err := http.Get(usersPage)
	helpers.HandlerError(err, "http.Get users")
	defer rsp.Body.Close()

	// 将字节转化为string
	pageBytes, err := ioutil.ReadAll(rsp.Body)
	helpers.HandlerError(err, "ioutil.ReadAll users")
	pageStr := string(pageBytes)
	// 打印信息
	//fmt.Println(pageStr)
	// 匹配正则表达式
	regx := regexp.MustCompile(regUser)
	res := regx.FindAllStringSubmatch(pageStr, -1)
	fr := deduplication(res)
	//fmt.Printf("find %d user\n", len(res))
	//fmt.Println(fr, len(fr))
	for _, r := range fr {
		go GetUser(r, userChan)
		users = append(users, <-userChan)
	}

	fmt.Println(users)
	end := time.Now()
	fmt.Println("spend time: ", end.Sub(start))
	return users
}

func GetUser(url string, userChan chan User) {
	var (
		email   string
		created string
		rank    string
	)
	s := strings.Split(url, "/")
	id := s[2]
	userPage := website + url
	rsp, err := http.Get(userPage)
	helpers.HandlerError(err, "http.Get userpage")
	defer rsp.Body.Close()

	pageBytes, err := ioutil.ReadAll(rsp.Body)
	helpers.HandlerError(err, "ioutil.ReadAll userpage")
	pageStr := string(pageBytes)
	// 匹配Email
	reEmail := regexp.MustCompile(regEmail)
	rE := reEmail.FindAllStringSubmatch(pageStr, -1)
	if len(rE) > 0 {
		email = rE[0][0]
	}
	//fmt.Println(url, rE)
	// 匹配排名
	reRank := regexp.MustCompile(regRank)
	rR := reRank.FindAllStringSubmatch(pageStr, -1)
	if len(rR) > 0 {
		email = rR[0][0]
	}
	//fmt.Println(url, rR)
	// 匹配日期
	reTime := regexp.MustCompile(regCreatedAt)
	rT := reTime.FindAllStringSubmatch(pageStr, -1)
	if len(rT) > 0 {
		email = rT[0][0]
	}
	//fmt.Println(url, rT)

	userChan <- User{
		ID:        id,
		Name:      "",
		Email:     email,
		CreatedAt: created,
		Rank:      rank,
	}
}

func GetFavorites(uid string) {

}

// 去重
func deduplication(data [][]string) map[string]string {
	res := make(map[string]string)
	for _, r := range data {
		//fmt.Printf("user: %v\n", r[0])
		res[r[0]] = r[0]
	}

	return res
}
