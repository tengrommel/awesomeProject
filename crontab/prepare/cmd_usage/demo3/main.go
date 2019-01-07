package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

var (
	ctx        context.Context
	cancelFunc context.CancelFunc
	cmd        *exec.Cmd
	resultChan chan *result
	res        *result
)

type result struct {
	err    error
	output []byte
}

func main() {
	// 执行一个cmd，让它在一个协程去执行，让它执行2秒：sleep 2；echo hello;
	// 1秒的时候，我们杀死cmd
	ctx, cancelFunc := context.WithCancel(context.TODO())
	resultChan = make(chan *result, 1000)
	// context: chan byte
	// cancelFunc: close(chan byte)

	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/zsh", "-c", "sleep 2;echo hello;")
		// select(case <- ctx.Done():)
		// kill pid，进程ID，杀死子进程
		output, err = cmd.CombinedOutput()
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()
	// 继续往下写
	time.Sleep(3 * time.Second)
	cancelFunc()
	res = <-resultChan
	fmt.Println(res.err, string(res.output))
}
