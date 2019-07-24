/**
* @Author: YuDC
* @Date: 2019-07-21 19:12
* @Description:
 */
package main

import (
	"awesomeProject/helloword/demo_13/model"
	"fmt"
)

func main() {

	var key string
	loop := true
	var (
		id      int
		age     int
		name    string
		gender  string
		address string
	)
	stu := model.NewStudent(id, name, age, gender, address)
	for {
		fmt.Println("1：添加学生")
		fmt.Println("2：修改学生")
		fmt.Println("3：删除学生")
		fmt.Println("4：学生列表")
		fmt.Println("5：退   出")
		fmt.Println("请选择(1-5):")
		_, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Println("请输入1-5")
		}
		switch key {
		case "1":
			fmt.Println("请一次输入: id name age gender address")
			fmt.Scanln(&id)
			fmt.Scanln(&name)
			fmt.Scanln(&age)
			fmt.Scanln(&gender)
			fmt.Scanln(&address)
			stu.SetId(id)
			stu.SetName(name)
			stu.SetAge(age)
			stu.SetGender(gender)
			stu.SetAddress(address)
			stu.AddStudent(*stu)
		case "2":
			fmt.Println("请输入id")
			var id int
			fmt.Scanln(&id)
			stuList := stu.GetStudentList()
			student := model.FindStudentById(stuList, id)
			student.SetName("唐僧")
		case "4":
			stuList := stu.GetStudentList()
			fmt.Println("编号\t\t\t\t姓名\t\t\t\t年龄\t\t\t\t性别\t\t\t\t住址")
			for _, v := range stuList {
				fmt.Println(v)
			}
		default:
			loop = false
		}
		if !loop {
			break
		}
	}

}
