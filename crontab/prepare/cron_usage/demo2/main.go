package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	// 需要有1个调度协程，它定时检查所有的Cron任务，谁过期了就执行谁
	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob
	)
	scheduleTable = make(map[string]*CronJob)

	now = time.Now()
	// 1、我们定义两个cronjob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	// 任务注册到调度表
	scheduleTable["job2"] = cronJob
	// 启动一个调度协程
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				// 判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					// 启动一个协程，执行这个任务
					go func(jobName string) {
						fmt.Println("执行：", jobName)
					}(jobName)
					// 计算下一次调度时间
					fmt.Println(jobName, "下次执行时间：", cronJob.nextTime)
				}
			}
			// 睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Microsecond).C:

			}
		}
	}()
	time.Sleep(20 * time.Second)
}
