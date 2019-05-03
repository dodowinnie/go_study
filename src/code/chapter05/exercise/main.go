package main
import (
	"fmt"
)

func main () {
	// var num1 int32 = 42
	// var num2 int32 = 24
	// if num1 + num2 >= 50 {
	// 	fmt.Println("hello world")
	// }

	// var f1 float64 = 15.64
	// var f2 float64 = 5.2
	// if f1 > 10.0 && f2 < 20.0 {
	// 	fmt.Println(f1 + f2)
	// }

	// var n1 int32 = 3
	// var n2 int32 = 12
	// if (n1 + n2) % 3 == 0 && (n1 + n2) % 5 == 0 {
	// 	fmt.Println("既能被三整出又能被5整除")
	// }
	// var year int
	// fmt.Println("请输入年份")
	// fmt.Scanln(&year)

	// if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
	// 	fmt.Printf("%v为闰年", year)
	// } else {
	// 	fmt.Printf("%v不是闰年", year)
	// }
	
	// var key byte
	// fmt.Println("please enter the letter")
	// fmt.Scanf("%c", &key)

	// switch key{
	// 	case 'a':
	// 		fmt.Println("A")
	// 	case 'b':
	// 		fmt.Println("B")
	// 	case 'c':
	// 		fmt.Println("C")
	// 	case 'd':
	// 		fmt.Println("D")
	// 	case 'e':
	// 		fmt.Println("E")
	// 	default:
	// 		fmt.Println("other")
	// }
	
	// var score int
	// fmt.Println("请输入分数")
	// fmt.Scanln(&score)

	// switch int(score/60) {
	// 	case 1:
	// 		fmt.Println("合格")
	// 	case 0 :
	// 		fmt.Println("不合格")
	// 	default:
	// 		fmt.Println("wrong")
	// }

	var month int
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
		case 3, 4, 5:
			fmt.Println("spring")
		case 6, 7, 8:
			fmt.Println("summer")
		case 9, 10, 11:
			fmt.Println("autumn")
		case 12, 1, 2:
			fmt.Println("winter")
		default:
			fmt.Println("wrong input")
	}



}