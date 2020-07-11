package main

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

// 循环构建缩略图
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Fatalln(err)
		}
	}
}

// 使用goroutine循环并发构建，忽略错误
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // Ignoring errors
	}
}

// 使用channel通知外部goroutine完成情况
func makeThumbnail3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// 等待goroutine完成
	for {
		<-ch
	}
}

// 往 main goroutine 报告错误
func makeThumbnail4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // 如果有错误产生会造成阻塞
		}
	}
	return nil
}

// 使用buffered channel
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <- ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// 返回新文件的总大小
// 使用 waitGroup

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		// 在 worker goroutine之前调用
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			// 获取文件信息
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()	// 等价与 wg.Add(-1) 计数器减1
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func main() {

}
