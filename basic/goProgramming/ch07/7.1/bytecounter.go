package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int, error) {
	// 分割字符串
	s := strings.NewReader(string(p))
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	sum := 0
	for bs.Scan(){
		sum++
	}
	*w = WordsCounter(sum)
	return sum, nil
}

func main() {
	var w WordsCounter
	test := "test1 test2 words1 words2-test1"
	sum, _ := w.Write([]byte(test))
	fmt.Println(sum)
}
