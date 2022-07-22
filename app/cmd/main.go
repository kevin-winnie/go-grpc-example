package main

import (
	"github.com/joho/godotenv"
	"go-grpc/config"
	"go-grpc/core"
	"log"
)

func init()  {
	//解析env文件
	_ = godotenv.Load()
}

func main()  {
	//加载配置文件
	cfg,err := config.NewConfig()
	if err != nil{
		log.Fatalf("Config error: %s", err)
	}
	//项目运行
	core.Run(cfg)
}
