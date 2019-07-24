/**
* @Author: YuDC
* @Date: 2019-07-21 10:42
* @Description: 小项目练习
 */
package main

import (
	"awesomeProject/helloword/demo_12/model"
	"fmt"
)

func main() {
	var key string
	loop := true
	newAccount := model.InitAccount(0.0)
	for {
		list := newAccount.GetAccountList()
		total := newAccount.GetSumBalance()
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("----------1：收支明细----------")
		fmt.Println("----------2：登记收入----------")
		fmt.Println("----------3：登记支出----------")
		fmt.Println("----------4：退   出----------")
		fmt.Println("请选择（1-4）：")
		_, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Print("请输入正确选项")
		}
		switch key {
		case "1":
			fmt.Println("----------收支明细----------")
			fmt.Println("收支\t账户金额\t收支金额\t说明")
			for _, account := range list {
				fmt.Println(account)
			}
		case "2":
			fmt.Println("----------登记收入----------")
			fmt.Println("请输入收入金额")
			inCome := 0.0
			_, err1 := fmt.Scanln(&inCome)
			if err1 != nil {
				fmt.Println("输入金额错误")
			}
			fmt.Println("请输入说明")
			desc := ""
			_, err2 := fmt.Scanln(&desc)
			if err2 != nil {
				fmt.Println("输入说明")
			}
			a := model.NewAccount("收入", total+inCome, inCome, desc)
			newAccount.AddAccount(*a)
		case "3":
			fmt.Println("----------登记支出----------")
			fmt.Println("请输入支出金额")
			inCome := 0.0
			_, err1 := fmt.Scanln(&inCome)
			if err1 != nil {
				fmt.Println("输入金额错误")
			}
			fmt.Println("请输入说明")
			desc := ""
			_, err2 := fmt.Scanln(&desc)
			if err2 != nil {
				fmt.Println("输入说明")
			}
			a := model.NewAccount("支出", total-inCome, -inCome, desc)
			newAccount.AddAccount(*a)
		case "4":
			loop = false
		default:
			fmt.Println("请输入正确选项")
		}
		if !loop {
			break
		}
	}
}
