package app

import (
	"fmt"
	"go-grpc/config"
)

func Run(cfg *config.Config)  {
	fmt.Println("开始进入项目")
	fmt.Println(cfg)
}
