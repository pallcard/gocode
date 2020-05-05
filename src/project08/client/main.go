package main
import (
	"fmt"
)

var userId int
var userPwd string

func main() {

	var key int

	var loop = true

	for loop {
		fmt.Println("\n----------------欢迎到登录多人聊天系统----------------")
		fmt.Println("                   1. 登录聊天室")
		fmt.Println("                   2. 注册用户")
		fmt.Println("                   3. 退出系统")
		fmt.Print("请选择（1-3）：")
		
		fmt.Scanf("%d\n", &key)
		
		switch key {
			case 1 :
				fmt.Print("登录聊天室")
				loop = false
			case 2 :
				fmt.Print("注册用户")
				loop = false
			case 3 :
				fmt.Print("退出系统")
				loop = false
			default :
				fmt.Println("输入错误")
		}
	}

	if key == 1 {
		fmt.Print("请输入用户的id:")
		fmt.Scanf("%d\n", &userId)
		fmt.Print("请输入用户的密码:")
		fmt.Scanf("%s\n", &userPwd)
		
		err := login(userId, userPwd)

		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}

	} else if key == 2 {
		fmt.Print("处理用户注册逻辑...")
	}

	fmt.Println("已退出")
}