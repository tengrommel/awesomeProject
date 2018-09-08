package model

import (
	"awesomeProject/net/chatroom/common/message"
	"net"
)

// 因为在客户端，我们很多地方会使用到curUser，我们将使用全局的
type CurUser struct {
	Conn net.Conn
	message.User
}
