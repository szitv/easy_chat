package model

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

type Config struct {
	host     string
	port     string
	database string
	username string
	password string
}

type Home struct {
	id        int    `db:"id"`
	address   string `db:"address"`
	room      string `db:"room"`
	startTime string `db:"start_time"`
}

func Init() {
	c := LoadConfig()
	Db, err := sql.Open("mysql", c.username+":"+c.password+"@("+c.host+":"+c.port+")/"+c.database+"?charset=utf8")

	// 最大连接数
	Db.SetMaxOpenConns(100)
	// 闲置连接数
	Db.SetMaxIdleConns(20)
	// 最大连接周期
	Db.SetConnMaxLifetime(100 * time.Second)

	if err = Db.Ping(); nil != err {
		panic("数据库链接失败: " + err.Error())
	}

	if err != nil {
		fmt.Printf("%v", err)
	}

	defer Db.Close()
	homes := make([]Home, 0)
	rows, err := Db.Query("SELECT * FROM home limit 0, 10")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", rows)
	var home Home

	for rows.Next() {
		rows.Scan(&home.id, &home.address, &home.room, &home.startTime)
		homes = append(homes, home)
		// fmt.Println(id)
		// fmt.Println(address)
		// fmt.Println(room)
		// fmt.Println(start_time)
	}
	// fmt.Println(homes)
}

// 读取本地配置文件
func LoadConfig() Config {
	c := &Config{}

	//创建一个结构体变量的反射
	f, err := os.Open("./config/app.config")
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("%v", err)
		}
	}()
	//我们要逐行读取文件内容
	s := bufio.NewScanner(f)
	for s.Scan() {
		//以=分割,前面为key,后面为value
		var str = s.Text()
		var index = strings.Index(str, "=")
		var key = strings.Replace(str[0:index], " ", "", -1)
		var value = strings.Replace(str[index+1:], " ", "", -1)
		switch key {
		case "host":
			c.host = value
			break
		case "port":
			c.port = value
			break
		case "username":
			c.username = value
			break
		case "password":
			c.password = value
			break
		case "database":
			c.database = value
			break

		}
	}
	return *c
}
