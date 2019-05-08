package main

import (
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

const REDIS_LOCK = "redis_lock"

type MysqlRedisLock struct {
}

func (MysqlRedisLock) GenID() uint64 {
	conn := RedisPool.Get()
	d := time.Millisecond
	for i := 0; ; i++ {
		ret, _ := redis.String(conn.Do("SET", REDIS_LOCK, REDIS_LOCK, "EX", "1", "NX"))
		if ret == "OK" {
			break
		} else {
			max := 16
			if i < 4 {
				max = 1 << uint(i)
			}
			t := rand.Intn(max)
			time.Sleep(d * time.Duration(t))
		}
	}

	defer func() {
		conn.Do("DEL", REDIS_LOCK)
		conn.Close()
	}()

	ins := &MyID{}
	err := MyDB.Where("id = ?", 4).First(ins).Error
	if err != nil {

	}
	ins.MyID += 1
	MyDB.Save(ins)
	return ins.MyID
}
