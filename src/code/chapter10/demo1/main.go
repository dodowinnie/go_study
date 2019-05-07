package main

import(
	"fmt"
)

func main (){
	// var a map[string]string
	// a = make(map[string]string, 10)
	// a["name"] = "brandon"
	// a["address"] = "shanghai"
	// fmt.Println(a)

	// cities := make(map[string]string)
	// cities["north"]="北京"
	// cities["south"]="福建"
	// fmt.Println(cities)

	// heroes := map[string]string{
	// 	"h1":"brandno",
	// 	"h2":"dodo",
	// }
	// fmt.Println(heroes)
	students := make(map[string]map[string]string)
	// students["s1"] = make(map[string]string)
	// students["s1"]["name"] = "tom"
	// students["s1"]["sex"] = "male"
	
	// students["s2"] = make(map[string]string)
	// students["s2"]["name"] = "dodo"
	// students["s2"]["sex"] = "female"
	students["s1"] = map[string]string{
	"name" : "tom",
	"sex" : "male",
	}

	fmt.Println(students["s1"]["name"])

}