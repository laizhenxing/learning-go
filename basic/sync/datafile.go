package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	RSN() int64
	// 获取最后写入的数据块的序列号
	WRN() int64
	// 获取数据块的长度
	DataLen() uint32
	// 关闭数据文件
	Close() error
}

type Data []byte

type myDataFile struct {
	f       *os.File     // 文件
	fmutex  sync.RWMutex // 文件读写锁
	woffset int64        // 写操作用到的偏移量
	roffset int64        // 读操作用到的偏移量
	wmutex  sync.Mutex   // 写操作用到的互斥量
	rmutex  sync.Mutex   // 读操作用到的互斥量
	dataLen uint32       // 数据块长度
	rcond   *sync.Cond   // 条件变量
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取一个数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	for {
		//df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				//df.fmutex.RUnlock()
				df.rcond.Wait()
				continue
			}
			//df.fmutex.RUnlock()
			return
		}
		d = bytes
		//df.fmutex.RUnlock()
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(len(d))
	df.wmutex.Unlock()

	// 写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	df.rcond.Signal()
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WRN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	return df.f.Close()
}

// 获取一个数据块文件的实例
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

var FileLength uint32 = 2048

func main() {
	df, err := NewDataFile("datafile.txt", FileLength)
	defer df.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("datafile ext: ", df)
	data := "test111222211"
	wsn, err := df.Write(Data(data))
	if err != nil {
		fmt.Println("write data error: ", err)
	}
	fmt.Println("wsn is: ", wsn, df.WRN())

	// test1.md read
	//rsn, d, err := df.Read()
	//if err != nil {
	//	fmt.Println("read file error: ", err)
	//}
	//fmt.Println("rsn: ", rsn)
	//fmt.Println("data: ", d)
}
