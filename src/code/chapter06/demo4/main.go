package main
import (
	"fmt"
	// "math/rand"
	// "time"
)

func main () {
	// var num int
	// var count int
	// for {
	// 	rand.Seed(time.Now().Unix())
	// 	num = rand.Intn(100) + 1
	// 	if num == 99 {
	// 		count++
	// 		fmt.Printf("总共试了%d次", count)
	// 		break
	// 	}
	// 	count++
	// }
	// fmt.Println(time.Now().Unix())

	// label2:// 设置标签
	// for i := 0; i < 4; i++{
	// 	// label1: // 设置标签
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			break label2
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// }
	// var sum int
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// 	if sum > 20 {
	// 		fmt.Printf("当期数是：%v,和为：%v", i, sum)
	// 		break
	// 	}
	// }

	
	var username string
	var password string
	label:
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		if username == "brandon" && password == "123"{
			fmt.Println("登录成功")
			break label
		}else{
			
			fmt.Printf("用户名或密码错误，还有%v次机会\n", 3-i)
		}

	}



}