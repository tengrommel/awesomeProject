package worker

import (
	"awesomeProject/crontab/common"
	"fmt"
	"time"
)

type Scheduler struct {
	jobEventChan      chan *common.JobEvent              // etcd 任务事件队列
	jobPlanTable      map[string]*common.JobSchedulePlan // 任务调度计划表
	jobExecutingTable map[string]*common.JobExecuteInfo
	jobResultChan     chan *common.JobExecuteResult
}

var (
	G_scheduler *Scheduler
)

// 处理任务事件
func (scheduler *Scheduler) handleJobEvent(jobEvent *common.JobEvent) {
	var (
		jobSchedulePlan *common.JobSchedulePlan
		jobExisted      bool
		err             error
	)
	switch jobEvent.EventType {
	case common.JOB_EVENT_SAVE:
		if jobSchedulePlan, err = common.BuildJobSchedulePlan(jobEvent.Job); err != nil {
			return
		}
		scheduler.jobPlanTable[jobEvent.Job.Name] = jobSchedulePlan
	case common.JOB_EVENT_DELETE:
		if jobSchedulePlan, jobExisted = scheduler.jobPlanTable[jobEvent.Job.Name]; jobExisted {
			delete(scheduler.jobPlanTable, jobEvent.Job.Name)
		}
	}
}

func (scheduler *Scheduler) handleJobResult(result *common.JobExecuteResult) {
	delete(scheduler.jobExecutingTable, result.ExecuteInfo.Job.Name)
}

// 调度协程
func (scheduler *Scheduler) scheduleLoop() {
	var (
		jobEvent      *common.JobEvent
		scheduleAfter time.Duration
		scheduleTimer *time.Timer
		jobReuslt     *common.JobExecuteResult
	)
	// 定时任务
	scheduleAfter = scheduler.TrySchedule()
	scheduleTimer = time.NewTimer(scheduleAfter)

	for {
		select {
		case jobEvent = <-scheduler.jobEventChan: // 监听任务变化事件
			// 对内存中维护的任务列表做增删改查
			scheduler.handleJobEvent(jobEvent)
		case <-scheduleTimer.C:
		case jobReuslt = <-scheduler.jobResultChan:
			scheduler.handleJobResult(jobReuslt)
		}
		scheduleAfter = scheduler.TrySchedule()
		scheduleTimer.Reset(scheduleAfter)
	}
}

func (scheduler *Scheduler) TryStartJob(jobSchedulePlan *common.JobSchedulePlan) {
	// 调度和执行是两个事件
	// 执行可能运行很久
	var (
		jobExecuteInfo *common.JobExecuteInfo
		jobExecuting   bool
	)
	if jobExecuteInfo, jobExecuting = scheduler.jobExecutingTable[jobSchedulePlan.Job.Name]; jobExecuting {
		fmt.Println("尚未跳过执行")
		return
	}
	jobExecuteInfo = common.BuildJobExecuteInfo(jobSchedulePlan)
	scheduler.jobExecutingTable[jobSchedulePlan.Job.Name] = jobExecuteInfo
	G_executor.ExecuteJob(jobExecuteInfo)

}

// 重新计算任务调度状态
func (scheduler *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var (
		jobPlan  *common.JobSchedulePlan
		now      time.Time
		nearTime *time.Time
	)
	if len(scheduler.jobPlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}
	now = time.Now()
	for _, jobPlan = range scheduler.jobPlanTable {
		if jobPlan.NextTime.Before(now) {
			scheduler.TryStartJob(jobPlan)
			jobPlan.NextTime = jobPlan.Expr.Next(now)
		}

		// 统计最近一个要过期的任务事件
		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}
	scheduleAfter = (*nearTime).Sub(now)
	return
}

// 推送任务变化事件
func (scheduler *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	scheduler.jobEventChan <- jobEvent
}

// 初始化调度器
func InitScheduler() (err error) {
	G_scheduler = &Scheduler{
		jobEventChan:      make(chan *common.JobEvent, 1000),
		jobPlanTable:      make(map[string]*common.JobSchedulePlan),
		jobExecutingTable: make(map[string]*common.JobExecuteInfo),
		jobResultChan:     make(chan *common.JobExecuteResult, 1000),
	}
	// 启动调度协程
	go G_scheduler.scheduleLoop()
	return
}

func (scheduler *Scheduler) PushJobResult(jobResult *common.JobExecuteResult) {
	scheduler.jobResultChan <- jobResult
}
