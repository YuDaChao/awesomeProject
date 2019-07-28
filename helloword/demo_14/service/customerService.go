/**
* @Author: YuDC
* @Date: 2019-07-28 11:50
* @Description:
 */
package service

import (
	"awesomeProject/helloword/demo_14/model"
)

type CustomerService struct {
	list  []model.Customer
	total int
}

func NewCustomerService() *CustomerService {
	//// 创建一个customer实例
	//customer := model.NewCustomer(1, "张三", "男", 23, "0101", "123@qq.com")
	//customerService := &CustomerService{}
	//customerService.list = append(customerService.list, *customer)
	//customerService.total = len(customerService.list)
	//return customerService
	return &CustomerService{}
}

// 返回客户列表
func (cs *CustomerService) List() []model.Customer {
	return cs.list
}

// 添加一个客户
func (cs *CustomerService) AddCustomer(customer model.Customer) {
	cs.total++
	customer.SetId(cs.total)
	cs.list = append(cs.list, customer)
}

// 删除客户
func (cs *CustomerService) DeleteCustomer(id int) bool {
	var success bool
	for i, v := range cs.list {
		if v.GetId() == id {
			success = true
			cs.list = append(cs.list[0:i], cs.list[i+1:len(cs.list)]...)
		}
	}
	return success
}

// 修改客户信息
func (cs *CustomerService) UpdateCustomer(id int, name string) {
	//for _, v := range cs.list {
	//	if v.GetId() == id {
	//		v.SetName(name)
	//	}
	//}
	for i := 0; i < len(cs.list); i++ {
		if cs.list[i].GetId() == id {
			cs.list[i].SetName(name)
		}
	}
}
