package core

import (
	"fire_eye/pkg/logger"
	"fmt"
	"go-grpc/config"
	"go-grpc/controller/rpc/service"
	"go-grpc/pkg/mysql"
	"google.golang.org/grpc"
	"log"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	//初始化mysql
	dbObject := mysql.Init()
	db, err := dbObject.New("default", cfg.Mysql)
	if err != nil {
		log.Fatalf("db connect err:%s", err)
	}

	//rs, err := model.QueryOrder(db)
	//if err != nil {
	//	log.Fatalf("db query err:%s", err)
	//}
	//fmt.Println(rs)

	//grpc注册
	grpcServer := grpc.NewServer()
	service.Register(grpcServer)

	//crontab注册

	//http服务 gin框架
	// HTTP
	//handler := gin.New()

	// Mysql Close
	if err := db.Close(); err != nil {
		l.Error(fmt.Errorf("app - Run - MysqlMap.Close: %w", err))
	}

}
