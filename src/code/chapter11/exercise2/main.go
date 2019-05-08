package main

import (
	"fmt"

)

type Student struct{
	Name string
	Gender string
	Age int
	Id int
	Score float64

}


func (spt *Student) say() string{
	str := fmt.Sprintf("name=%v,gender=%v,age=%v,id=%v,score=%v", spt.Name, spt.Gender, spt.Age, spt.Id, spt.Score)
	return str
}


func main() {
	sutudent := Student {"brandon", "male", 25, 10001, 98.9}
	str := sutudent.say()
	fmt.Println(str)
}