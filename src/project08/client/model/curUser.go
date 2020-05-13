package model

import (
	"net"
	"project08/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}

