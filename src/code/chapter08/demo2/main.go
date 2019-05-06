package main

import (
	"fmt"
)

func main(){
	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Println("请输入成绩")
	// 	fmt.Scanln(&score[i])
	
	// }
	// fmt.Println(score)

	// var numArr [3]int = [3]int{1,2,3}
	// numArr := [3]int{1,2,3}
	// var numArr = [...]int{4,5,6}
	// var numArr = [...]int{1:800,0:999,3:654}

	heros := [...]string{"brandon", "iron man", "captain American"}
	// for index, value := range heros {
	// 	fmt.Printf("index=%v, value=%v\n", index, value)
	// }

	for _, value := range heros {
		fmt.Printf("value=%v\n", value)
	}
	

}