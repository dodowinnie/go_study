package main

import(
	"fmt"
)

type Usb interface{
	Start()
	Stop()
}

type Phone struct{
	Name string
}
func (phone Phone) Start(){
	fmt.Println(phone.Name, "手机开始工作。。。。")
}

func (phone Phone) Stop(){
	fmt.Println(phone.Name, "手机停止工作。。。。")
}


type Camera struct{
	Name string
}
func (camera Camera) Start(){
	fmt.Println(camera.Name, "相机开始工作。。。。")
}

func (camera Camera) Stop(){
	fmt.Println(camera.Name, "相机停止工作。。。。")
}

func main(){
	var usbs [3]Usb
	usbs[0] = Phone{"vivo"}
	usbs[1] = Phone{"小米"}
	usbs[2] = Camera{"vivo"}

	fmt.Println(usbs)

}