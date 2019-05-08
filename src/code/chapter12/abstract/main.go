package main

import (
	"fmt"
)

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

func (account *Account) Deposite(money float64, pwd string){
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	if money <= 0 {
		fmt.Println("金额错误")
		return
	}

	account.Balance += money

	fmt.Println("存款成功")
	
}

func (account *Account) Withdraw(money float64, pwd string){
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
	
}

func (account *Account) Query(pwd string){
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	fmt.Printf("你的账号为=%v,余额=%v\n", account.AccountNo, account.Balance)
	
}


func main(){
	account := &Account{
		AccountNo:"brandon",
		Pwd:"123456",
		Balance:100,
	}

	account.Query("123456")
	account.Deposite(200, "123456")
	account.Query("123456")
	account.Withdraw(100, "123456")
	account.Query("123456")
}

