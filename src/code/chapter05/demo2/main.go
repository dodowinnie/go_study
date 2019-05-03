package main
import (
	"fmt"
	// "math"
)

func main(){
	// var score float32
	// fmt.Println("请输入分数")
	// fmt.Scanln(&score)
	// if score == 100 {
	// 	fmt.Println("奖励BMW一套")
	// } else if score > 80 && score <= 99 {
	// 	fmt.Println("奖励Iphone7 plus一套")
	// } else if score > 60 && score <= 80 {
	// 	fmt.Println("奖励Ipad一套")
	// } else {
	// 	fmt.Println("奖励 nothing")
	// }

	// var flag = false
	// if flag == false {
	// 	fmt.Println("flag is false")
	// }
	// var (
	// 	a float64 = 3.0
	// 	b float64 = 10.0 
	// 	c float64 = 6.0
	// ) 
	
	// m := b * b - 4 * a * c
	// fmt.Println(math.Sqrt(b))
	// if m > 0 {
	// 	x1 := (-b + math.Sqrt(m)) / 2 * a
	// 	x2 := (-b - math.Sqrt(m)) / 2 * a
	// 	fmt.Printf("x1=%v x2=%v", x1, x2)
	// } else if m == 0 {
	// 	x1 := (-b + math.Sqrt(m)) / 2 * a
	// 	fmt.Printf("x1=x2=%v", x1)
	// }else {
	// 	fmt.Printf("无解")
	// }

	var (
		height int32
		money float32
		handsome string
		isHandsome bool
	)

	fmt.Println("请输入身高(cm)")
	fmt.Scanln(&height)

	fmt.Println("请输入身价（千万）")
	fmt.Scanln(&money)

	fmt.Println("是否帅 请输入yes或no")
	fmt.Scanln(&handsome)
	if handsome == "yes"{
		isHandsome = true
	} else {
		isHandsome = false
	}

	if height > 180 && money > 1.0 && isHandsome {
		fmt.Println("must marry")
	}else if height > 180 || money > 1.0 || isHandsome{
		fmt.Println("maybe marry")
	}else {
		fmt.Println("can not marry")
	}




}