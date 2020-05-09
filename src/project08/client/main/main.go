package main

import (
	"fmt"
	"project08/client/process"
)

var userId int
var userPwd string
var userName string

func main() {

	var key int

	for true {
		fmt.Println("\n----------------欢迎到登录多人聊天系统----------------")
		fmt.Println("                   1. 登录聊天室")
		fmt.Println("                   2. 注册用户")
		fmt.Println("                   3. 退出系统")
		fmt.Print("请选择（1-3）：")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Print("请输入用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Print("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Print("用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Print("密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Print("昵称:")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Print("退出系统")
		default:
			fmt.Println("输入错误")
		}
	}

	// if key == 1 {
	// 	fmt.Print("请输入用户的id:")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Print("请输入用户的密码:")
	// 	fmt.Scanf("%s\n", &userPwd)

	// 	login(userId, userPwd)

	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功")
	// 	// }

	// } else if key == 2 {
	// 	fmt.Print("处理用户注册逻辑...")
	// }

	fmt.Println("已退出")
}
