package main

import "fmt"

func main() {

	key := ""
	
	loop := true

	balance := 10000.0
	money := 0.0
	note := ""
	details := "收支\t账户金额\t收支金额\t说明"

	// 是否有收支情况
	flag := false

	for {
		fmt.Println("\n----------------家庭收支记账软件----------------")
		fmt.Println("                  1. 收支明细")
		fmt.Println("                  2. 登记收入")
		fmt.Println("                  3. 登记支出")
		fmt.Println("                  4. 退    出")
		fmt.Print("请选择（1-4）：")
		
		fmt.Scanln(&key)
		
		switch key {
			case "1" :
				if flag {
					fmt.Println("当前收支明细记录")
					fmt.Println(details)
				} else {
					fmt.Println("当前无收支明细")
				}
			case "2" :
				fmt.Print("本次收入金额：")
				fmt.Scanln(&money)
				balance += money
				fmt.Print("本次收入说明：")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
				flag = true
			case "3" :
				fmt.Print("本次支出金额：")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("余额不足")
					break
				}
				balance -= money
				fmt.Print("本次支出说明：")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				flag = true
			case "4" :
				fmt.Print("确认退出吗？y/n:")
				choise := ""
				for {
					fmt.Scanln(&choise)
					if choise == "y" || choise == "n" {
						break
					}
					fmt.Print("输入有误，请重新输入 y/n:")
				}
				if choise == "y" {
					loop = false
				}
			default :
				fmt.Println("输入错误")
		}


		if !loop {
			break
		}
	}
	fmt.Println("已退出")
}