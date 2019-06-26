package main

import "fmt"

// 定义结构体

type student struct {
	id      int
	name    string
	age     int
	gender  byte
	address string
}

func main() {
	// 如果不指定变量名， 必须全部成员变量都初始化
	var s = student{1, "悟空", 23, 1, "北京"}
	fmt.Println(s)

	var s2 = student{name: "猪八戒"}

	s3 := student{id: 1, name: "唐僧"}
	fmt.Println(s2, s3)

	// 配合指针

	var stu = &student{id: 1, name: "悟空"}

	//fmt.Println("stu = ", stu)
	fmt.Println("*stu : ", *stu)

	// 成员操作

	// 获取成员

	var name = stu.name
	fmt.Println("name: ", name)

	// 赋值
	stu.address = "上海"
	fmt.Println("address: ", stu.address)

	// 比较

	//stu1 := student{id: 1, name: "ydc", gender: 1, address: "bj"}
	//stu2 := student{id: 1, name: "ydc", gender: 1, address: "bj"}
	//fmt.Println(stu1 == stu2) // true

	stu1 := new(student)
	stu2 := new(student)

	stu1.id = 1
	stu2.id = 1

	fmt.Printf("%T \n", stu1)
	fmt.Println(stu1 == stu2) // false
}
