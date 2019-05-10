package service

import (
	// "fmt"
	"code/customer_manage/model"
)

type CustomerService struct {
	customers []model.Customer
	// 表示当前customer切片有多少客户
	customerNum int

}

// 工厂方式返回service实例
func NewCustomerService() *CustomerService{
	customerService := &CustomerService{}
	customerService.customerNum = 2
	customer := model.NewCustomer(1, "brandon", "男",20,"15021147853","brandon@163.com")
	customerService.customers = append(customerService.customers, customer)
	customer2 := model.NewCustomer(2, "dodo", "女",20,"15021147853","dodo@163.com")
	customerService.customers = append(customerService.customers, customer2)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}


// 根据id查询用户,返回对应的下标，没有返回-1
func (this *CustomerService) FindById(id int) int{
	index := -1
	for in, customer := range this.customers {
		if id == customer.Id {
			index = in
		}
	}
	return index
}

// 删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1{
		return false
	}else {
		this.customers = append(this.customers[:index], this.customers[index + 1:]...)
		this.customerNum--
		return true
	}
}

func (this *CustomerService) Update(customer model.Customer) bool{
	id := customer.Id
	for _ , cus := range this.customers {
		if id == cus.Id {
			newCustomer := model.Customer{}
			newCustomer.Id = id

			if customer.Name != ""{
				newCustomer.Name = customer.Name
			}else {
				newCustomer.Name = cus.Name
			}

			if customer.Age != 0{
				newCustomer.Age = customer.Age
			}else {
				newCustomer.Age = cus.Age
			}

			if customer.Gender != ""{
				newCustomer.Gender = customer.Gender
			}else {
				newCustomer.Gender = cus.Gender
			}

			if customer.Phone != ""{
				newCustomer.Phone = customer.Phone
			}else {
				newCustomer.Phone = cus.Phone
			}

			if customer.Email != ""{
				newCustomer.Email = customer.Email
			}else {
				newCustomer.Email = cus.Email
			}

			// 先删除
			this.Delete(id)
			// 在插入
			this.customers = append(this.customers, newCustomer)
			this.customerNum++
			
		}
	}
	return true


}
