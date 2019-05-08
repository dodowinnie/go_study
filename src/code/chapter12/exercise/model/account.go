package model

import(
	"fmt"
)

type account struct {
	AccountNo string
	balance float64
	password string
}

func NewAccount(accountNo string, password string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账户长度6~10")
		return nil
	}

	if len(password) != 6 {
		fmt.Println("密码设置错误")
		return nil
	}

	if balance < 20 {
		fmt.Println("金额不对")
		return nil
	}

	return &account{
		AccountNo:accountNo,
		password:password,
		balance:balance,
	}
	
}

func (acc *account) setPassword(password string) {
	if len(password) != 6 {
		fmt.Println("密码设置错误")
	}else{
		acc.password = password
	}
}


func (acc *account) SetBalance(balance float64, password string){
	if password != acc.password {
		fmt.Println("密码错误")
	}else {
		acc.balance = balance
	}
}

func (acc *account) GetBalance( password string) float64{
	if password != acc.password {
		fmt.Println("密码错误")
		return 0
	}else {
		return acc.balance
	}
}

