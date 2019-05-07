package main

import(
	"fmt"
)

func main (){
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	// // fmt.Println(cities)
	cities["no4"] = "深圳"
	// // delete(cities, "no5")
	fmt.Println(len(cities))
	
	// // cities = make(map[string]string)
	// // fmt.Println(cities)

	// // val := cities["no5"]
	// // fmt.Println(val)
	// // fmt.Println(res)

	// // 遍历
	// for k, v := range cities {
	// 	fmt.Printf("key=%v, value=%v\n", k, v)
	// }

	// students := make(map[string]map[string]string)
	// students["s1"] = map[string]string{
	// 	"name":"brandon",
	// 	"sex":"man",
	// }
	// students["s2"] = map[string]string{
	// 	"name":"dodo",
	// 	"sex":"women",
	// }

	// for k, v := range students {
	// 	fmt.Println("k1=", k)
	// 	for k1, v1 := range v {
	// 		fmt.Printf("\n k2:v2 = %v:%v",k1,v1)
	// 	}
	// 	fmt.Println()
	// }

}