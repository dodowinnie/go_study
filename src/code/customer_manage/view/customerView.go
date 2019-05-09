package main

import (
	"fmt"
	"code/customer_manage/service"
)

type customerView struct{
	key string // 接受用户输入
	loop bool // 表示是否循环显示主菜单
	customerService *service.CustomerService
}

func (this *customerView) list(){
	customerList := this.customerService.List()
	// 显示
	fmt.Println("-------------客户信息管理软件-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range customerList {
		customer.GetInfo()
	}
}


// 显示主菜单

func (this *customerView) mainMenu(){
	for {
		fmt.Println("-------------客户信息管理软件-------------")
		fmt.Println()
		fmt.Println("                1.添加客户")
		fmt.Println("                2.修改客户")
		fmt.Println("                3.删除客户")
		fmt.Println("                4.客户列表")
		fmt.Println("                5.退   出")
		fmt.Println()
		fmt.Print("                请选择(1-5):")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			this.list()
		case "5":
			fmt.Println("感谢使用客户信息管理软件,再见")
			this.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
			
		}
		
		if !this.loop {
			break
		}
	}

} 


func main(){
	customerView := customerView{
		key:"",
		loop:true,
	}
	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()

}