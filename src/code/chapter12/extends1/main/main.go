package main

import (
	"fmt"
)

type A struct {
	Name string
	age int
}

func (a *A) Sayok() {
	fmt.Println("A sayok", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}


type B struct{
	A
}

func (b *B) hello() {
	fmt.Println("B hello", b.Name)
}

func (b *B) Sayok() {
	fmt.Println("B sayok", b.Name)
}

func main()  {
	b := &B {}
	b.A.Name = "brandon"
	b.A.age = 18
	b.A.Sayok()
	b.A.hello()

	// 简化
	b.Name = "dodo"
	b.age = 28
	b.Sayok()
	b.hello()


}