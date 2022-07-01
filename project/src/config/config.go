package config

import (
	"fmt"
	"github.com/spf13/viper"
	. "project/src/log"
)

func init() {
	//InitConfig("")
}

func InitConfig(fileName string) {
	if fileName == "" {
		fileName = "application"
	}

	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("config/")         // path to look for the config file in
	viper.AddConfigPath("./src/config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	fmt.Println("=================== >>> app config is ", viper.AllSettings())
	er := viper.Unmarshal(&Config)
	if er != nil {
		panic(er)
	}
	Log.Infoln("=================== >>> context config is ", Config)
}

var Config AppConfig = AppConfig{}

type AppConfig struct {
	Redis redisProperties
}

type redisProperties struct {
	Network string
	// host:port address.
	Addr string
	// Optional password. Must match the password specified in the
	// requirepass server configuration option.
	Password string
	// Database to be selected after connecting to the server.
	DB int
	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int
}
