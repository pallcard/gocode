package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"project08/client/utils"
	"project08/common/message"
)

func ShowMenu(userId int) {
	fmt.Printf("\n----------------恭喜%d登录成功----------------\n", userId)
	fmt.Println("              1. 显示在线用户列表")
	fmt.Println("              2. 发送消息")
	fmt.Println("              3. 私聊")
	fmt.Println("              4. 退出系统")
	fmt.Print("请选择（1-4）：")
	var key int
	var toUserId int
	var content string

	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Print("发送消息内容:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("私聊")
		fmt.Print("userId:")
		fmt.Scanf("%d\n", &toUserId)
		fmt.Print("发送消息内容:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendPersonMes(toUserId, content)
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误")
	}

}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		fmt.Println("mes=", mes)

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesToGroupType:
			outPutGroupMes(&mes)
		case message.SmsMesToPersonType:
			outPutPersonMes(&mes)
		default:
			fmt.Println("未知消息类型")

		}
	}
}
