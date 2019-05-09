package main

import(
	"fmt"
)

type Monkey struct {
	Name string

}

func (monkey *Monkey) climbing(){
	fmt.Println(monkey.Name, " 生来会爬树。。。。")
}

type BirdAble interface{
	Fly()
}

type FishAble interface{
	Swim()
}

func (monkey *Monkey) Fly(){
	fmt.Println(monkey.Name, " 会飞。。。。")
}

func (monkey *Monkey) Swim(){
	fmt.Println(monkey.Name, " 会游泳。。。。")
}

type LittleMonkey struct {
	Monkey
}



func main()  {
	littleMonkey := LittleMonkey {
		Monkey {
			Name : "brandon",
		},
	}
	littleMonkey.climbing()
	littleMonkey.Fly()
	littleMonkey.Swim()
}