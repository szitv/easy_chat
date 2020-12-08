package model

import (
	"fmt"
	"../lib/model"
	"strconv"
	"strings"
)

type User struct {
	id int
	address string
}

var Db = model.ConnectMysql()

const table string = "user"


func find(sql string) {
	rows, err := Db.Query(sql)
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var user User

		err := rows.Scan(&user.id, &user.address)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println(user.id, user.address)
	}

	//关闭结果集（释放连接）
	rows.Close()
}

func UserList(start int, size int) {
	defer Db.Close()
	
	columns := []string{"id", "name"}
	fields := strings.Join(columns, ",")
	sql := fmt.Sprintf("select %s from %s", fields, table)
	sql += " where id > 0 limit "+strconv.Itoa(start)+", "+strconv.Itoa(size)
	find(sql)
}