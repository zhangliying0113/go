package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_learn"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return nil
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed,err:%v\n", err)
		return
	}

	sqlStr1 := `update users set age = age-2 where rec_id=6`
	sqlStr2 := `update users set age = age+2 where rec_id=7`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行 sql1 出错了，要回滚")
	}

	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("执行sql2出错了，要回滚！")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("事务提交失败，要回滚")
	}
	fmt.Println("事务提交成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功!")
	transactionDemo()
}
