package main

import(
	"fmt"
	"sort"
)

func main (){
	// map1 := make(map[int]int, 10)
	map1 := map[int]int{
		10:100,
		1:13,
		4:56,
		8:90,
	}

	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(map1[k])
	}


}