package main
import (
	"fmt"
)

func main(){
	// var age int
	// fmt.Println("请输入你的年龄")
	// fmt.Scanln(&age)
	// if age >= 18 {
	// 	fmt.Println("你的年龄大于18岁")
	// }


	
	// if age:= 20; age >= 18 {
	// 	fmt.Println("你的年龄大于18岁")
	// }

	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年龄大于18岁")
	} else {
		fmt.Println("你的年龄不到18岁")
	}
}