/**
* @Author: YuDC
* @Date: 2019-07-14 12:44
* @Description: 继承
 */
package model

import "fmt"

// 定义一个动物结构体

type Animal struct {
	name   string
	age    int
	weight float64
}

// 动物可以吃
func (animal *Animal) Eat() {
	fmt.Println("animal eat ....")
}

func NewAnimal(name string, age int, weight float64) *Animal {
	return &Animal{
		name:   name,
		age:    age,
		weight: weight,
	}
}

// 定义一个猫结构体

type cat struct {
	Animal
	color string
}

// 定义猫自己的方法

func (cat *cat) Jump() {
	fmt.Println("cat jump")
}

// 工厂方法
func NewCat(animal *Animal, color string) *cat {
	return &cat{
		Animal: *animal,
		color:  color,
	}
}
