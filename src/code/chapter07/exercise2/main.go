package main
import(
	"fmt"
	"math/rand"
	"time"
)

func getRandom() int {
	rand.Seed(time.Now().Unix())
	 return rand.Intn(10) + 1
}


func main(){
	num := getRandom()
	fmt.Println(num)
	var guess int
	// var count int = 10
	var i int = 1
	label:
	for ; i <= 10; i++ {
		fmt.Println("请猜数字")
		fmt.Scanln(&guess)
		if guess != num {
			if i == 10 {
				fmt.Println("说你啥好")
			}
			continue
		}else {
			switch i {
				case 1 :
					fmt.Println("你真是个天才")
					break label
				case 2, 3:
					fmt.Println("你很聪明，赶上我了")
					break label
				case 4, 5, 6, 7, 8, 9 :
					fmt.Println("一般般")
					break label
				case 10:
					fmt.Println("可算猜对了")
					break label
				default:
					break label
			}
		}
	}
	fmt.Println("游戏结束。。。。")
	


}