package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)


func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countNode: %s", err)
		os.Exit(1)
	}
	countNodeRes := make(map[string]int)
	countNode(countNodeRes, doc)
	fmt.Println(countNodeRes)
}

func countNode(res map[string]int, doc *html.Node) {
	//res[doc.Data]++
	if doc.Type == html.ElementNode {
		res[doc.Data]++
	}
	if doc.FirstChild != nil {
		countNode(res, doc.FirstChild)
	}
	if doc.NextSibling != nil {
		countNode(res, doc.NextSibling)
	}
	//for c := doc.FirstChild; c != nil; c = c.NextSibling {
	//	res = countNode(res, c)
	//}

	return
}
