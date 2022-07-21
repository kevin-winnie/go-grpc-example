package app

import (
	"fmt"
	"go-grpc/config"
	"go-grpc/pkg/mysql"
	"log"
)

type Order struct {
	order_no string
	pay_progress string
}

func Run(cfg *config.Config) {
	log.Println("项目开始运行")
	//初始化mysql
	dbObject := mysql.Init()
	mysql, err := dbObject.New("default", cfg.Mysql)
	if err != nil {
		log.Fatalf("db connect err:%s", err)
	}


	rows,err :=mysql.Query("select order_no,pay_progress from gops_order where order_no ='10104620210928151836273578'")
	if err != nil{
		log.Fatalf("db query err:%s", err)
	}

	for rows.Next(){	// 循环读取结果集中的数据
		var u Order
		err := rows.Scan(&u.order_no,&u.pay_progress)
		if err != nil{
			log.Fatalf("db scan err:%s", err)
		}
		fmt.Println(u)
	}
}
