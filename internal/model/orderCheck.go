package model

import (
	"database/sql"
	"log"
)

type Order struct {
	order_no     string
	pay_progress string
}

func QueryOrder(db *sql.DB) (data []Order, err error) {

	rows, err := db.Query("select order_no,pay_progress from gops_order where project_code ='101037'")
	if err != nil {
		return nil, err
	}

	for rows.Next() { // 循环读取结果集中的数据
		var u Order
		err := rows.Scan(&u.order_no, &u.pay_progress)
		if err != nil {
			log.Fatalf("db scan err:%s", err)
		}
		data = append(data, u)
	}
	return data, nil
}
