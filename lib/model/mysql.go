package model

import (
	"fmt"
	// "database/sql"
	"os"
	"bufio"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	host string
	port string
	database string
	username string
	password string
}

func Init() {
	c := LoadConfig()
	fmt.Printf("*v", c)
	// db, err := sql.Open("mysql", c.username+":"+c.password+"@tcp("+c.host+":"+c.port+")/"+c.database+"?charset=utf8")
    // if err != nil {
    //     panic(err)
	// }
}

// 读取本地配置文件
func LoadConfig() Config {
	c := &Config{}
	
  	//创建一个结构体变量的反射
  	f,err := os.Open("./config/app.config")
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
		var value = strings.Replace(str[index + 1:], " ", "", -1)
		switch key {
			case "host":
				c.host = value
				break;
			case "port":
				c.port = value
				break;
			case "username":
				c.username = value
				break;
			case "password":
				c.password = value
				break;
			case "database":
				c.database = value
				break;
			
		}
	}
	return *c
}

func query(sql string) {
	rows,err := db.Query(sql);
    if err != nil{
        fmt.Printf("select fail [%s]",err)
    }

    var mapUser map[string]int
    mapUser = make(map[string]int)

    for rows.Next(){
        var id int
        var username string
        rows.Columns()
        err := rows.Scan(&id,&username)
        if err != nil{
            fmt.Printf("get user info error [%s]",err)
        }
        mapUser[username] = id
    }
}