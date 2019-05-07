package main

import(
	"fmt"
	
)

func modify(map1 map[int]int){
	map1[10] = 10000
}


// 定义结构体
type Stu struct{
	Name string
	Age int
	Address string
}

func main (){
	// map1 := make(map[int]int)
	// map1 := map[int]int{
	// 	1:90,
	// 	2:88,
	// 	10:1,
	// 	20:2,
	// }
	// fmt.Println(map1)
	// modify(map1)
	// fmt.Println(map1)

	// map1 := make(map[int]int, 2)
	// map1[1] = 100
	// map1[2] = 200

	// map1[3] = 300
	// map1[4] = 400
	// fmt.Println(len(map1))
	// fmt.Println(map1)

	students := make(map[string]Stu)
	// 创建学生
	stu1 := Stu{"brandon",22, "shanghai"}
	stu2 := Stu{"dodo",21, "beijing"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	for k, v := range students{
		fmt.Printf("学生的编号=%v，名字=%v，年龄=%v,住址=%v\n", k, v.Name, v.Age, v.Address)
	}



}