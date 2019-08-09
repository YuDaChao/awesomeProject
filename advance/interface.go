/**
* @Author: YuDC
* @Date: 2019-08-09 10:45
* @Description: 接口
 */
package main

import (
	"fmt"
	"math"
)

// 定义一个接口
type Car interface {
	Run()
	GetName() string
}

// 定一个几何图形接口
// 面积 周长
type geometry interface {
	area() float64
	perimeter() float64
}

// 矩形
type rect struct {
	width, height float64
}

// 圆
type circle struct {
	radius float64
}

type BMW struct {
	Name string
}

// 让BMW实现Car接口
// 实现接口必须实现接口中的所有方法
func (b *BMW) GetName() string {
	return b.Name
}

func (b *BMW) Run() {
	fmt.Printf("%v 正在跑...", b.Name)
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {

	//car := BMW{Name: "宝马"}
	//name := car.GetName()
	//fmt.Println(name)
	//car.Run()

	//var car Car = &BMW{"奔驰"}
	//
	//fmt.Println(car.GetName())
	//car.Run()

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)

}
