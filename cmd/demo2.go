package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resChan    chan *result
	)

	// 执行一个cmd,让它在一个协程中执行，执行两秒；sleep 2; echo hello;
	// 1s时，kill cmd

	ctx, cancelFunc = context.WithCancel(context.TODO())
	resChan = make(chan *result, 100)

	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "C:\\Users\\84576\\.babun\\cygwin\\bin\\bash.exe", "-c", "sleep 2; echo hello;")

		output, err = cmd.CombinedOutput()
		// 把结果传递给 main goroutine
		resChan <- &result{
			output: output,
			err: err,
		}
	}()

	time.Sleep(1 * time.Second)

	// 取消上下文
	cancelFunc()

	res := <-resChan
	fmt.Println(res.err, string(res.output))

	// 开源库 Cronexpr
}
