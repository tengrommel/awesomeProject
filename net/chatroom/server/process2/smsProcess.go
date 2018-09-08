package process2

import (
	"awesomeProject/net/chatroom/common/message"
	"awesomeProject/net/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	// 暂时不需要字段
}

// 写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	// 遍历服务器端的map
	// 将信息发送出去
	// 取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}

	for id, up := range userMgr.onlineUsers {
		// 这里我们还要过滤掉自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMessageToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMessageToEachOnlineUser(data []byte, conn net.Conn) {
	// 创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}
