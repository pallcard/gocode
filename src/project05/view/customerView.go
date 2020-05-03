package main

import (
	"fmt"
	"project05/service"
	"project05/model"
)


type customerView struct {
	key string
	loop bool
	choise string
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customres := this.customerService.List()
	fmt.Println("------------------------客户列表------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customres); i++ {
		fmt.Println(customres[i].GetInfo())
	}
 
	fmt.Println("----------------------客户列表完成-------------------------")

}

func (this *customerView) add() {
	fmt.Println("------------------------添加客户------------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomerWithoutId(name, gender, age, phone, email)
	
	if this.customerService.Add(customer) {
		fmt.Println("----------------------添加客户完成-------------------------")
	} else {
		fmt.Println("----------------------添加客户失败-------------------------")
	}
 
}

func (this *customerView) delete() {
	fmt.Println("----------------------删除客户-------------------------")
	fmt.Print("请输入删除客户编号(-1退出)：")
	id := 0
	fmt.Scanln(&id)
	if id == -1 {
		return 
	}

	fmt.Println("确认删除吗？y/n:")
	confirm := ""
	for {
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "n" {
			break
		}
		fmt.Print("输入有误，请重新输入 y/n:")
	}

	if confirm == "y" {
		if this.customerService.Delete(id) {
			fmt.Print("----------------------删除客户完成-------------------------")
		} else {
			fmt.Print("----------------------删除失败，id不存在-------------------------")
		}
	}	
}

func (this *customerView) exit() {
	fmt.Print("确认退出吗？y/n:")
	for {
		fmt.Scanln(&this.choise)
		if this.choise == "y" || this.choise == "n" {
			break
		}
		fmt.Print("输入有误，请重新输入 y/n:")
	}
	if this.choise == "y" {
		this.loop = false
	}
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
				this.add()
			case "2" :
				fmt.Println("修改客户")
			case "3" :
				this.delete()
			case "4" :
				this.list()
			case "5" :
				this.exit()
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

	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu();
}