package main

import (
	"fmt"
	"io"
	"net"
	"project08/common/message"
	process2 "project08/server/process"
	"project08/server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	fmt.Printf("serverProcessMes:%v,type:%v\n",mes, mes.Type)
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在")
	}

	return
}

func (this *Processor) Process2() (err error) {

	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Process2 json.Unmarshal fail err=", err)
				return err
			} else {
				fmt.Println("Process2 json.Unmarshal fail err=", err)
				return err
			}
		}

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
