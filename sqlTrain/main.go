package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {

	c := cron.New()
	c.AddFunc("@hourly", func() { fmt.Println("Every minute on the half hour") })
	c.Start()
	defer c.Stop()
	select {}
}
