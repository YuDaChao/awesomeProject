/**
* @Author: YuDC
* @Date: 2019-07-28 11:43
* @Description:
 */
package model

import "fmt"

type Customer struct {
	id     int
	name   string
	gender string
	age    int
	phone  string
	email  string
}

func NewCustomer(name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}

func (customer *Customer) GetId() int {
	return customer.id
}

func (customer *Customer) SetId(id int) {
	customer.id = id
}

func (customer *Customer) SetName(name string) {
	customer.name = name
}

func (customer Customer) String() string {
	return fmt.Sprintf("%d\t\t\t\t%s\t\t\t\t%s\t\t\t\t%d\t\t\t\t%s\t\t\t\t%s",
		customer.id, customer.name, customer.gender, customer.age, customer.phone, customer.email)
}
