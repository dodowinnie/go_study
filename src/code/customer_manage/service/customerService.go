package service

import (
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
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "brandon", "男",20,"15021147853","brandon@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

