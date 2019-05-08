package main

import (
	"fmt"
)

type Usb interface{
	Start()
	Stop()
}

type Phone struct{
	
}
// 实现接口方法
func (phone Phone) Start(){
	fmt.Println("手机开始工作")
}

func (phone Phone) Stop(){
	fmt.Println("手机停止工作")
}

type Camera struct{
	
}
// 实现接口方法
func (camera Camera) Start(){
	fmt.Println("相机开始工作")
}

func (camera Camera) Stop(){
	fmt.Println("相机停止工作")
}

type Computer struct{
	
}
// 实现接口方法
func (computer Computer) Working(usb Usb){
	usb.Start()
	fmt.Println("计算机开始工作")
	usb.Stop()
}


func main()  {
	var phone = Phone{}
	var computer = Computer{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}