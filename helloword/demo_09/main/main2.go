/**
* @Author: YuDC
* @Date: 2019-07-14 15:32
* @Description:
 */
package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A sayOK", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.age)
}

type B struct {
	A
	Name string // 字段名可以和A中的重复 访问时会优先在B中查找 找不到时才会到A中查找
}

func main() {
	var b B

	b.A.Name = "Jack"
	b.A.age = 12

	b.hello()
	b.SayOk()

	b.Name = "Tome"

	// 一下两个访问Name效果是一样的
	fmt.Println(b.A.Name) // Jack
	fmt.Println(b.Name)   // Tom

}
