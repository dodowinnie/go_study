package main

import(
	"fmt"
	"reflect"
)

type Cal struct{
	Num1 int
	Num2 int
}

func (this Cal) GetSub(name string){
	fmt.Printf("%v完成了减法运行，%d - %d = %d", name, this.Num1, this.Num2, (this.Num1 - this.Num2))
}

func doReflect(s interface{}){
	// 获取Value值
	rVal := reflect.ValueOf(s)
	// 获取字段个数
	fieldNum := rVal.NumField()
	// 遍历获取字段信息
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("字段 %d 的值=%v\n", i, rVal.Field(i))
	}

	// 调用方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	rVal.MethodByName("GetSub").Call(params)


}


func main(){
	c := Cal{
		Num1 : 8,
		Num2 : 5,
	}
	doReflect(c)

	

}