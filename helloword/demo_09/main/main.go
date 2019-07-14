/**
* @Author: YuDC
* @Date: 2019-07-11 00:57
* @Description:
 */
package main

import (
	"awesomeProject/helloword/demo_09/model"
	"fmt"
)

func main() {
	// 创建student
	stu := model.CreateStudent("猪八戒", 430)

	fmt.Println(*stu) // {猪八戒 430}

	fmt.Println(stu.GetName()) // 猪八戒

	fmt.Println(stu.GetAge()) // 430

	stu.SetName("孙悟空")

	fmt.Println(stu.GetName()) // 孙悟空

	account := model.NewAccount("123456", "123456", 50000)

	fmt.Println("account: ", *account)

	person := model.NewPerson(account, "与大潮", 26, "北京")

	fmt.Println(person)

	person.SetAccountNo("135780")

	fmt.Println(person.GetAccount())

	animal := model.NewAnimal("小花猫", 1, 10.0)
	cat := model.NewCat(animal, "白色")

	//cat.Animal.name = "小猫" //  cat.Animal.name undefined
	cat.Animal.Eat()
	cat.Jump()
	fmt.Println(*cat)

}
