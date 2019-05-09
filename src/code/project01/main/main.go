package main
import (
	"code/project01/utils"	
)

// type Record struct {
// 	Amount float64
// 	Opt string
// 	Message string
// 	Balance float64
// }

// var RecordList []Record = make([]Record, 1)

// var Balance float64


// var amount float64
// var message string


// func MoneyIn(amount float64, message string, preBalance *float64){
// 	*preBalance += amount
// 	record := Record{
// 		Amount : amount,
// 		Message : message,
// 		Opt : "收入",
// 		Balance : *preBalance,
// 	}
// 	RecordList = append(RecordList, record)
// }


// func MoneyOut(amount float64, message string, preBalance *float64){
// 	*preBalance -= amount
// 	record := Record{
// 		Amount : amount,
// 		Message : message,
// 		Opt : "支出",
// 		Balance : *preBalance,
// 	}
// 	RecordList = append(RecordList, record)
// }

// func MoneyInfo(){
// 	fmt.Println()
// 	fmt.Println("------------------当前收支明细------------------")
// 	fmt.Println("收支\t账户金额\t收支金额\t说	明")
// 	for _, record := range RecordList{
// 		fmt.Printf("%v\t%v\t\t%v\t\t%v\n", record.Opt, record.Balance, record.Amount, record.Message)
// 	}
	
// }

func main(){
	family := utils.NewFamilyAcocunt()
	family.Login()
	// for {
	// key := ""
	// fmt.Println("------------------家庭收支记账软件------------------")
	// fmt.Println()
	// fmt.Println("                  1 收支明细                       ")
	// fmt.Println("                  2 登记收入                       ")
	// fmt.Println("                  3 登记支出                       ")
	// fmt.Println("                  4 退    出                       ")
	// fmt.Println()
	// fmt.Print("                  请选择(1-4):")
	// fmt.Scanln(&key)

	
	// 	switch key {
	// 	case "1":
	// 		MoneyInfo()
	// 	case "2":
	// 		fmt.Print("本次收入金额:")
	// 		fmt.Scanln(&amount)
	// 		fmt.Print("本次收入说明:")
	// 		fmt.Scanln(&message)
	// 		MoneyIn(amount, message, &Balance)
	// 	case "3":
	// 		fmt.Print("本次支出金额:")
	// 		fmt.Scanln(&amount)
	// 		fmt.Print("本次支出说明:")
	// 		fmt.Scanln(&message)
	// 		MoneyOut(amount, message, &Balance)
	// 	case "4":
	// 		fmt.Println("感谢使用家庭收支记账软件，再见。。。")
	// 		return
	// 	}
	// }


}
