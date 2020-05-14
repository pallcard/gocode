package process

import (
	"encoding/json"
	"fmt"
	"project08/common/message"
)

func outPutGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outPutGroupMes json.Unmarshal, err=", err)
	}

	info := fmt.Sprintf("用户id:%d\t 对大家说:%s\t", smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
}

func outPutPersonMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outPutGroupMes json.Unmarshal, err=", err)
	}

	info := fmt.Sprintf("用户id:%d\t 对你说:%s\t", smsMes.User.UserId, smsMes.Content)
	fmt.Println(info)
}