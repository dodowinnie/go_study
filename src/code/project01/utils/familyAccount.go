package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明一个字段，保存接受用户输入的选项
	key string
	// 控制是否退出for
	loop bool
	// 账户余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 是否收支
	flag bool
	// 收支的详情使用字符串记录
	details string

	// 用户名
	username string
	// 密码
	password string
	loginloop bool

}

func NewFamilyAcocunt() *FamilyAccount{
	return &FamilyAccount{
		key : "",
		loop : true,
		balance:10000.0,
		money:0.0,
		note:"",
		flag:false,
		details:"收支\t账户金额\t收支金额\t说	明",
		username:"brandon", // 默认
		password:"666666", // 默认
		loginloop:true,
	}
}

// 显示明细
func (familyAccount *FamilyAccount) showDetail(){
	fmt.Println("------------------当前收支明细------------------")
	if familyAccount.flag {
		fmt.Println(familyAccount.details)
	}else{
		fmt.Println("当前没有收支明细。。。来一笔")
	}
}

// 收入的方法
func (familyAccount *FamilyAccount) income(){
	fmt.Print("本次收入金额:")
	fmt.Scanln(&familyAccount.money)
	familyAccount.balance += familyAccount.money
	fmt.Print("本次收入说明:")
	fmt.Scanln(&familyAccount.note)
	// 拼接细节
	familyAccount.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", familyAccount.balance, familyAccount.money, familyAccount.note)
	familyAccount.flag = true
}


// 支出的方法
func (familyAccount *FamilyAccount) pay(){
	fmt.Print("本次支出金额:")
	fmt.Scanln(&familyAccount.money)
	if familyAccount.money > familyAccount.balance{
		fmt.Println("余额不足")
		return
	}
	familyAccount.balance -= familyAccount.money
	fmt.Print("本次支出说明:")
	fmt.Scanln(&familyAccount.note)
	// 拼接细节
	familyAccount.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", familyAccount.balance, familyAccount.money, familyAccount.note)
	familyAccount.flag = true
}


func (familyAccount *FamilyAccount) exit(){
	fmt.Println("确定套退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		familyAccount.loop = false
		familyAccount.loginloop = false
		fmt.Println("感谢使用家庭收支记账软件，再见。。。")
	}
}


func (familyAccount *FamilyAccount) Login(){
	for {

		fmt.Println("欢迎使用家庭收支记账软件")
		var username string
		var password string
		fmt.Print("请输入用户名:")
		fmt.Scanln(&username)
		fmt.Print("请输入密码:")
		fmt.Scanln(&password)
	
		if username != familyAccount.username || password != familyAccount.password {
			fmt.Println("用户名或密码错误")
		}else {
			familyAccount.MainMenu()
		}

		if !familyAccount.loginloop {
			break
		}
	}

}


func (familyAccount *FamilyAccount) MainMenu(){
	for {
		
			fmt.Println("------------------家庭收支记账软件------------------")
			fmt.Println()
			fmt.Println("                  1 收支明细                       ")
			fmt.Println("                  2 登记收入                       ")
			fmt.Println("                  3 登记支出                       ")
			fmt.Println("                  4 退    出                       ")
			fmt.Println()
			fmt.Print("                  请选择(1-4):")
			fmt.Scanln(&familyAccount.key)

			switch familyAccount.key {
			case "1":
				familyAccount.showDetail()
			case "2":
				familyAccount.income()
			case "3":
				familyAccount.pay()
			case "4":
				familyAccount.exit()
			}

			if !familyAccount.loop {
				break
			}
		}
	
}
