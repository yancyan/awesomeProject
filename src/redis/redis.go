package redis

import (
	"awesomeProject/config"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 声明一个全局的redisDb变量
var RedisDb *redis.Client

// 根据redis配置初始化一个客户端
func InitRedisClient() (err error) {

	redisPros := config.Config.Redis

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     redisPros.Addr,     // redis地址
		Password: redisPros.Password, // redis密码，没有则留空
		DB:       redisPros.DB,       // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = RedisDb.Ping().Result()
	if err != nil {
		return err
	}

	fmt.Println("the redis client init success.")
	return nil
}

func TestRedis() {
	err := InitRedisClient()
	if err != nil {
		panic(err)
	}

	a1 := RedisDb.Get("test:a:a_1")
	fmt.Printf("before a1 is %s \n", a1.Val())

	RedisDb.Set("test:a:a_1", "abcValue"+time.Now().String(), 1*time.Minute)

	aa1 := RedisDb.Get("test:a:a_1")
	fmt.Printf("after a1 is %s \n", aa1.Val())

	time.Sleep(5 * time.Second)
}
