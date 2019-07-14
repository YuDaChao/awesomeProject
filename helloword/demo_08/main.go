/**
* @Author: YuDC
* @Date: 2019-07-06 19:33
* @Description: struct
 */
package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体

type Person struct {
	name string
	age  int
	addr string
}

type Monster struct {
	Name  string `json:"name"` // 结构体 tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	// 结构体是值类型
	var person Person

	fmt.Println(person) // { 0 }

	// 初始化 并赋值
	person = Person{"孙悟空", 23, "北京"}

	person1 := person

	person1.name = "猪八戒"

	fmt.Println(person)  // {孙悟空 23 北京}
	fmt.Println(person1) // {猪八戒 23 北京}

	// 获取结构体地址

	ptr := &Person{}
	// 如果需要初始化的时候需要赋值 那么必须初始化所有字段 否则会编译不通过
	//ptr := &Person{"jack", 12, "北京"}

	fmt.Println("ptr = ", ptr)

	// 使用new创建结构体 此时的类型为指针类型

	var p = new(Person)

	// 赋值
	(*p).name = "沙沙"

	// 上面等价于
	p.name = "沙沙"

	fmt.Println("p = ", p)   // &{"沙沙" 0 ""}
	fmt.Println("*p = ", *p) // {"沙沙" 0 ""}

	fmt.Printf("p 的类型是: %v, 地址是: %p", p, &p) //  &{"沙沙" 0 } 0xc00000e028

	// 结构体json序列化

	monster := Monster{"孙悟空", 400, "七十二变"}

	fmt.Println(monster)

	// 序列化

	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("序列化错误：", err)
	}

	fmt.Println("jsonStr: ", string(jsonStr))

}
