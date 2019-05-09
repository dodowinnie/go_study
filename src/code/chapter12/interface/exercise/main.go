package main

import(
	"fmt"
	"sort"
	"math/rand"
)

type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func (heroSlice HeroSlice) Len() int{
	return len(heroSlice)
} 

func (heroSlice HeroSlice) Less(i, j int) bool{
	return heroSlice[i].Age > heroSlice[j].Age
}

func (heroSlice HeroSlice) Swap(i, j int){
	// tmp := heroSlice[i]
	// heroSlice[i] = heroSlice[j]
	// heroSlice[j] = tmp
	// 等价于上面三行代码
	heroSlice[i], heroSlice[j] = heroSlice[j], heroSlice[i]
}


func main(){

	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	for _, val := range heros{
		fmt.Println(val)
	}
	sort.Sort(heros)
	fmt.Println("====================排序后===========================")
	for _, val := range heros{
		fmt.Println(val)
	}

}