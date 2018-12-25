package redisTool

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var (
	redis_host     string
	redis_port     string
	redis_password string
	redisClient    *redis.Pool
)

func RedisLink() {
	redis_host = beego.AppConfig.String("redis.host")
	redis_port = beego.AppConfig.String("redis.port")
	redis_password = beego.AppConfig.String("redis.password")
	redisClient = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1024,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_host+":"+redis_port)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			// 选择db
			//c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

//获取redisClient链接
func GetRedisClient() *redis.Conn {
	redisItem := redisClient.Get()
	return &redisItem
}

//GET字符串
func GetStringKey(key string) string {
	redisTem := redisClient.Get()
	str, err := redis.String(redisTem.Do("GET", key))
	redisTem.Close()
	if err != nil {
		return str
	}
	return str
}

//set字符串
func setStringKey(key string, value string) bool {
	redisTem := redisClient.Get()
	n, err := redisTem.Do("SET", key, value)
	redisTem.Close()
	if err != nil {
		fmt.Println(n)
		return false
	}
	return true
}

//队列前插入
func ListEdit(key string, value string) bool {
	redisTem := redisClient.Get()
	redisTem.Do("LPUSH", key, value)
	redisTem.Close()
	return false
}

//查询队列
func RangeList(key string) {}

//队列情况
func ClearList(key string) {}

//是不是存在这个队列
func IsExeistList(key string) {}
