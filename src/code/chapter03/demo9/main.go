package main
// import "fmt"
// import "unsafe"

import(
	"fmt"
	
	
)

func main() {
	var str string = "我是\nbrandon"
	
	fmt.Println("str=", str)

	str2 := `
		package main
		// import "fmt"
		// import "unsafe"
		
		import(
			"fmt"
			
			
		)`

	fmt.Println(str2)

	// 字符串拼接
	var str3 = "hello " + 
	"world "
	str3 += "bardnon"
	fmt.Println(str3)


}