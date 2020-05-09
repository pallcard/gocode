package main

import (
	_ "encoding/binary"
	_ "encoding/json"
	_ "errors"
	"fmt"
	_ "io"
	"net"
	_ "project08/common/message"
	"project08/server/model"
	"time"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	fmt.Println("读取客户端发送的数据.....")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		// err = errors.New("read pkg header error")
// 		return
// 	}

// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])

// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		err = errors.New("read pkg body error")
// 		return
// 	}

// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err=", err)
// 		return
// 	}
// 	return
// }

// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err=", err)
// 		return
// 	}

// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	var loginResMes message.LoginResMes

// 	// id=100, pwd=123456
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		loginResMes.Code = 200

// 	} else {
// 		loginResMes.Code = 500
// 		loginResMes.Error = "改用户不存在，请注册使用"
// 	}

// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail err=", err)
// 		return
// 	}

// 	resMes.Data = string(data)

// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail err=", err)
// 		return
// 	}

// 	err = writePkg(conn, data)
// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

// 	// 发送长度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(bytes) fail", err)
// 		return
// 	}

// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(data) fail", err)
// 		return
// 	}
// 	return
// }

// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 		case message.LoginMesType:
// 			err = serverProcessLogin(conn, mes)
// 		case message.RegisterMesType:
// 			//
// 		default:
// 			fmt.Println("消息类型不存在")
// 	}

// 	return
// }

func process(conn net.Conn) {

	defer conn.Close()

	// for {

	// 	mes, err := readPkg(conn)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("json.Unmarshal fail err=", err)
	// 			return
	// 		} else {
	// 			fmt.Println("json.Unmarshal fail err=", err)
	// 			return
	// 		}
	// 	}
	// 	// fmt.Println("mes=", mes)

	// 	err = serverProcessMes(conn, &mes)
	// 	if err != nil {
	// 		return
	// 	}
	// }

	processor := &Processor{
		Conn: conn,
	}

	err := processor.Process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 err=", err)
		return
	}

}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听.....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器.....")

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accpet err=", err)
		}

		go process(conn)

	}

}
