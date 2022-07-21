package main

import (
	"github.com/joho/godotenv"
	"go-grpc/config"
	"go-grpc/internal/app"
	_ "go-grpc/internal/app"
	"log"
)

func init()  {
	_ = godotenv.Load()
}

func main()  {
	cfg,err := config.NewConfig()
	if err != nil{
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
