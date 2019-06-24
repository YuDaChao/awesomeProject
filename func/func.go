package main

import "fmt"

// 定义一个函数 无参数 无返回值

func test() {
	fmt.Printf("这是一个无参数 无返回值的函数 \n")
}

// 有参数 需要指定参数类型

func say(a int) {
	fmt.Printf("函数参数为 a = %d \n", a)
}

// 有参数 又返回值

func getName(name string) (newName string) {
	var helloName = "hello " + name
	return helloName
}

// 多个参数 多个返回值
func getNames(firstName string, lastName string) (newFirstName string, newLastName string) {
	newFirstName = "xxx"
	newLastName = "yyy"
	return firstName, lastName
}

func sum(num int) (total int) {
	if num <= 1 {
		return 1
	}
	return num + sum(num-1)
}

func main() {
	// 调用函数
	test()
	say(123)
	name := getName("go")
	fmt.Printf("name = %q \n", name)
	firstName, lastName := getNames("jamson", "haden")
	fmt.Printf("firstName = %q, lastname = %q \n", firstName, lastName)
	sum := sum(100)
	fmt.Printf("sum = %d \n", sum)
}
