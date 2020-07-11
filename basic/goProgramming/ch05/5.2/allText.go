package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "allText: %s", err)
		os.Exit(1)
	}
	res := make([]string, 0)
	res = allText(res, doc)
	fmt.Println(res)
}

func allText(res []string, doc *html.Node) []string {
	if doc.Type == html.ElementNode && doc.Data != "script" && doc.Data != "style" {
		for _, attr := range doc.Attr {
			res = append(res, attr.Val)
		}
	}
	if doc.FirstChild != nil {
		res = allText(res, doc.FirstChild)
	}
	if doc.NextSibling != nil {
		res = allText(res, doc.NextSibling)
	}
	return res
}