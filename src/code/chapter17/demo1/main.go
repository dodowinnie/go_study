package main

import(
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string `json:"monster_name"`
	Age int `json:"monster_age"`
	Sal float64 `json:"monster_salary"`
	Sex string `json:"monster_sex"`
}

func main(){
	monster := Monster{
		Name : "brandon",
		Age : 20,
		Sal : 4565.54,
		Sex : "male",
	}

	data,err := json.Marshal(monster)
	if err != nil{
		fmt.Println("转化出错， err=", err)
	}
	fmt.Println(string(data))

}