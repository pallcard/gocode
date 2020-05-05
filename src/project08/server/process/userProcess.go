package process2

import (
	"fmt"
	"net"
	"project08/common/message"
	"encoding/json"
	_"encoding/binary"
	"project08/server/utils"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	// id=100, pwd=123456
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
		
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "改用户不存在，请注册使用"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}