package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initMySQL() (err error) {
	// DSN: Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 去初始化全局的db对象而不是声明一个新的db变量
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 做完错误检查之后，确保db不为nil
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed, err: %v\n", err)
		return
	}
	// 数值需要业务具体情况来确定
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed, err: %v\n", err)
	}
	// Close() 用来释放掉数据库连接相关的资源
	defer db.Close() // 注意这行代码要写在上面err判断的下面
	fmt.Println("connect to db success")
}
