package utils

import (
	"fmt"
)

var Num int = 32

func Cal(n1 float64, n2 float64, operator byte) float64 {
	switch operator{
		case '+':
			return n1 + n2
		case '-':
			return n1 - n2
		case '*':
			return n1 * n2
		case '/':
			return n1 / n2
		default:
			fmt.Println("操作符号有误。。。。")
			return 0.0
	}
}