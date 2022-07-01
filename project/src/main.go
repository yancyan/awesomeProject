package main

import (
	_ "project/src/config"
	"project/src/orm"
)

func init() {
	//log.InitLog("")
	//config.InitConfig("dev-f1")
	//redis.InitRedisClient()

}

func main() {
	//business.Test()
	//db.TestOracle()
	orm.TestOrm()
}
