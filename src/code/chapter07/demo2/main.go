package main
import (
	"fmt"
)

func test(n1 int) {
	n1 += 1
	fmt.Println("test n1=", n1)
}

func main(){
	// n1 := 5
	// test(n1)
	// fmt.Println("main n1=", n1)
	// sum := sum(1, 4)
	// fmt.Println("sum=", sum)

	res1, res2 := getSumAndSub(20, 10)
	fmt.Printf("res1=%v, res2=%v", res1, res2)

	// 想要忽略某个返回值，使用 _ 占位符号
	_, res3 := getSumAndSub(40, 10)
	fmt.Printf("res3=%v", res3)


}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}