/**
* @Author: YuDC
* @Date: 2019-07-14 21:08
* @Description: 接口
 */
package main

import "fmt"

// 定义一个Usb接口

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

// 让Phone实现Usb的 所有 接口方法
func (phone Phone) Start() {
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Compute struct {
}

func (compute Compute) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// 空接口 任何接口都可以赋值给空接口
type T interface {
}

func main() {
	compute := Compute{}
	phone := Phone{}

	compute.Working(phone)

	var t T = 1

	fmt.Println(t)

	var u Usb = Phone{}

	fmt.Println(u)
}
