package main

import(
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}){
	// 获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)
	// 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v, rVal Type=%T\n", rVal, rVal)

	// reflect.Value{}
	iV := rVal.Interface()
	// 将interface通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

	kind := rVal.Kind()
	fmt.Println("kind is ", kind)
}

func reflectTest02(b interface{}){
	// 获取refelctValue
	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iV=%v, iv Type = %T\n", iV, iV)
	stu, ok := iV.(Sutdent)
	if ok {
		fmt.Printf("stu.Name=%v,stu.Age=%v", stu.Name, stu.Age)
	}
}

type Sutdent struct{
	Name string
	Age int
}

// type Monster struct{
// 	Name string
// 	Age int
// }


func main(){
	var num int = 100
	reflectTest01(num)

	// stu := Sutdent{
	// 	Name : "brandon",
	// 	Age : 20,
	// }
	// reflectTest02(stu)

}