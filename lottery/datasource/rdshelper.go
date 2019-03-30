package datasource

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

var rdsLock sync.Mutex
var cacheInstance *redis.Conn

type RedisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (rds *RedisConn) Do(commandName string,
	args ...interface{}) (reply interface{}, err error) {
	conn := rds.pool.Get()
	defer conn.Close()
	t1 := time.Now().UnixNano()
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println("rbshelper.Do", err, e)
		}
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		fmt.Printf("[redis] [info] [%dus]cmd=%s, err=%s, args=%v, reply=%s\n", (t2-t1)/1000, commandName, err, args, reply)
	}
	return reply, err
}
