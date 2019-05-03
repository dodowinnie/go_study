package main

import(
	"fmt"
)

func main(){
	n1 := 9
	n2 := 8

	fmt.Println(n1 == n2) // false
	fmt.Println(n1 != n2) // true
	fmt.Println(n1 <= n2) // fasle
	fmt.Println(n1 >= n2) // true
	fmt.Println(n1 < n2) // false
	fmt.Println(n1 > n2) // true
	flag := n1 >= n2
	fmt.Println(flag)

}