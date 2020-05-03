package main

import (
	"fmt"
)


type customerView struct {
	key string
	loop bool
}


func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n----------------客户信息管理软件----------------")
		fmt.Println("                   1. 添加客户")
		fmt.Println("                   2. 修改客户")
		fmt.Println("                   3. 删除客户")
		fmt.Println("                   4. 客户列表")
		fmt.Println("                   5. 退    出")
		fmt.Print("请选择（1-5）：")
		
		fmt.Scanln(&this.key)
		
		switch this.key {
			case "1" :
				fmt.Println("添加客户")
			case "2" :
				fmt.Println("修改客户")
			case "3" :
				fmt.Println("删除客户")
			case "4" :
				fmt.Println("客户列表")
			case "5" :
				fmt.Println("退    出")
			default :
				fmt.Println("输入错误")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("已退出")
}

func main() {
	customerView := customerView {
		key: "",
		loop: true,
	}
	customerView.mainMenu();
}