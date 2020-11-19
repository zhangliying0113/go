package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入包但是不直接使用包，用 _
)

func main() {
	// 连接数据库
	// 用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_learn"

	// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                   // dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}

	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功！")

}
