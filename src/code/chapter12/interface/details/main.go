package main

import (
	"fmt"
)

type AInterface interface{
	Say()
}

type Stu struct {

}

func (stu Stu) Say(){
	fmt.Println("stu say()...")
}

type Integer int

func (i Integer) Say(){
	fmt.Println("stu say i=", i)
}

type BInterface interface{
	Hello()
}

type Monster struct{

}

func (monster Monster) Hello(){
	fmt.Println("moster hello")
}


func (monster Monster) Say(){
	fmt.Println("moster say")
}

func main()  {
	// stu := Stu{}
	// var a AInterface = stu
	// a.Say()

	// var i Integer
	// var b AInterface = i
	// b.Say()

	moster := Monster{}
	var a2 AInterface = moster
	var b2 BInterface = moster
	a2.Say()
	b2.Hello()
}