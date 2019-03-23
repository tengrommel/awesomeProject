package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*
考场签到，名字丢入管道；
只有5个车道，最多5个人同时考试；
考生按签到顺序依次考试，给予考生10%的违规几率；
每3秒巡视一次，发现违规的清出考场时序良好
所有考试者考完后，生成考试记录
当前目录下的成绩录入MySQL数据库，数据库允许一写多读
成绩录入完毕通知考生，考生查阅自己的成绩
*/

var (
	chNames = make(chan string, 100)
	examers = make([]string, 0)
	// 信号量
	chLanes = make(chan int, 5)
	// 违纪者
	chFouls = make(chan string, 100)
	// 考试成绩
	scoreMap = make(map[string]int)
	wg       sync.WaitGroup
	// 随机数互斥锁（确保GetRandomInt不能被并发访问）
	randomMutex sync.Mutex
	//数据库读写锁
	dbMutext sync.RWMutex
	// 姓氏
	familyNames = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "楚", "卫", "蒋", "沈", "韩", "杨", "张", "欧阳", "东门", "西门", "上官", "夏侯"}
	// 辈分（宗的永其光...）
	middleNamesMap = map[string][]string{}
	// 名字
	lastNames = []string{"春", "夏", "秋", "冬", "风", "霜", "雨", "雪", "木", "禾", "米", "竹", "山", "石", "田", "土", "福", "禄", "寿", "喜", "文", "武", "才", "华"}
)

func init() {
	for _, x := range familyNames {
		if x != "欧阳" {
			middleNamesMap[x] = []string{"德", "惟", "守", "世", "令", "子", "伯", "师", "希", "与", "孟", "由", "宜", "顺", "元", "允", "宗", "仲", "士", "不", "善", "汝", "崇", "必", "良", "友", "季", "同"}
		} else {
			middleNamesMap[x] = []string{"宗", "的", "永", "其", "光"}
		}
	}
}

func main() {
	for i := 0; i < 100; i++ {
		chNames <- GetRandomName()
	}
	close(chNames)
	/*
		巡考
		考生并发考试
	*/
	go Patrol()
	for name := range chNames {
		wg.Add(1)
		go TakeExam(name)
	}
	fmt.Println("考试完毕")
	wg.Wait()
	wg.Add(1)
	go WriteScore2DB()
	// 故意给一个时间间隔，确保WriteScore2DB先抢到数据库的读写锁
	<-time.After(1 * time.Second)
	// 考生查询成绩
	for _, name := range examers {
		wg.Add(1)
		go QueryScore(name)
	}
	wg.Wait()
	fmt.Println("考生查询自己的成绩")
}

type ExamScore struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Score int    `db:"score"`
}

/*
查询成绩
*/
func QueryScore(name string) {
	// 吸入期间不允许数据库的读访问
	dbMutext.RLock()
	db, err := sqlx.Connect("mysql", "root@teng@tcp(localhost:3306)/driving_exam")
	HandlerError(err, `sqlx.Connect("mysql", "root@teng@tcp(localhost:3306)/driving_exam")`)
	examScore := make([]ExamScore, 0)
	err = db.Select(&examScore, "select * from score where name=?;", name)
	HandlerError(err, `db.Select(&examScore, "select * from score where name=?;", name)`)
	dbMutext.RUnlock()
	fmt.Println(examScore)
	wg.Done()
}

func Patrol() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		fmt.Println("战狼正在巡考")
		select {
		case name := <-chFouls:
			fmt.Println(name, "考试违纪！！！！")
		default:
			fmt.Println("考场秩序良好")
		}
		<-ticker.C
	}
}

func TakeExam(name string) {
	chLanes <- 123
	fmt.Println(name, "正在考试...")
	// 将参与考试的考生姓名
	examers = append(examers, name)
	score := GetRandomInt(0, 100)
	scoreMap[name] = score
	if score < 10 {
		chFouls <- name
		fmt.Println(name, "考试违纪!!!")
	}
	<-time.After(5 * time.Second)
	<-chLanes
	wg.Done()
}

func GetRandomName() (name string) {
	familyName := familyNames[GetRandomInt(0, len(familyNames)-1)]
	middleName := middleNamesMap[familyName][GetRandomInt(0, len(middleNamesMap[familyName])-1)]
	lastName := lastNames[GetRandomInt(0, len(lastNames)-1)]
	return familyName + middleName + lastName
}

func GetRandomInt(start, end int) int {
	randomMutex.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start+1)
	randomMutex.Unlock()
	return n
}

func HandlerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func WriteScore2DB() {
	// 锁定为写模式，写入期间不允许读访问
	dbMutext.Lock()
	db, err := sqlx.Connect("mysql", "root:teng@tcp(localhost:3306)/driving_exam")
	HandlerError(err, `sqlx.Connect("mysql", "root:teng@tcp(localhost:3306)/driving_exam")`)
	for name, score := range scoreMap {
		_, err := db.Exec("insert into score(name, score) values (?, ?);", name, score)
		HandlerError(err, `db.Exec("insert into score(name, score) values (?, ?);", name, score)`)
		fmt.Println(name, "插入成功!")
	}
	fmt.Println("成绩录入完成")
	dbMutext.Unlock()
	wg.Done()
}
