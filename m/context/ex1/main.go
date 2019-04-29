package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	mySleepAndTalk(ctx, 1*time.Second, "hello")
	cancel()
}

func mySleepAndTalk(ctx context.Context, duration time.Duration, msg string) {
	for {
		select {
		case <-time.After(duration):
			fmt.Println(msg)
		case <-ctx.Done():
			log.Print("err: ", ctx.Err())
			return
		}
	}
}
