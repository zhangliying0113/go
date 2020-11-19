package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // 定义一个数据库连接池对象

type user struct {
	recID int
	name  string
	age   int
}

func initDB() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_learn"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确,只有 dsn 格式不正确才会报错
	if err != nil {                  // dsn格式不正确的时候会报错
		return
	}
	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数，夜晚连接数较少时，空闲连接数较多时，就会自动关闭多余的连接数
	return
}

// 插入数据
func insert() {
	sqlStr := `insert into users(name, age) values("小白", 1),("xiaobai", 2)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	// 如果是插入数据操作，能够拿到插入数据的 id，插入多条时只能拿到插入的最小 id LastInsertId
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("rec_id:", id)
}

// 查询单个记录
func queryOne(recID int) {
	var u user
	sqlStr := "select rec_id, name, age from users where rec_id = ?" // 必须使用字段查询，使用 * 在 Scan 处匹配不上字段
	db.QueryRow(sqlStr, recID).Scan(&u.recID, &u.name, &u.age)
	fmt.Println("u:", u)
}

// 查询多个记录
func queryMore(recID int) {
	sqlStr := `select rec_id, name, age from users where rec_id > ?`
	rows, err := db.Query(sqlStr, recID)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}

	// 一定要关闭 rows
	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.recID, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("u1:%#v\n", u)
	}
}

// 更新操作
func updateRow(newAge, recID int) {
	sqlStr := `update users set age = ? where rec_id > ?`
	ret, err := db.Exec(sqlStr, newAge, recID)
	if err != nil {
		fmt.Println("update users failed, err:", err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get affected num failed, err:", err)
	}
	fmt.Printf("更新了%d行", n)
}

// 删除数据
func deleteRow(recID int) {
	sqlStr := `delete from users where rec_id = ?`
	ret, err := db.Exec(sqlStr, recID)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get affected num failed, err:", err)
	}
	fmt.Printf("删除了%d行", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into users(name,age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("insert prepare failed, err:", err)
	}
	defer stmt.Close()
	var m = map[string]int{
		"小白1": 2,
		"小白2": 3,
		"小白3": 4,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功!")
	// insert()
	// queryOne(1)
	// queryMore(2)
	// updateRow(3, 1)
	// deleteRow(4)
	prepareInsert()
}
