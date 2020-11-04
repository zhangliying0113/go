package main

import (
	"fmt"
	"os"
)

// 学生管理系统
var sms studentMgr // 声明一个全局的变量 学生管理系统对象

func showMenu() {
	fmt.Println("-------------Welcome sms!-------------")
	fmt.Println(`
	1. 查看所有学成信息
	2. 添加学生信息
	3. 修改学生信息
	4. 删除学生信息
	5. 退出系统
	`)
}

func main() {
	// 1. 学生管理系统首页显示
	sms = studentMgr{
		allStudents: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 2. 提示用户输入操作选项
		fmt.Print("请输入操作序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("您输入的是：", choice)

		switch choice {
		case 1:
			sms.showStudents()
		case 2:
			sms.addStudents()
		case 3:
			sms.updateStudents()
		case 4:
			sms.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("服务暂未提供")
		}
	}
}
