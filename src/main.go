package main

import (
	"awesomeProject/src/config"
	_ "awesomeProject/src/config"
	"awesomeProject/src/log"
	"awesomeProject/src/redis"
)

func init() {
	log.InitLog("")
	config.InitConfig("dev-f1")
	// ===============
	redis.InitRedisClient()
}

func main() {

}
