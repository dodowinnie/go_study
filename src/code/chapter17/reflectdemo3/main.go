package main

import(
	"fmt"
	"reflect"
)

type Monster struct{
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex string
}

func (s Monster) GetSum(n1, n2 int) int{
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string){
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func (s Monster) Print(){
	fmt.Println("------start------")
	fmt.Println(s)
	fmt.Println("------end------")
}

func TestStruct(s interface{}){
	// 获取reflect.Type
	rType := reflect.TypeOf(s)
	// 获取reflect.Value
	rVal := reflect.ValueOf(s)
	// 获取kind
	kind := rVal.Kind()
	fmt.Printf("type=%v, value=%v, kind=%v\n", rType, rVal, kind)

	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取该结构体的字段数
	filedNums := rVal.NumField()
	fmt.Printf("该结构体有%d个字段\n", filedNums)

	for i := 0; i < filedNums; i++ {
		value := rVal.Field(i)
		fmt.Printf("Field %d 的值是%v\n", i, value)
		// 获取标签
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != ""{
			fmt.Printf("Filed %d:tag为=%v\n", i, tagVal)
		}
	}

	// 获取结构体方法
	methodsNum := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", methodsNum)

	// val.Method(1).Call()
	var params []reflect.Value // 声明切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rVal.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())

}

func main(){
	monster := Monster{
		Name:"brandon",
		Age:40,
		Score:30.78,
	}
	TestStruct(monster)
}