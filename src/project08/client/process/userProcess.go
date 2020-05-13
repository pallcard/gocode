package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"project08/client/utils"
	"project08/common/message"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// fmt.Printf(" userId = %d userPwd = %s", userId, userPwd)
	// return nil

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg() fail", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg() fail", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("注册成功")

	} else {
		fmt.Println(registerResMes.Error)
	}
	os.Exit(0)
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Printf(" userId = %d userPwd = %s", userId, userPwd)
	// return nil

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)

	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	// time.Sleep(20 * time.Second)
	// fmt.Printf("客户端发送完长度=%d\n", len(data), string(data))

	tf := &utils.Transfer{
		Conn: conn,
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) fail", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			// 过滤自己
			if v == userId{
				continue
			}
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

