/**
* @Author: YuDC
* @Date: 2019-07-21 14:27
* @Description:
 */
package model

import "fmt"

type account struct {
	accountType string
	balance     float64
	inCome      float64
	desc        string
	list        []account
}

// 创建一笔流水
func InitAccount(balance float64) *account {
	return &account{
		balance: balance,
	}
}

func NewAccount(accountType string, balance float64, inCome float64, desc string) *account {
	return &account{
		accountType: accountType,
		balance:     balance,
		inCome:      inCome,
		desc:        desc,
	}
}

func (account *account) AddAccount(a account) {
	account.list = append(account.list, a)
}

func (account *account) GetAccountList() []account {
	return account.list
}

func (account *account) GetSumBalance() float64 {
	total := 0.0
	for _, v := range account.list {
		total += v.inCome
	}
	return total
}

func (account account) String() string {
	return fmt.Sprintf("%s\t%.2f\t%.2f\t%s\n",
		account.accountType, account.balance, account.inCome, account.desc)
}
