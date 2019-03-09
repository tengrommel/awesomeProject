package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type Human struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func HandleErr(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
		os.Exit(1)
	}
}

func main() {
	var cmd string
	for {
		fmt.Println("请输入命令：")
		fmt.Scan(&cmd)
		switch cmd {
		case "getall":
			// 显示所有人员信息
			GetAllPeople()
		case "exit":
			goto GAMEOVER
		default:
			fmt.Println("什么破命令，fuck off!")
		}
	}
GAMEOVER:
	fmt.Println("GAME OVER")
}

func GetAllPeople() {
	fmt.Println("GetAllPeople")

	peopleList := GetPeopleFromRedis()
	if peopleList == nil || len(peopleList) == 0 {
		db, _ := sqlx.Connect("mysql", "root:teng@tcp(localhost:3306)/china")
		defer db.Close()
		var people []Human
		err := db.Select(&people, "select name, age from person")
		if err != nil {
			fmt.Println("查询失败，err=", err)
		}
		fmt.Println(people)
		// 缓存查询结果到Redis
		CachePeople2Redis(people)
	} else {
		fmt.Println(peopleList)
	}
}

func GetPeopleFromRedis() (peopleStrs []string) {
	conn, _ := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	reply, err := conn.Do("lrange", "people", "0", "-1")
	HandleErr(err, "@lrange people 0 -1")
	peopleStrs, err = redis.Strings(reply, err)
	return
}

// 缓存查询结果到redis
func CachePeople2Redis(people []Human) {
	conn, _ := redis.Dial("tcp", "localhost:6379")
	defer conn.Close()
	for _, human := range people {
		humanString := fmt.Sprint(human)
		_, err := conn.Do("rpush", "people", humanString)
		HandleErr(err, "rpush people "+humanString)
	}
	_, err := conn.Do("expire", "people", 60)
	HandleErr(err, "@ expire people 60")
	fmt.Println("缓存成功")
}
