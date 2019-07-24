/**
* @Author: YuDC
* @Date: 2019-07-21 19:12
* @Description:
 */
package model

import "fmt"

type student struct {
	id      int
	name    string
	age     int
	gender  string
	address string
	list    []student
}

func (student *student) SetId(id int) {
	student.id = id
}

func (student *student) SetName(name string) {
	student.name = name
}

func (student *student) SetAge(age int) {
	student.age = age
}

func (student *student) SetGender(gender string) {
	student.gender = gender
}

func (student *student) SetAddress(address string) {
	student.address = address
}

// 获取学生列表
func (student *student) GetStudentList() []student {
	return student.list
}

// 添加一个学生信息
func (student *student) AddStudent(stu student) {
	student.list = append(student.list, stu)
}

// 创建一个学生对象
func NewStudent(id int, name string, age int, gender string, address string) *student {
	return &student{
		id:      id,
		name:    name,
		gender:  gender,
		age:     age,
		address: address,
	}
}

// 根据id查找学生信息
func FindStudentById(students []student, id int) *student {
	var s student
	for _, v := range students {
		if v.id == id {
			s = v
		}
	}
	return &s
}

func (student student) String() string {
	str := fmt.Sprintf("%d\t\t\t\t%s\t\t\t\t%d\t\t\t\t%s\t\t\t\t%s",
		student.id, student.name, student.age, student.gender, student.address)
	return str
}
