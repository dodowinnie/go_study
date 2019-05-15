package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"monster_name"`
	Age int `json:"monster_age"`
	Birhtday string `json:"monster_birthday"`
	Sal float64 `json:"monster_salary"`
	Skill string `json:"monster_skill"`
}

func test(){
	monster := Monster{
		Name : "brandon",
		Age : 23,
		Birhtday : "2011-04-15",
		Sal : 8000.01,
		Skill : "Eating",
	}

	// 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化出错, ", err)
		return
	}

	fmt.Printf("序列化结果=%v\n", string(data))
}

func mapTest(){
	var amap map[string]interface{}

	amap = make(map[string]interface{})
	amap["name"] = "brandon"
	amap["age"] = 30
	amap["address"] = "上海"

	// 序列化
	data,err := json.Marshal(amap)
	if err != nil {
		fmt.Println("序列化错误 ", err)
	}
	fmt.Printf("amap=%v", string(data))
}

// 切片序列化
func sliceTest(){
	var slice []map[string]interface{}
	m1 := map[string]interface{}{
		"name" : "jack",
		"age" :22,
		"address" : "北京",
	}
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "brandon"
	m2["age"] = 24
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	m2["childres"] = [...]string{"aa", "bb"}
	slice = append(slice, m2)

	// 序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误 ", err)
	}
	fmt.Printf("slice=%v", string(data))


}

// 基本类型序列化
func float64Test(){
	n1 := 235.014
	// 序列化
	data, err := json.Marshal(n1)
	if err != nil {
		fmt.Println("序列化错误 ", err)
	}
	fmt.Printf("基本数据类型=%v", string(data))

}

func main(){
	test()
	// mapTest()
	// sliceTest()
	// float64Test()
}