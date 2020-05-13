package process

import (
	"encoding/json"
	"fmt"
	"project08/client/utils"
	"project08/common/message"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User.UserId = CurUser.UserId

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes tf.WritePkg fail=", err)
		return
	}
	fmt.Printf("SendGroupMes success data=%v\n", smsMes)
	return
}