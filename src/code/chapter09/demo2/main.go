package main

import (
	"fmt"
)

func main(){
	// var slice []float64 = make([]float64, 5, 10)
	// slice[2] = 10
	// slice[3] = 20
	// fmt.Println(slice)
	// fmt.Println("slice size:", cap(slice))
	// fmt.Println("slice length:", len(slice))

	// var slice []string = []string{"tom", "jack","mary"}
	// fmt.Println(slice)
	// fmt.Println("slice size", cap(slice))
	// fmt.Println("slcie len", len(slice))

	// var arr [5]int = [...]int{10,20,30,40,50}

	// slice := arr[1:4]
	// slice := arr[:]
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("slice[%v]=%v\n", i, slice[i])
	// }

	// slice2 := slice[1:3]
	// slice2[0] = 88
	// fmt.Println(&slice[0])
	// slice = append(slice, 99)
	// fmt.Println(&slice[0])

	// for index, value := range slice {
	// 	fmt.Printf("slice[%v]=%v\n", index, value)
	// }

	a := []int{1,2,3,4,5}

	var slice = make([]int, 10)
	slice[8]=100
	slice[9]=200
	copy(slice, a)
	fmt.Println(slice)


}