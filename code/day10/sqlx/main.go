package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	RecID int `db:"rec_id"` // 如果不使用 tag，只能命名为 Rec_id 才能映射上，但是命名不规范
	Name  string
	Age   int
}

func initDB() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_learn"
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数，夜晚连接数较少时，空闲连接数较多时，就会自动关闭多余的连接数
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}

	fmt.Println("连接数据库成功")

	sqlStr := `select rec_id, name, age from users where rec_id = 1`
	var u user
	err = db.Get(&u, sqlStr)
	if err != nil {
		fmt.Println("select failed, err:", err)
	}
	// db.QueryRow(sqlStr).Scan(&u.RecID, &u.Name, &u.Age)
	fmt.Println("u:", u)

	var userList []user
	sqlStr1 := `select rec_id, name, age from users` // 使用 select * 映射不上，只能按字段查找
	err = db.Select(&userList, sqlStr1)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}
