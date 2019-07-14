/**
* @Author: YuDC
* @Date: 2019-07-11 22:42
* @Description:
 */
package model

type account struct {
	accountNo string
	password  string
	balance   float64
}

func NewAccount(accountNo string, password string, balance float64) *account {
	return &account{
		accountNo: accountNo,
		password:  password,
		balance:   balance,
	}
}

// 继承
// person 拥有account的属性
type person struct {
	account
	name string
	age  int
	addr string
}

func NewPerson(account *account, name string, age int, addr string) *person {
	return &person{
		account: *account,
		name:    name,
		age:     age,
		addr:    addr,
	}
}

func (person *person) GetAccount() account {
	return person.account
}

func (person *person) SetAccountNo(accountNo string) {
	person.accountNo = accountNo
}
