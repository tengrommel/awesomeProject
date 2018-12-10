package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"job_name"`  // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点信息
}

// jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"` // JobName 赋值为10
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		cond       *FindByJobName
		cursor     mongo.Cursor
		record     *LogRecord
	)

	clientOptions := options.ClientOptions{}
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientOptions.SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}
	database = client.Database("cron")
	collection = database.Collection("log")

	cond = &FindByJobName{JobName: "job10"}
	findopt := options.FindOptions{}
	if cursor, err = collection.Find(context.TODO(), cond, findopt.SetSkip(int64(0)), findopt.SetLimit(int64(2))); err != nil {
		fmt.Println(err)
		return
	}
	for cursor.Next(context.TODO()) {
		record = &LogRecord{}
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*record)
	}
}
