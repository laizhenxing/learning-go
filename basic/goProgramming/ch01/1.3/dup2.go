// dup2 打印输入的数量和文本。从标准输入或者文件读取。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileArg := range files {
			f, err := os.Open(fileArg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	fmt.Println("count\tfile\tstr")
	for line, n := range counts {
		if n > 1 {
			fs := strings.Split(line, "&")
			fmt.Printf("%d\t%s\t%s\n", n, fs[0], fs[1])
		}
	}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// 控制退出
		if input.Text() == "end" {
			break
		}
		counts[f.Name() + "&" + input.Text()]++
	}
}
