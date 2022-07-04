package main

import (
	"project/src/business"
	_ "project/src/config"
)

func init() {
	//log.InitLog("")
	//config.InitConfig("dev-f1")
	//redis.InitRedisClient()

}

func main() {
	business.Test()
	//db.TestOracle()
	//orm.TestOrm()
}
