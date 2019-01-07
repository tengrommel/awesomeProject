package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		result     *mongo.InsertOneResult
		docId      primitive.ObjectID
	)

	clientOptions := options.ClientOptions{}
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientOptions.SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}
	database = client.Database("cron")
	collection = database.Collection("log-test")
	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}
	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}
	// id:默认生成一个全局唯一ID,ObjectID: 12字节的二进制
	docId = result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())
}
