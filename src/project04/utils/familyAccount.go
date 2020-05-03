package utils

import "fmt"

type FamilyAccount struct {
	key string
	loop bool
	balance float64
	money float64
	note string
	details string
	// 是否有收支情况
	flag bool
	choise string
}


func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount {
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		details : "收支\t账户金额\t收支金额\t说明",
		flag : false,
		choise : "",
	}
}

func (this *FamilyAccount) showDetails() {
	if this.flag {
		fmt.Println("当前收支明细记录")
		fmt.Println(this.details)
	} else {
		fmt.Println("当前无收支明细")
	}
}

func (this *FamilyAccount) income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) outcome() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		return
	}
	this.balance -= this.money
	fmt.Print("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit() {
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

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n----------------家庭收支记账软件----------------")
		fmt.Println("                  1. 收支明细")
		fmt.Println("                  2. 登记收入")
		fmt.Println("                  3. 登记支出")
		fmt.Println("                  4. 退    出")
		fmt.Print("请选择（1-4）：")
		
		fmt.Scanln(&this.key)
		
		switch this.key {
			case "1" :
				this.showDetails()
			case "2" :
				this.income()
			case "3" :
				this.outcome()
			case "4" :
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