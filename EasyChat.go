package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"./lib/model"
)

type Config struct {
	host string
	port string
	database string
	username string
	password string
}

func main () {
	// config := LoadConfig()
	model.Init()
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
		fmt.Println("key===")
		fmt.Printf("%s\n", strings.Replace(key, " ", "", -1))
		fmt.Println("value===")
		fmt.Printf("%s\n", strings.Replace(value, " ", "", -1))
	}
	return *c
}
