package main

import (
	"fmt"
)

func main(){
	// 1. 定义数组
	var hens [6]float64
	// 2.赋值
	hens[0] = 3.0
	hens[1] = 3.6
	hens[2] = 5.6
	hens[3] = 7.5
	hens[4] = 2.9
	hens[5] = 8.1

	var sum float64
	for i := 0; i < len(hens); i++{
		sum += hens[i]
	}

	avg := fmt.Sprintf("%.2f", sum / float64(len(hens)))
	fmt.Printf(avg)

}