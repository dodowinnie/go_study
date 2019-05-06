package main

import (
	"fmt"
)

func main(){
	var year int
	var month int
	for {
		fmt.Println("请输入年份\n")
		fmt.Scanln(&year)
		fmt.Println("请输入月份\n")
		fmt.Scanln(&month)
		if month > 12 || month < 1 {
			fmt.Println("月份输入有误，请重输入\n")
			continue
		}

		switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Printf("%d年%d月有%d天\n", year, month, 31)
			case 4, 6, 9, 11:
				fmt.Printf("%d年%d月有%d天\n", year, month, 30)
			case 2:
				if year % 400 == 0 {
					// 闰年
					fmt.Printf("%d年%d月有%d天\n", year, month, 29)
				}else{
					fmt.Printf("%d年%d月有%d天\n", year, month, 28)
				}

		}

		



	}
}