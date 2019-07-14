/**
* @Author: YuDC
* @Date: 2019-07-11 00:55
* @Description:
 */
package model

// 定义一个私有的结构体
type student struct {
	name string
	age  int
}

// 创建一个student
func CreateStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func (student *student) GetName() string {
	return student.name
}

func (student *student) SetName(name string) {
	student.name = name
}

func (student *student) GetAge() int {
	return student.age
}
