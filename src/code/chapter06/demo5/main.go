package main
import (
	"fmt"
)

func main () {
	// label:
	// for i := 0; i < 4; i++ {
	// 	// label1: // 设置标签
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue label
	// 		}
	// 		fmt.Println("j=", j)
	// 	}

	// }

	// for i := 1; i <= 100; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	var num int
	var positive int
	var negative int
	for {
		fmt.Println("请输入数字")
		fmt.Scanln(&num)
		if num > 0 {
			positive++
			continue
		}else if num < 0{
			negative++
			continue
		}else{
			break
		}

	}
	fmt.Printf("正数的个数%v,负数的个数%v", positive, negative)
}