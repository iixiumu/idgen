package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     128,
		MaxActive:   1024,
		IdleTimeout: 60 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", "127.0.0.1:6379") },
	}
}

type RedisIncr struct {
}

func (RedisIncr) GenID() uint64 {
	conn := RedisPool.Get()
	// fmt.Printf("%d %d\n", RedisPool.IdleCount(), RedisPool.ActiveCount())
	defer conn.Close()
	id, _ := redis.Uint64(conn.Do("INCR", "id"))
	return id
}
