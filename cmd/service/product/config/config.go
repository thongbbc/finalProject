package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Username   string
	Password   string
	DbName     string
	DbHost     string
	Port       string
	HostServer string
	PortRedis  string
	RedisHost  string
}

func SetupConfig() *Config {
	e := godotenv.Load("./cmd/service/.env") //Load .env file
	if e != nil {
		fmt.Print(e)
		panic(e)
	}

	return &Config{
		Username:   os.Getenv("db_user"),
		Password:   os.Getenv("db_pass"),
		DbName:     os.Getenv("db_name"),
		DbHost:     os.Getenv("db_host"),
		Port:       os.Getenv("port"),
		HostServer: os.Getenv("host"),
		PortRedis:  os.Getenv("portRedis"),
		RedisHost:  os.Getenv("redis_host"),
	}
}