package main

import (
	"fmt"
	"encoding/json"
)


type A struct {
	Num int
	
}

type B struct {
	Num int
	
}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}


type Point struct {
	x int
	y int
}

func main(){
	var a A
	var b B	
	a = A(b)
	fmt.Println(a, b)

	monster := Monster{"brandon", 500, "eat"}
	// 序列化
	jsonMonster, err:= json.Marshal(monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonMonster))
}