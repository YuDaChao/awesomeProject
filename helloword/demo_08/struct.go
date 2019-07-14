/**
* @Author: YuDC
* @Date: 2019-07-09 22:24
* @Description: 结构体 方法
 */
package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	Addr string
}

// 给Student绑定一个read方法
func (student Student) read() {
	student.Name = "如来佛"
	fmt.Printf("read student: %v\n", student)
}

// 实现String方法
func (student Student) String() string {
	str := fmt.Sprintf("Name = %v, Age = %v, Addr = %v", student.Name, student.Age, student.Addr)
	return str
}

func main() {

	student := Student{Name: "猪八戒", Age: 400, Addr: "北京"}
	student.read()                         // read student: {如来佛 400 北京}
	fmt.Println("main student: ", student) // main student: {猪八戒 400 北京}

	// 如果实现了String方法，则打印时会自动调用String方法
	fmt.Println(student) // Name = 猪八戒, Age = 400, Addr = 北京

	stu := Student{"刘德华", 400, "香港"}

	fmt.Println(stu)
}
