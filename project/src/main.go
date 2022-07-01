package main

import (
	"fmt"
	"github.com/yancyan/merci-tools"
	utils "github.com/yancyan/tools"
	"project/src/config"
	_ "project/src/config"
	"project/src/log"
	"project/src/redis"
)

func init() {
	log.InitLog("")
	config.InitConfig("dev-f1")
	// ===============
	redis.InitRedisClient()

}

func main() {
	ab := merci_tools.StringMd5("abc")
	fmt.Println(ab)

	utils.Print("abc")

}
