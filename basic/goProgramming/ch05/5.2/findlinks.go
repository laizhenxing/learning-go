package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, doc *html.Node) []string {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, a := range doc.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if doc.FirstChild != nil {
		links = visit(links, doc.FirstChild)
	}
	if doc.NextSibling != nil {
		links = visit(links, doc.NextSibling)
	}
	//for c := doc.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}

	return links
}
