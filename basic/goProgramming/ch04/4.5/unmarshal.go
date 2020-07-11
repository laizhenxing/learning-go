package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
	// 当前时间
	now := time.Now().Unix()
	// 一个月前
	preMonth := now - 24*30*3600
	// 一年前
	preYear := now - 365*24*3600

	nowRes := make([]*Issue, 0)
	monthRes := make([]*Issue, 0)
	yearRes := make([]*Issue, 0)

	for _, item := range result.Items {
		createdTime := item.CreatedAt.Unix()
		if createdTime < preMonth && createdTime >= preYear {
			// 一个月前
			monthRes = append(monthRes, item)
		} else if createdTime < preYear {
			 // 一年前
			yearRes = append(yearRes, item)
		} else {
			nowRes = append(nowRes, item)
		}
	}
	fmt.Println("-------------------------------nowRes-----------------------------")
	print(nowRes)
	fmt.Println("-------------------------------monthRes-----------------------------")
	print(monthRes)
	fmt.Println("-------------------------------yearRes-----------------------------")
	print(yearRes)
}

func print(result []*Issue)  {
	for _, item := range result {
		fmt.Printf("#%-5d %9.9s %.55s \t%s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt.String())
	}
}
const IssuesURL = "https://api.github.com/search/issues?q=windows+label:bug+language:python+state:open&sort=created&order=asc"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string 
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	//q := url.QueryEscape(strings.Join(terms, " "))
	rsp, err := http.Get(IssuesURL)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", rsp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}