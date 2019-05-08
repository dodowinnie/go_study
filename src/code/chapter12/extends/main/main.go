package main

import (
	"fmt"
)

type Student struct {
	Name string
	Score float64
	Age int
}

type Pupil struct {
	Student // 嵌入Student匿名结构体
}

type Graduate struct {
	Student // 嵌入Student匿名结构体
}

func (stu *Student) showInfo() {
	fmt.Printf("学生名=%v,年龄=%v，成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) setScore(score float64){
	stu.Score = score
}

func (stu *Student) getSum(n1 int, n2 int) int{
	return n1 + n2
}

func (p *Pupil) testing(){
	fmt.Println("小学生考试中。。。")
}

func (g *Graduate) testing(){
	fmt.Println("大学生考试中。。。")
}


func main()  {
	pupil := &Pupil {}
	pupil.Student.Name = "brandon"
	pupil.Student.Age = 10
	pupil.testing()
	pupil.Student.setScore(80.5)
	pupil.Student.showInfo()
	fmt.Println(pupil.Student.getSum(1,2))

	graduate := &Graduate{}
	graduate.Student.Name = "dodo"
	graduate.Student.Age = 27
	graduate.testing()
	graduate.Student.setScore(100)
	graduate.Student.showInfo()
	fmt.Println(graduate.Student.getSum(10,20))
}