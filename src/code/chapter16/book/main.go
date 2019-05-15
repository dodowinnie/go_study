package main
import(
	"fmt"
	// "time"
)

func worker(ch chan string){
	fmt.Println("get into worker....")
	ch <- "hello"
}

func main()  {
	ch := make(chan string)
	go worker(ch)
	<-ch

}
