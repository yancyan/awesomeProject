package redis

import (
	"awesomeProject/src/config"
	. "awesomeProject/src/log"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var RedisDb *redis.Client

func InitRedisClient() {

	redisPros := config.Config.Redis

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     redisPros.Addr,     // redis地址
		Password: redisPros.Password, // redis密码，没有则留空
		DB:       redisPros.DB,       // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err := RedisDb.Ping().Result()
	if err != nil {
		panic(err)
	}

	Log.Println("the redis client init success.")
}

func TestRedis() {
	InitRedisClient()

	a1 := RedisDb.Get("test:a:a_1")
	fmt.Printf("before a1 is %s \n", a1.Val())

	RedisDb.Set("test:a:a_1", "abcValue"+time.Now().String(), 1*time.Minute)

	aa1 := RedisDb.Get("test:a:a_1")
	fmt.Printf("after a1 is %s \n", aa1.Val())

	time.Sleep(5 * time.Second)
}
