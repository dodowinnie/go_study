package main
import (
	"fmt"
	
)

func main(){
	// var day byte
	// fmt.Println("请输入代号")
	// fmt.Scanf("%c", &day)

	// switch day {
	// 	case 'a' :
	// 		fmt.Println("today is moday")
	// 	case 'b' :
	// 		fmt.Println("today is tuesday")
	// 	case 'c' :
	// 		fmt.Println("today is wensday")
	// 	case 'd' :
	// 		fmt.Println("today is thursday")
	// 	case 'e' :
	// 		fmt.Println("today is friday")
	// 	case 'f' :
	// 		fmt.Println("today is saturday")
	// 	case 'g' :
	// 		fmt.Println("today is sunday")
	// 	default:
	// 		fmt.Println("wrong input")

	// 	}

// 	var age int = 10
// 	switch {
// 		case age == 10 :
// 			fmt.Println("age == 10")
// 		case age == 20 :
// 			fmt.Println("age == 20")
// 		default :
// 			fmt.Println("age is nothing")
// 	}

// 	var score int = 30

// 	switch {
// 	case score > 90 :
// 		fmt.Println("good")
// 	case score >= 70 && score <= 90 :
// 		fmt.Println("common")
// 	case score >= 60 && score < 70 :
// 		fmt.Println("ok")
// 	default :
// 		fmt.Println("fail")
// }

// switch穿透 fallthrough
var num int = 10
switch num {
	case 10 :
		fmt.Println("ok1")
		fallthrough
	case 20 :
		fmt.Println("ok2")
	case 30 :
		fmt.Println("ok3")
	default:
		fmt.Println("nothing")
}
	



}