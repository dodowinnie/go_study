package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string 
	Age int 
	Birhtday string 
	Sal float64 
	Skill string 
}

func unmarshalStruct(){
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后monster=%v, monster.Name=%v\n", monster, monster.Name)
}

func unmarshalMap(){
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	// amap := make(map[string]interface{})
	var amap map[string]interface{}
	err := json.Unmarshal([]byte(str), &amap)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后:%v\n", amap)
}

func unmarshalSlice(){
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后:%v\n", slice)



}



func main(){
	// unmarshalStruct()
	// unmarshalMap()
	unmarshalSlice()
	// float64Test()
}