package model

import (
	"fmt"
)

type person struct{
	Name string
	age int
	sal float64
}

func NewPerson(name string) *person{
	return &person{
		Name:name,
	}
}

func (ppt *person) SetAge(age int){
	if age >0 && age < 150 {
		ppt.age = age
	}else {
		fmt.Println("年龄范围不正确。。。")
	}
}

func (ppt *person) GetAge() int{
	return ppt.age
}

func (ppt *person) SetSal(sal float64){
	if sal >= 3000 && sal <= 30000 {
		ppt.sal = sal
	}else {
		fmt.Println("薪水输入不正确。。。")
	}
}

func (ppt *person) GetSal() float64{
	return ppt.sal
}

