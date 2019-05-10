package main

import (
	"fmt"
	"code/customer_manage/service"
	"code/customer_manage/model"
)

type customerView struct{
	key string // 接受用户输入
	loop bool // 表示是否循环显示主菜单
	customerService *service.CustomerService
}

// 客户列表
func (this *customerView) list(){
	customerList := this.customerService.List()
	// 显示
	fmt.Println("-------------客户信息管理软件-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, customer := range customerList {
		fmt.Println(customer.GetInfo())
	}
}

// 添加客户
func (this *customerView) add(){
	fmt.Println("-------------添加客户-------------")

	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话:")
	tel := ""
	fmt.Scanln(&tel)

	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, tel, email)
	res := this.customerService.Add(customer)
	if res {
		fmt.Println("-------------添加完成-------------")
	}else {
		fmt.Println("-------------添加失败-------------")
	}
}

// 删除客户
func (this *customerView) delete(){
	fmt.Println("-------------删除客户-------------")
	fmt.Println("请选择待删除的客户编号(-1)退出:")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("-----------删除客户取消-----------")
		return
	}
	fmt.Println("确认是否删除(Y/N):")
	flag := ""
	fmt.Scanln(&flag)
	if flag == "y" || flag == "Y" {
		res := this.customerService.Delete(id)
		if res {
			fmt.Println("-------------删除客户成功-------------")
		}else {
			fmt.Println("-------------删除客户失败-------------")
		}
	}

}

// 修改客户
func (this *customerView) update(){
	fmt.Println("-------------修改客户-------------")
	fmt.Println("请选择待修改的客户编号(-1)退出:")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("-----------删除客户取消-----------")
		return
	}

	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话:")
	tel := ""
	fmt.Scanln(&tel)

	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer(id, name, gender, age, tel, email)
	res := this.customerService.Update(customer)
	if res {
		fmt.Println("-------------修改完成-------------")
	}else {
		fmt.Println("-------------修改失败-------------")
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
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			for{
				fmt.Print("确认退出(Y/N):")
				flag := ""
				fmt.Scanln(&flag)
				if flag == "y" || flag == "Y"{
					fmt.Println("感谢使用客户信息管理软件,再见")
					this.loop = false
					break
				}else if flag == "n" || flag == "N" {
					break
				}else {
					fmt.Println("您的输入有误，请重新输入")
				}
			}
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