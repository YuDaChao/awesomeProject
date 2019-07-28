/**
* @Author: YuDC
* @Date: 2019-07-28 12:03
* @Description:
 */
package main

import (
	"awesomeProject/helloword/demo_14/model"
	"awesomeProject/helloword/demo_14/service"
	"fmt"
)

type CustomerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

// 显示客户列表
func (cv *CustomerView) showCustomerList() {
	list := cv.customerService.List()
	fmt.Println("---------------------客户列表---------------------")
	fmt.Println("编号\t\t\t\t姓名\t\t\t\t性别\t\t\t\t年龄\t\t\t\t电话\t\t\t\t邮箱")
	for _, c := range list {
		fmt.Println(c)
	}
}

// 添加一个客户
func (cv *CustomerView) addCustomer(customer model.Customer) {
	cv.customerService.AddCustomer(customer)
}

// 删除客户
func (cv *CustomerView) deleteCustomer() {
	fmt.Println("请输入id: ")
	id := 0
	_, _ = fmt.Scanln(&id)
	success := cv.customerService.DeleteCustomer(id)
	fmt.Println(success)
}

// 修改客户信息
func (cv *CustomerView) updateCustomer() {
	fmt.Println("请输入id， name")
	var (
		id   int
		name string
	)
	_, _ = fmt.Scanln(&id)
	_, _ = fmt.Scanln(&name)
	cv.customerService.UpdateCustomer(id, name)
}

func (cv *CustomerView) showMainMenu() {
	for {
		fmt.Println("----------客户管理系统----------")
		fmt.Println("          1 添加客户")
		fmt.Println("          2 删除客户")
		fmt.Println("          3 修改客户")
		fmt.Println("          4 客户列表")
		fmt.Println("          5 退   出")
		fmt.Print("请输入1-5：")
		_, _ = fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			var (
				name   string
				gender string
				age    int
				phone  string
				email  string
			)
			fmt.Println("请一次输入： name，gender，age，phone，email")
			_, _ = fmt.Scanln(&name)
			_, _ = fmt.Scanln(&gender)
			_, _ = fmt.Scanln(&age)
			_, _ = fmt.Scanln(&phone)
			_, _ = fmt.Scanln(&email)
			customer := model.NewCustomer(name, gender, age, phone, email)
			cv.addCustomer(*customer)
		case "2":
			cv.deleteCustomer()
		case "3":
			cv.updateCustomer()
		case "4":
			cv.showCustomerList()
		case "5":
			cv.loop = false
		default:
			fmt.Println("default")
		}
		if !cv.loop {
			break
		}
	}
}

func main() {
	cv := &CustomerView{loop: true, customerService: service.NewCustomerService()}
	cv.showMainMenu()
}
