package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Reader struct {
	s string
	n int
}

func (r *Reader) Read(p []byte) (n int, err error) {
	//fmt.Println("interface main.Reader")
	if r.n >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(p, r.s[n:])
	r.n += n
	return n, nil
}

func NewReader(s string) *Reader {
	return &Reader{
		s: s,
		n: 0,
	}
}

type LimitR struct {
	R io.Reader
	N int64
}

func (l *LimitR) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(n) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitR{
		R: r,
		N: n,
	}
}

func main()  {
	//r := NewReader("12345678910111213141516")
	//s, err := ioutil.ReadAll(r)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "%s\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println(string(s))
	//Copy()
	// limitReader test
	limitReaderTest()
}

func limitReaderTest()  {
	r := NewReader("qbcdkljouhihjpuiyhrwejnkjshgdf;agfasdfjadhfok")
	nr := LimitReader(r, 4)
	ns, err := ioutil.ReadAll(nr)
	if err != nil {
		fmt.Println("read err", err)
	}
	fmt.Println(string(ns))
}

func Copy()  {
	//s1 := []int{1, 2, 3, 4, 5}
	//s2 := []int{6, 7, 8}
	//n := copy(s2, s1)
	c1 := "test123456test"
	c2 := make([]byte, 10)
	n := copy(c2, c1[3:])
	fmt.Println(n)
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c2 string", string(c2))
}
