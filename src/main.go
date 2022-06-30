package main

import (
	_ "awesomeProject/config"
	"awesomeProject/redis"
)

func main() {
	redis.TestRedis()

}
