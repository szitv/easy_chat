package model

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

 
var (
    userName  string = "root"
    password  string = "test123456"
    ipAddrees string = "localhost"
    port      int    = 3306
    dbName    string = "test"
    charset   string = "utf8"
)
 
func ConnectMysql() (*sqlx.DB) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
    Db, err := sqlx.Open("mysql", dsn)
    if err != nil {
        fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
    }
    return Db
}