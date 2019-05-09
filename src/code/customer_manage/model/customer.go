package model

import(
	"fmt"
)

// 定义用户结构体
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 返回Customer实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer{
	return Customer {
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

// 返回用户的信息，格式化的字符串
func (this Customer) GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
