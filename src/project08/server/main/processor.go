package main

import (
	"fmt"
	"project08/common/message"
	"project08/server/process"
	"net"
	"io"
	"project08/server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//
		default:
			fmt.Println("消息类型不存在")
	}

	return
}

func (this *Processor) Process2() (err error) {

	for {
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("json.Unmarshal fail err=", err)
				return err
			} else {
				fmt.Println("json.Unmarshal fail err=", err)
				return err
			}
		}
		// fmt.Println("mes=", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}