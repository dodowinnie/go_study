package main
import (
	"fmt"
)

func main () {
	// var totalScore int
	// var tmp int
	// var i int = 1
	// for ; i <= 3; i++{
	// 	var score int
	// 	var j int = 1
	// 	for ; j <= 5; j++ {
	// 		fmt.Printf("请输入%d班学生的成绩", i)
	// 		fmt.Scanln(&tmp)
	// 		score += tmp
	// 	}
	// 	fmt.Printf("%d班的平均分是%v\n", i, score/(j - 1))
	// 	totalScore += score

	// }
	// fmt.Printf("总平均分是%v", totalScore/(i - 1))

	// var ceng int
	// fmt.Println("请输入层数")
	// fmt.Scanln(&ceng)
	// for i := 1; i <= ceng; i++ {
	// 	for k := 1; k <= ceng - i; k++{
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= 2 * i - 1; j++ {
			
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// 打印99乘法表
	for i := 1; i < 10; i++{
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}




}