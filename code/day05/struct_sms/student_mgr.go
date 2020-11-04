package main

import "fmt"

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudents map[int64]student
}

// 1. 查看所有学生信息
func (s studentMgr) showStudents() {
	for _, student := range s.allStudents {
		fmt.Printf("学生id:%d 学生姓名：%s\n", student.id, student.name)
	}
}

// 2. 增加学生
func (s studentMgr) addStudents() {
	var (
		stuID   int64
		stuName string
	)

	fmt.Println("请输入学生学号：")
	fmt.Scanln(&stuID)

	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&stuName)

	newStudent := student{
		id:   stuID,
		name: stuName,
	}

	// s.allStudents[stuID] = newStudent
	s.allStudents[newStudent.id] = newStudent
	fmt.Println("添加成功")
}

// 3. 修改学生信息
func (s studentMgr) updateStudents() {
	var (
		stuID   int64
		newName string
	)

	fmt.Println("请输入学生学号：")
	fmt.Scanln(&stuID)

	// 判断有无此学生
	stuObj, ok := s.allStudents[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你想要修改的学生信息如下：学号：%d，姓名：%s\n", stuObj.id, stuObj.name)

	fmt.Println("请输入学生的新名字：")
	fmt.Scanln(&newName)
	stuObj.name = newName

	s.allStudents[stuID] = stuObj
}

// 4. 删除学生
func (s studentMgr) deleteStudents() {

	var stuID int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&stuID)
	_, ok := s.allStudents[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudents, stuID)
	fmt.Println("删除成功")
}
