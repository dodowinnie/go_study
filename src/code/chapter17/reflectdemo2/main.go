package main

import(
	"fmt"
	"reflect"
)



func main(){
	var v float64 = 1.2
	rVal := reflect.ValueOf(v) // 获取Value
	rNum := rVal.Float() // 获取值
	// rType := reflect.TypeOf(v)
	rType := rVal.Type() // 获取type
	kind := rVal.Kind() // 获取kind
	fmt.Printf("value=%v,data=%v,kind=%v,type=%v\n", rVal, rNum, kind, rType)

	// 转interface
	iV := rVal.Interface()
	res := iV.(float64)
	fmt.Println("res=", res)
}