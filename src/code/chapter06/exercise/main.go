package main

import(
	"fmt"
)

func main(){
	// var sum int
	// var count int
	// for i := 1; i <= 100; i++{
	// 	if i % 9 == 0 {
	// 		sum += i
	// 		count++
	// 	}
	// }
	// fmt.Printf("1~100内总和：%d,9的倍数为：%d", sum, count)

	var i int = 0
	var j int = 6
	var res int
	for i <= 6{
		res = i * j
		fmt.Printf("%d * %d = %d\n", i, j, res)
		i++
		j--
	}


}